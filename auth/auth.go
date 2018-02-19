package auth

import (
	"strconv"
	"strings"
	"time"

	ovh "github.com/toorop/go-ovh"
	"github.com/toorop/go-ovh/auth/authswag"
)

// Auth wrap /auth
type Auth struct {
	*ovh.Client
}

// New return new Auth client
func New(client *ovh.Client) *Auth {
	return &Auth{
		client,
	}
}

// RequestCredential return  PostAuthCredentialParamsBodyAccessRules
func (a *Auth) RequestCredential(params authswag.PostAuthCredentialParamsBody) (*authswag.AuthCredential, error) {
	data, err := params.MarshalBinary()
	if err != nil {
		return nil, err
	}
	body, err := a.POST("/auth/credential", data)
	if err != nil {
		return nil, err
	}
	authCredential := new(authswag.AuthCredential)
	err = authCredential.UnmarshalBinary(body)
	return authCredential, err
}

// GetCurrentCredential return current credential
func (a *Auth) GetCurrentCredential() (*authswag.APICredential, error) {
	body, err := a.GET("/auth/currentCredential")
	if err != nil {
		return nil, err
	}
	credential := new(authswag.APICredential)
	err = credential.UnmarshalBinary(body)
	return credential, err
}

// GetOVHTime return OVH time
func (a *Auth) GetOVHTime() (OVHTime time.Time, err error) {
	body, err := a.GET("/auth/time")
	if err != nil {
		return
	}
	t, err := strconv.Atoi(strings.TrimSpace(string(body)))
	if err != nil {
		return
	}
	OVHTime = time.Unix(int64(t), 0)
	return
}

// ExpireCredential expire current credential
func (a *Auth) ExpireCredential() error {
	_, err := a.POST("/auth/logout", nil)
	return err
}
