package ovh

import (
	"bytes"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/toorop/go-ovh/common"
)

const (
	// Version API version
	Version = "1.0"
	// OvhEU for OVH EU
	OvhEU = "https://eu.api.ovh.com/1.0"
	// OvhCA for OVH CA
	OvhCA = "https://ca.api.ovh.com/1.0"
	// OvhUS for OVH US
	OvhUS = "https://api.ovh.us/1.0"
	// KimsufiEU for Kimsufi EU
	KimsufiEU = "https://eu.api.kimsufi.com/1.0"
	// KimsufiCA for Kimsufi CA
	KimsufiCA = "https://ca.api.kimsufi.com/1.0"
	// SoyoustartEU for SoYouStart EU
	SoyoustartEU = "https://eu.api.soyoustart.com/1.0"
	// SoyoustartCA for SoYouStart CA
	SoyoustartCA = "https://ca.api.soyoustart.com/1.0"
	// Runabove is for runabove
	Runabove = "https://api.runabove.com/1.0"
)

// Endpoints map name -> endpoint
var Endpoints = map[string]string{
	"ovh-eu":        OvhEU,
	"ovh-ca":        OvhCA,
	"ovh-us":        OvhUS,
	"kimsufi-eu":    KimsufiEU,
	"kimsufi-ca":    KimsufiCA,
	"soyoustart-eu": SoyoustartEU,
	"soyoustart-ca": SoyoustartCA,
	"runabove":      Runabove,
}

// Client is an HTTP client wrapper for OVH API
type Client struct {
	*http.Client
	Endpoint string
	Keyring  *Keyring
}

// Response reprents a OVH client response
type Response struct {
	Body   []byte
	APIErr common.APIError
}

// NewClient return a new base client
func NewClient(subOVH string) *Client {
	return &Client{
		&http.Client{},
		Endpoints[subOVH],
		nil,
	}
}

// WithKeyring return a client with keyring
func (c *Client) WithKeyring(AK, AS, CK string) *Client {
	c.Keyring = &Keyring{
		AK: AK,
		AS: AS,
		CK: CK,
	}
	return c
}

// WithKeyringFromEnv get keyring from environment var
// Warning no check is done !
func (c *Client) WithKeyringFromEnv() *Client {
	c.Keyring = &Keyring{
		AK: os.Getenv("OVH_AK"),
		AS: os.Getenv("OVH_AS"),
		CK: os.Getenv("OVH_CK"),
	}
	return c
}

// WithEndpoint changes default endpoint
// Usefull for tests
func (c *Client) WithEndpoint(endpoint string) *Client {
	c.Endpoint = endpoint
	return c
}

// GET for HTTP GET requests
func (c *Client) GET(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", c.Endpoint+url, nil)
	if err != nil {
		return nil, err
	}
	return c.do(req)
}

// POST for HTTP POST requests
func (c *Client) POST(url string, data []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", c.Endpoint+url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	return c.do(req)
}

// do do the request
func (c *Client) do(req *http.Request) ([]byte, error) {
	var err error
	req.Header.Add("Content-type", "application/json")
	if c.Keyring != nil {
		// if application key add header X-Ovh-Application
		if c.Keyring.AK != "" {
			req.Header.Add("X-Ovh-Application", c.Keyring.AK)
		}
		if c.Keyring.AS != "" && c.Keyring.CK != "" {
			timestamp := fmt.Sprintf("%d", int32(time.Now().Unix()))
			req.Header.Add("X-Ovh-Timestamp", timestamp)
			req.Header.Add("X-Ovh-Consumer", c.Keyring.CK)
			if err = c.signRequest(req); err != nil {
				return nil, err
			}
		}
	}

	// Do request
	r, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	// get Body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	// Bad status ?
	if r.StatusCode >= 400 {
		apiErr := new(common.APIError)
		if len(body) != 0 {
			err = json.Unmarshal(body, apiErr)
			if apiErr.HTTPCode == 0 {
				apiErr.HTTPCode = int32(r.StatusCode)
			}
			return nil, fmt.Errorf("HTTP status code: %d - Error code: %d : %s", apiErr.HTTPCode, apiErr.ErrorCode, apiErr.Message)
		}
		return nil, fmt.Errorf("HTTP status code: %d - Error code: %d : %s", r.StatusCode, 0, r.Status)
	}
	return body, nil
}

// Timestamp parameter is here for testing purpose (mock current time)
func (c *Client) signRequest(req *http.Request, timestampIn ...string) error {
	var err error
	var timestamp string
	if len(timestampIn) < 1 {
		timestamp = fmt.Sprintf("%d", int32(time.Now().Unix()))
	} else {
		timestamp = timestampIn[0]
	}
	query := strings.Split(req.URL.String(), "?")[0]
	payload := []byte{}
	if req.Body != nil {
		payload, err = ioutil.ReadAll(req.Body)
		if err != nil {
			return err
		}
		req.Body = ioutil.NopCloser(bytes.NewBuffer(payload))
	}
	h := sha1.New()
	toSign := fmt.Sprintf("%s+%s+%s+%s+%s+%s", c.Keyring.AS, c.Keyring.CK, req.Method, query, string(payload), timestamp)
	h.Write([]byte(toSign))
	req.Header.Add("X-Ovh-Signature", fmt.Sprintf("$1$%x", h.Sum(nil)))
	return nil

}
