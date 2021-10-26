package api

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type ClientType int

const (
	CLIENTTYPE_NONE ClientType = iota
	CLIENTTYPE_DNS
	CLIENTTYPE_SONAR
	CLIENTTYPE_SONARDNSV2
)

const dnsBaseURL = "https://api.dns.constellix.com/v4/"
const sonarBaseURL = "https://api.sonar.constellix.com/rest/api/"

type ApiClient struct {
	httpclient 		*http.Client
	apiKey     		string
	secretKey  		string
	baseUrl			string
	clientType		ClientType
	LastResponse 	*http.Response
}

//singleton implementation of dns and sonar clients
var dnsApiClietnImpl *ApiClient = nil
var sonarApiClietnImpl *ApiClient = nil

func initApiClient(apiKey, secretKey string) *ApiClient {
	//existing information about client
	client := &ApiClient{
		apiKey:    apiKey,
		secretKey: secretKey,
	}

	//Setting up the HTTP client for the API call
	var transport *http.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			CipherSuites: []uint16{
				tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
				tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
				tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			},
			PreferServerCipherSuites: true,
			InsecureSkipVerify:       true,
			MinVersion:               tls.VersionTLS11,
			MaxVersion:               tls.VersionTLS12,
		},
	}

	client.httpclient = &http.Client{
		Transport: transport,
	}

	return client
}

//Returns a dns singleton
func GetDnsApiClient(apiKey, secretKey string) *ApiClient {
	if dnsApiClietnImpl == nil {
		dnsApiClietnImpl = initApiClient(apiKey, secretKey)
		dnsApiClietnImpl.baseUrl = dnsBaseURL
		dnsApiClietnImpl.clientType = CLIENTTYPE_DNS
		dnsApiClietnImpl.apiKey = apiKey
		dnsApiClietnImpl.secretKey = secretKey
	}
	return dnsApiClietnImpl
}

//Returns a sonar singleton
func GetSonarApiClient(apiKey, secretKey string) *ApiClient {
	if sonarApiClietnImpl == nil {
		sonarApiClietnImpl = initApiClient(apiKey, secretKey)
		sonarApiClietnImpl.baseUrl = sonarBaseURL
		sonarApiClietnImpl.clientType = CLIENTTYPE_SONAR
		sonarApiClietnImpl.apiKey = apiKey
		sonarApiClietnImpl.secretKey = secretKey
	}
	return sonarApiClietnImpl
}

func getToken(apiKey, secretKey string) string {
	//Extracts epoch time in miliseconds
	time := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)

	//Calculate hmac using secrest key and epoch time
	h := hmac.New(sha1.New, []byte(secretKey))
	h.Write([]byte(time))
	sha := base64.StdEncoding.EncodeToString(h.Sum(nil))

	//Building token as 'apikey:hmac:time'
	token := string(apiKey) + ":" + string(sha) + ":" + string(time)
	return token
}

type DNSErrorResponse struct {
	Message 	string 				`json:"message,omitempty"`
	Errors 		map[string]string  	`json:"errors,omitempty"`
}

func (d *DNSErrorResponse) toError() error {
	var allerrs string = d.Message + "\n"
	for key, val := range d.Errors {
		allerrs = allerrs + fmt.Sprintf("%s: %s, ", key, val)
	}
		
	return fmt.Errorf("%s", allerrs)
}

type DNSErrorArrayResponse struct {
	Message 	string 					`json:"message,omitempty"`
	Errors 		map[string][]string  	`json:"errors,omitempty"`
}

func (d *DNSErrorArrayResponse) toError() error {
	var allerrs string = d.Message + "\n"
	for key, val := range d.Errors {
		for _, errval := range(val){
			allerrs = allerrs + fmt.Sprintf("%s: %s, ", key, errval)
		}
	}
		
	return fmt.Errorf("%s", allerrs)
}

func checkForErrors(resp *http.Response, bodyString string, clientType ClientType) error {
	if resp.StatusCode != http.StatusOK && 
		resp.StatusCode != 201 && 
		resp.StatusCode != 202 && 
		resp.StatusCode != 203 && 
		resp.StatusCode != 204 {

		if clientType == CLIENTTYPE_DNS {
			limitRemaining, ok := resp.Header["X-RateLimit-Remaining"]
			if ok {
				if limitRemaining[0] == "0" {
					return fmt.Errorf("Request Limit Exceeded. Status code %d", resp.StatusCode)
				}
			}
		}
		if len(bodyString) == 0 {
			return fmt.Errorf("Status code %d", resp.StatusCode)
		}

		if (clientType == CLIENTTYPE_SONAR) || (clientType == CLIENTTYPE_SONARDNSV2) {
			return fmt.Errorf("Status code %d. Error: %s", resp.StatusCode, bodyString)
		}

		var data DNSErrorResponse
		err := json.Unmarshal([]byte(bodyString), &data)
		if err != nil {
			var arrdata DNSErrorArrayResponse
			err := json.Unmarshal([]byte(bodyString), &arrdata)
			if err != nil {
				return fmt.Errorf("Status code %d", resp.StatusCode)
			}
			return arrdata.toError()
		}

		return data.toError()
	}
	return nil
}

func (c *ApiClient) makeRequest(method, url string, payload []byte, clientType ClientType) (*http.Request, error) {
	//Defining http request
	var req *http.Request
	var err error
	if method == "POST" || method == "PUT" {
		req, err = http.NewRequest(method, url, bytes.NewBuffer(payload))
	} else {
		req, err = http.NewRequest(method, url, nil)
	}
	if err != nil {
		return nil, err
	}

	//Calling for token and setting headers
	token := getToken(c.apiKey, c.secretKey)

	if clientType == CLIENTTYPE_SONAR {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("x-cns-security-token", token)
	} else if clientType == CLIENTTYPE_SONARDNSV2 {
		req.Header.Set("Content-Type", "application/vnd.sonar.v2+json")
		req.Header.Set("x-cns-security-token", token)
	} else {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	}

	return req, nil
}

func (c *ApiClient) DoGet(endpoint string, clientType ClientType) (string, error) {
	
	var url = c.baseUrl + endpoint

	req, err := c.makeRequest("GET", url, nil, clientType)
	if err != nil {
		return "", err
	}

	resp, err1 := c.httpclient.Do(req)
	if err1 != nil {
		return "", err1
	}
	
	c.LastResponse = resp
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	return bodyString, checkForErrors(resp, bodyString, clientType)
}

func (c *ApiClient) DoPost(endpoint string, payload []byte, clientType ClientType) (string, error) {
	
	var url = c.baseUrl + endpoint

	req, err := c.makeRequest("POST", url, payload, clientType)
	if err != nil {
		return "", err
	}

	resp, err1 := c.httpclient.Do(req)
	if err1 != nil {
		return "", err1
	}

	c.LastResponse = resp
	bodyBytes, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		return "", err2
	}

	bodyString := string(bodyBytes)

	return bodyString, checkForErrors(resp, bodyString, clientType)
}

func (c *ApiClient) DoPut(endpoint string, payload []byte, clientType ClientType) (string, error) {
	
	var url = c.baseUrl + endpoint

	req, err := c.makeRequest("PUT", url, payload, clientType)
	if err != nil {
		return "", err
	}

	resp, err1 := c.httpclient.Do(req)
	if err1 != nil {
		return "", err1
	}

	c.LastResponse = resp
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	return bodyString, checkForErrors(resp, bodyString, clientType)
}

func (c *ApiClient) DoDelete(endpoint string, clientType ClientType) (string, error) {

	var url = c.baseUrl + endpoint

	req, err := c.makeRequest("DELETE", url, nil, clientType)
	if err != nil {
		return "", err
	}

	resp, err1 := c.httpclient.Do(req)
	if err1 != nil {
		return "", err1
	}

	c.LastResponse = resp
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	return bodyString, checkForErrors(resp, bodyString, clientType)
}
