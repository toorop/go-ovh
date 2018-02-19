package auth

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	ovh "github.com/toorop/go-ovh"
	"github.com/toorop/go-ovh/auth/authswag"
)

func TestRequestCredential(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"consumerKey":"Ck","state":"expired","validationUrl":"https://api"}`)
	}))
	accessRules := authswag.PostAuthCredentialParamsBodyAccessRules{
		&authswag.AuthAccessRule{
			Method: "GET",
			Path:   "/me/*",
		},
	}

	body := authswag.PostAuthCredentialParamsBody{
		AccessRules: accessRules,
	}

	authCredential, err := New(ovh.NewClient("ovh-eu").WithEndpoint(ts.URL)).RequestCredential(body)
	assert.NoError(t, err)
	assert.Equal(t, "Ck", authCredential.ConsumerKey)
	assert.Equal(t, "expired", authCredential.State)
	assert.Equal(t, "https://api", authCredential.ValidationURL)
}

// GetCurrentCredential
func TestGetCurrentCredential(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"applicationId":0,"creation":"2018-02-17T16:22:23Z","credentialId":0,"expiration":"2018-02-17T16:22:23Z","lastUse":"2018-02-17T16:22:23Z","ovhSupport":true,"rules":[{"method":"DELETE","path":"string"}],"status":"expired"}`)
	}))
	c := New(ovh.NewClient("ovh-eu").WithEndpoint(ts.URL))
	credential, err := c.GetCurrentCredential()
	assert.NoError(t, err)
	assert.Equal(t, int64(0), credential.ApplicationID)
	assert.Equal(t, true, credential.OvhSupport)
	assert.Equal(t, "DELETE", credential.Rules[0].Method)
}

// GetOVHTime
func TestGetOVHTime(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "1518890161")
	}))
	OVHTime, err := New(ovh.NewClient("ovh-eu").WithEndpoint(ts.URL)).GetOVHTime()
	assert.NoError(t, err)
	assert.Equal(t, int64(1518890161), OVHTime.Unix())
}

// ExpireCredential
func TestExpireCredential(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "")
	}))
	err := New(ovh.NewClient("ovh-eu").WithEndpoint(ts.URL)).ExpireCredential()
	assert.NoError(t, err)
}
