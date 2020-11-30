/*
Copyright 2020 Elotl Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package nodeclient

import (
	"crypto/tls"
	"net/http"
	"net/http/httptest"
	"testing"
)

const okResponseBody = "123"

func OKResponse(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(okResponseBody))
}

func ErrorResponse(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(400)
}

func setupClientServer(success bool) (*ItzoClient, *httptest.Server) {
	var s *httptest.Server
	if success {
		s = httptest.NewTLSServer(http.HandlerFunc(OKResponse))
	} else {
		s = httptest.NewTLSServer(http.HandlerFunc(ErrorResponse))
	}
	c := NewItzoClient("1.2.3.4", &tls.Config{})
	c.baseURL = s.URL + "/"
	c.httpClient = s.Client()
	return c, s
}

func TestGetLogsHappyPath(t *testing.T) {
	c, s := setupClientServer(true)
	defer s.Close()
	logs, err := c.GetLogs("", 100, 0)
	if err != nil {
		t.Errorf("Error on logs happy path: %v", err)
	}
	if string(logs) != okResponseBody {
		t.Errorf("Expected response body to be %s, got %s", okResponseBody, logs)
	}
}

func TestGetLogsError(t *testing.T) {
	c, s := setupClientServer(false)
	defer s.Close()
	_, err := c.GetLogs("", 100, 0)
	if err == nil {
		t.Errorf("Gettings logs error path did not return any errors")
	}
}
