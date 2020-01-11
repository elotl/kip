package aws

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/awstesting/unit"
)

// Not used, but this is the signed plaintext contained in the PKCS7 data below
const instanceIdentityDocument = `{
    "devpayProductCodes" : null,
    "availabilityZone" : "us-east-1d",
    "privateIp" : "10.158.112.84",
    "version" : "2010-08-31",
    "region" : "us-east-1",
    "instanceId" : "i-1234567890abcdef0",
    "billingProducts" : null,
    "instanceType" : "t1.micro",
    "accountId" : "123456789012",
    "pendingTime" : "2015-11-19T16:32:11Z",
    "imageId" : "ami-5fb8c835",
    "kernelId" : "aki-919dcaf8",
    "ramdiskId" : null,
    "architecture" : "x86_64",
    "marketplaceProductCodes" : ["mkt-1234"]
}
`

const PKCS7 = `MIIFbwYJKoZIhvcNAQcCoIIFYDCCBVwCAQExCTAHBgUrDgMCGjCCAhgGCSqGSIb3
DQEHAaCCAgkEggIFewogICAgImRldnBheVByb2R1Y3RDb2RlcyIgOiBudWxsLAog
ICAgImF2YWlsYWJpbGl0eVpvbmUiIDogInVzLWVhc3QtMWQiLAogICAgInByaXZh
dGVJcCIgOiAiMTAuMTU4LjExMi44NCIsCiAgICAidmVyc2lvbiIgOiAiMjAxMC0w
OC0zMSIsCiAgICAicmVnaW9uIiA6ICJ1cy1lYXN0LTEiLAogICAgImluc3RhbmNl
SWQiIDogImktMTIzNDU2Nzg5MGFiY2RlZjAiLAogICAgImJpbGxpbmdQcm9kdWN0
cyIgOiBudWxsLAogICAgImluc3RhbmNlVHlwZSIgOiAidDEubWljcm8iLAogICAg
ImFjY291bnRJZCIgOiAiMTIzNDU2Nzg5MDEyIiwKICAgICJwZW5kaW5nVGltZSIg
OiAiMjAxNS0xMS0xOVQxNjozMjoxMVoiLAogICAgImltYWdlSWQiIDogImFtaS01
ZmI4YzgzNSIsCiAgICAia2VybmVsSWQiIDogImFraS05MTlkY2FmOCIsCiAgICAi
cmFtZGlza0lkIiA6IG51bGwsCiAgICAiYXJjaGl0ZWN0dXJlIiA6ICJ4ODZfNjQi
LAogICAgIm1hcmtldHBsYWNlUHJvZHVjdENvZGVzIiA6IFsibWt0LTEyMzQiXQp9
CqCCAfMwggHvMIIBWKADAgECAgUA72hG0TANBgkqhkiG9w0BAQsFADApMRAwDgYD
VQQKEwdBY21lIENvMRUwEwYDVQQDEwxFZGRhcmQgU3RhcmswHhcNMTkwNDIzMjMz
MzQxWhcNMjAwNDIzMjMzMzQxWjAlMRAwDgYDVQQKEwdBY21lIENvMREwDwYDVQQD
EwhKb24gU25vdzCBnzANBgkqhkiG9w0BAQEFAAOBjQAwgYkCgYEAwbDp7DNrzctV
eY85resud7vXZkY1SrSbCM87NMI2V57yyC7YLdLE9/6sGwSiG1nR3OYBcoKuVtF/
W2EQGfNhQlvMZOnRIyd9vI5YexpMh8FEKKkrCF6gttJzjxNqbFM1NHoJO7tTeowY
vjSkM3V8t4xqTl16x1UBYpIMNbpSkzsCAwEAAaMnMCUwDgYDVR0PAQH/BAQDAgWg
MBMGA1UdJQQMMAoGCCsGAQUFBwMEMA0GCSqGSIb3DQEBCwUAA4GBAGo4Hnefxb47
4ULNOBWHhH11hV7b43FGXkop2WbFeYf8Za+8ga/GsqRcXLdTXiOLHJx/Bu5m9SIn
BcWtsQFRWghlp7qeX6OuUQLqr13pjStVgMsC2LCGyt2AUiRK0XwSfMuF8JtasslP
n2nzEuvQ2BEmgVKELtpJ6xxHFPWOavWnMYIBNzCCATMCAQEwMjApMRAwDgYDVQQK
EwdBY21lIENvMRUwEwYDVQQDEwxFZGRhcmQgU3RhcmsCBQDvaEbRMAcGBSsOAwIa
oGEwGAYJKoZIhvcNAQkDMQsGCSqGSIb3DQEHATAgBgkqhkiG9w0BCQUxExcRMTkw
NDIzMTYzMzQxLTA3MDAwIwYJKoZIhvcNAQkEMRYEFCQjXHrlJgWvXqZCf8ZjwCqn
J/OTMAsGCSqGSIb3DQEBBQSBgKaER+G5wa86n/VmGtvFYlvRaHqsv79zpUo4cJQj
3TAGGv0I7mODyHdmQjP7GmfuXml03lWpJWS6PJzWkqYwJLI1ud2e+LKWBkOSdagq
e/cc/xiC47yvYIG9INXO5oGiJ0usMgsJKcHNaVGCEh9KqMRHT4P6ZPhXSIPoL+lc
OJ53`

const testCert = `-----BEGIN CERTIFICATE-----
MIIB7zCCAVigAwIBAgIFAO9oRtEwDQYJKoZIhvcNAQELBQAwKTEQMA4GA1UEChMH
QWNtZSBDbzEVMBMGA1UEAxMMRWRkYXJkIFN0YXJrMB4XDTE5MDQyMzIzMzM0MVoX
DTIwMDQyMzIzMzM0MVowJTEQMA4GA1UEChMHQWNtZSBDbzERMA8GA1UEAxMISm9u
IFNub3cwgZ8wDQYJKoZIhvcNAQEBBQADgY0AMIGJAoGBAMGw6ewza83LVXmPOa3r
Lne712ZGNUq0mwjPOzTCNlee8sgu2C3SxPf+rBsEohtZ0dzmAXKCrlbRf1thEBnz
YUJbzGTp0SMnfbyOWHsaTIfBRCipKwheoLbSc48TamxTNTR6CTu7U3qMGL40pDN1
fLeMak5desdVAWKSDDW6UpM7AgMBAAGjJzAlMA4GA1UdDwEB/wQEAwIFoDATBgNV
HSUEDDAKBggrBgEFBQcDBDANBgkqhkiG9w0BAQsFAAOBgQBqOB53n8W+O+FCzTgV
h4R9dYVe2+NxRl5KKdlmxXmH/GWvvIGvxrKkXFy3U14jixycfwbuZvUiJwXFrbEB
UVoIZae6nl+jrlEC6q9d6Y0rVYDLAtiwhsrdgFIkStF8EnzLhfCbWrLJT59p8xLr
0NgRJoFShC7aSescRxT1jmr1pw==
-----END CERTIFICATE-----`

func initTestServer(pathResponseMap map[string]string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := pathResponseMap[r.RequestURI]
		if response == "" {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		w.Write([]byte(response))
	}))
}

func TestEC2RoleProviderInstanceIdentity(t *testing.T) {
	server := initTestServer(map[string]string{
		"/latest/dynamic/instance-identity/pkcs7": PKCS7,
	})
	defer server.Close()
	c := ec2metadata.New(unit.Session, &aws.Config{Endpoint: aws.String(server.URL + "/latest")})

	doc, err := getInstanceIdentityDocument(c, testCert)
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
	if e, a := "123456789012", doc.AccountID; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := "us-east-1d", doc.AvailabilityZone; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := "us-east-1", doc.Region; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if len(doc.MarketplaceProductCodes) == 0 {
		t.Errorf("expected a marketplace product code in document")
	} else if e, a := "mkt-1234", doc.MarketplaceProductCodes[0]; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}
