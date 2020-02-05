package nodeclient

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	neturl "net/url"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/dghubble/sling"
	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/elotl/cloud-instance-provider/pkg/util"
	"github.com/elotl/cloud-instance-provider/pkg/util/timeoutmap"
	"github.com/elotl/wsstream"
	"github.com/gorilla/websocket"
	"k8s.io/klog"
)

var (
	clientTTL                = time.Minute * 10
	SAVE_LOG_BYTES           = 4096
	MULTIPART_PKG_FIELD_NAME = "package"
	ItzoPort                 = 6421
	ServerName               = "MilpaNode"
)

type ItzoClientFactoryer interface {
	// We pass in the whole NetworkAddress here since we might want to
	// connect to either the public or private IP of the node/pod,
	// depending on whether we are inside the cloud network
	GetClient([]api.NetworkAddress) NodeClient
	GetWSStream([]api.NetworkAddress, string) (*wsstream.WSStream, error)
	DeleteClient([]api.NetworkAddress)
}

type ItzoClientFactory struct {
	tlsConfig   *tls.Config
	clients     *timeoutmap.TimeoutMap
	usePublicIP bool
}

func NewItzoFactory(rootCert *x509.Certificate, cert tls.Certificate, usePublicIP bool) *ItzoClientFactory {
	caCertPool := x509.NewCertPool()
	caCertPool.AddCert(rootCert)
	clientFactory := &ItzoClientFactory{
		tlsConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
			RootCAs:      caCertPool,
			ServerName:   ServerName,
		},
		clients:     timeoutmap.New(false, nil),
		usePublicIP: usePublicIP,
	}
	go clientFactory.clients.Start(30 * time.Second)
	return clientFactory
}

func (fac *ItzoClientFactory) getAddress(addy []api.NetworkAddress) string {
	if fac.usePublicIP {
		return api.GetPublicIP(addy)
	} else {
		return api.GetPrivateIP(addy)
	}
}

func (fac *ItzoClientFactory) GetClient(addy []api.NetworkAddress) NodeClient {
	// get client from map if exists, otherwise create the client
	var newClient *ItzoClient
	ip := fac.getAddress(addy)
	client, exists := fac.clients.Get(ip)
	if !exists {
		newClient = NewItzoClient(ip, fac.tlsConfig)
		fac.clients.Add(ip, newClient, clientTTL, timeoutmap.Noop)
	} else {
		newClient = client.(*ItzoClient)
		fac.clients.Checkin(ip)
	}
	return newClient
}

func (fac *ItzoClientFactory) GetWSStream(addy []api.NetworkAddress, path string) (*wsstream.WSStream, error) {
	ip := fac.getAddress(addy)
	addr := fmt.Sprintf("%s:%d", ip, ItzoPort)
	u := url.URL{
		Scheme: "wss",
		Host:   addr,
		Path:   path,
	}
	dialer := &websocket.Dialer{
		TLSClientConfig:  fac.tlsConfig.Clone(),
		HandshakeTimeout: 10 * time.Second,
		Proxy:            http.ProxyFromEnvironment,
	}
	header := http.Header{}
	header.Add("Content-Type", "application/x-www-form-urlencoded")
	conn, resp, err := dialer.Dial(u.String(), header)
	if err != nil {
		if resp != nil {
			defer resp.Body.Close()
			bodyContents, bodyerr := ioutil.ReadAll(resp.Body)
			if bodyerr == nil {
				e := fmt.Errorf("Websocket dial error: %v - %s",
					err, string(bodyContents))
				klog.Error(e)
				return nil, e
			}
		}
		return nil, util.WrapError(err, "Dial error")
	}
	ws := wsstream.NewWSStream(conn)
	return ws, nil
}

func (fac *ItzoClientFactory) DeleteClient(addy []api.NetworkAddress) {
	ip := fac.getAddress(addy)
	fac.clients.Delete(ip)
}

type ItzoClient struct {
	instanceIp string
	baseURL    string

	// Go's http clients are safe for concurrent use by goroutines:
	// https://golang.org/src/net/http/client.go
	//
	// Todo, see if we can get away with only one client somehow (and
	// make short timeouts work)
	httpClient        *http.Client
	healthcheckClient *http.Client
}

func NewItzoClient(instanceIp string, tlsConfig *tls.Config) *ItzoClient {
	return &ItzoClient{
		instanceIp: instanceIp,
		baseURL:    fmt.Sprintf("https://%s:%d/", instanceIp, ItzoPort),
		// The main timeout was arbitrarily chosen.  It was made to be
		// very large since large containers might take a long time to
		// download.  We might need to specify different timeouts for
		// the stages of the connection lifetime.  For now use a
		// relaxed timeout.  I've added a dial & handshake timeout to
		// help with hangs that can happen after a node goes away.  In
		// that case, operations might fail but at least we don't hang
		// milpa.
		httpClient: &http.Client{
			Timeout: 120 * time.Second,
			Transport: &http.Transport{
				TLSClientConfig: tlsConfig.Clone(),
				Dial: (&net.Dialer{
					Timeout: 10 * time.Second,
				}).Dial,
				TLSHandshakeTimeout: 10 * time.Second,
			},
		},
		// This should be less than the node controller's heartbeat
		// interval
		healthcheckClient: &http.Client{
			Timeout: 3 * time.Second,
			Transport: &http.Transport{
				TLSClientConfig: tlsConfig.Clone(),
			},
		},
	}
}

// Used in testing, not sure if this is kosher
func NewItzoWithClient(serverURL string, client *http.Client) *ItzoClient {
	baseURL := serverURL + "/"
	return &ItzoClient{
		instanceIp:        serverURL,
		baseURL:           baseURL,
		httpClient:        client,
		healthcheckClient: client,
	}
}

func handleResponse(resp *http.Response, err error) (string, error) {
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bodyContents, err := ioutil.ReadAll(resp.Body)
	// todo: if debug level is really low, print out the request URL
	// in the message
	if err != nil {
		return "", util.WrapError(err, "Error getting response body for request")
	}
	if resp.StatusCode != 200 {
		// not all response bodies from Itzo are json, don't decocde
		// json response body at this time
		err := fmt.Errorf(
			"Server responded with status code %d.  Response body: %s",
			resp.StatusCode,
			bodyContents,
		)
		return string(bodyContents), err
	}
	return string(bodyContents), nil
}

func createUrl(base, upath string) string {
	u, err := neturl.Parse(base)
	if err != nil {
		panic("Can't parse base URL")
	}
	u.Path = path.Join(u.Path, upath)
	return u.String()
}

func (c *ItzoClient) ResizeVolume() error {
	req, err := sling.New().Post(c.baseURL).
		Path("rest/v1/resizevolume").Request()
	if err != nil {
		return util.WrapError(err, "Error creating request to resize volume")
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return util.WrapError(err, "Error with request to resize volume")
	}
	_, err = handleResponse(resp, err)
	if err != nil {
		return util.WrapError(err, "Error getting response for resize request")
	}
	return nil
}

func (c *ItzoClient) GetStatus() (*api.PodStatusReply, error) {
	url := createUrl(c.baseURL, "rest/v1/status")
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	s := api.PodStatusReply{}
	err = json.Unmarshal(body, &s) //dumpStruct, "", "\t")
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (c *ItzoClient) Healthcheck() error {
	url := c.baseURL + "rest/v1/ping"
	resp, err := c.healthcheckClient.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	// see if we can cast body to int and make sure it's positive
	responseBody := strings.ToLower(string(body))
	if responseBody != "pong" {
		return fmt.Errorf("Expected pong, got %s", responseBody)
	}
	return nil
}

func (c *ItzoClient) GetLogs(unit string, lines, bytes int) ([]byte, error) {
	url := c.baseURL + "rest/v1/logs/" + unit
	if lines > 0 || bytes > 0 {
		url = url + fmt.Sprintf("?lines=%d&bytes=%d", lines, bytes)
	}
	resp, err := c.httpClient.Get(url)
	if err != nil {
		klog.Errorf("Error getting logs from %s: %s", c.instanceIp, err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		klog.Errorf("Error reading log reply from %s: %s", c.instanceIp, err)
		return nil, err
	}
	if resp.StatusCode/200 != 1 {
		klog.Errorf("HTTP error getting log from %s: %s (%d); %s",
			c.instanceIp, resp.Status, resp.StatusCode, string(body))
		return nil, fmt.Errorf("Failed to fetch logs: %s (%d); %s",
			resp.Status, resp.StatusCode, string(body))
	}
	return body, nil
}

func (c *ItzoClient) GetFile(path string, lines, bytes int) ([]byte, error) {
	v := neturl.Values{}
	v.Set("path", path)
	if lines > 0 {
		v.Set("lines", strconv.Itoa(lines))
	}
	if bytes > 0 {
		v.Set("bytes", strconv.Itoa(bytes))
	}
	qs := v.Encode()
	url := c.baseURL + "rest/v1/file/?" + qs

	// Todo: combine with logs getter?
	resp, err := c.httpClient.Get(url)
	if err != nil {
		klog.Errorf("Error getting file from %s: %s", c.instanceIp, err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		klog.Errorf("Error reading file reply from %s: %s", c.instanceIp, err)
		return nil, err
	}
	if resp.StatusCode/200 != 1 {
		klog.Errorf("HTTP error getting file from %s: %s (%d); %s",
			c.instanceIp, resp.Status, resp.StatusCode, string(body))
		return nil, fmt.Errorf("Failed to fetch file: %s (%d); %s",
			resp.Status, resp.StatusCode, string(body))
	}
	return body, nil
}

func (c *ItzoClient) UpdateUnits(pp api.PodParameters) error {
	url := c.baseURL + "rest/v1/updatepod"
	b, err := json.Marshal(pp)
	if err != nil {
		return util.WrapError(err, "Could not serialize pod update")
	}
	buf := bytes.NewBuffer(b)
	resp, err := c.httpClient.Post(url, "application/json", buf)
	_, err = handleResponse(resp, err)
	if err != nil {
		klog.Errorf("Error sending pod update to %s: %v", c.instanceIp, err)
		return util.WrapError(err, "Error sending pod update to %s",
			c.instanceIp)
	}
	return nil
}

func (c *ItzoClient) Deploy(pod, name string, data io.Reader) error {
	pr, pw := io.Pipe()
	writer := multipart.NewWriter(pw)
	defer writer.Close()
	// Read from the pipe in a goroutine. Each write to the pipe blocks until
	// the reader side consumes that particular chunk of data.
	ch := make(chan error, 1)
	go func() {
		fullUrl := createUrl(
			c.baseURL, fmt.Sprintf("rest/v1/deploy/%s/%s", pod, name))
		klog.V(2).Infof("deploying package to %s", fullUrl)
		req, err := http.NewRequest("POST", fullUrl, pr)
		if err != nil {
			klog.Errorf("Error creating new deploy POST request: %v\n", err)
			ch <- err
			return
		}
		req.Header.Add("Content-Type", writer.FormDataContentType())
		resp, err := c.httpClient.Do(req)
		if err != nil {
			klog.Errorf("Error sending deploy POST request: %v\n", err)
			ch <- err
			return
		}
		if _, err = handleResponse(resp, err); err != nil {
			klog.Errorf("Error response %+v to deploy POST request: %v\n",
				*resp, err)
			ch <- err
			return
		}
		ch <- nil
	}()
	// Now we can create the actual multipart form. Writing it will go through
	// the pipe, read by the goroutine performing the POST request.
	part, err := writer.CreateFormFile(MULTIPART_PKG_FIELD_NAME, name)
	if err != nil {
		return util.WrapError(err, "Error creating multipart form")
	}
	_, err = io.Copy(part, data)
	if err != nil {
		return util.WrapError(err, "Error copying package data")
	}
	if err = writer.Close(); err != nil {
		return util.WrapError(err, "Error closing multipart writer")
	}
	if err = pw.Close(); err != nil {
		return util.WrapError(err, "Error closing multipart pipe writer")
	}
	err = <-ch
	if err != nil {
		return util.WrapError(err, "Error uploading package")
	}
	return nil
}

func (c *ItzoClient) RunCmd(cmdParams api.RunCmdParams) (string, error) {
	url := c.baseURL + "rest/v1/runcmd/"
	b, err := json.Marshal(cmdParams)
	if err != nil {
		return "", util.WrapError(err, "Could not serialize command")
	}
	buf := bytes.NewBuffer(b)
	resp, err := c.httpClient.Post(url, "application/json", buf)
	output, err := handleResponse(resp, err)
	if err != nil {
		return "", util.WrapError(err, "Error running command on node")
	}
	return output, nil
}

func StreamLogsEndpoint(unitName string, withMetadata bool) string {
	metadataflag := 0
	if withMetadata {
		metadataflag = 1
	}
	return fmt.Sprintf("rest/v1/logs/%s?follow=1&metadata=%d", unitName, metadataflag)
}

func PortForwardEndpoint() string {
	return fmt.Sprintf("rest/v1/portforward/")
}

func ExecEndpoint() string {
	return fmt.Sprintf("rest/v1/exec/")
}

func AttachEndpoint() string {
	return fmt.Sprintf("rest/v1/attach/")
}
