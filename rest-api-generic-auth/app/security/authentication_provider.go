// Copyright (c) Jeevanandam M. (https://github.com/jeevatkm)
// go-aah/tutorials source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package security

import (
	"fmt"
	"strings"

	"aahframework.org/config.v0"
	"aahframework.org/security.v0/authc"
	"github.com/go-aah/tutorials/rest-api-generic-auth/app/models"
)

var _ authc.Authenticator = (*AuthenticationProvider)(nil)

// AuthenticationProvider struct implements `authc.Authenticator` interface.
type AuthenticationProvider struct {
}

// Init method initializes the AuthenticationProvider, this method gets called
// during server start up.
func (a *AuthenticationProvider) Init(cfg *config.Config) error {

	fmt.Println("AuthenticationProvider init called")

	// NOTE: Init is called on application startup

	return nil
}

// GetAuthenticationInfo method is `authc.Authenticator` interface. It is called by Security Manager.
func (a *AuthenticationProvider) GetAuthenticationInfo(authcToken *authc.AuthenticationToken) (*authc.AuthenticationInfo, error) {

	// authcToken.Identity is complete value of "Authorization" HTTP header or
	// configured `security.auth_schemes.generic_auth.header.identity` header value
	values := strings.Fields(authcToken.Identity)
	if len(values) != 2 {
		return nil, authc.ErrAuthenticationFailed
	}

	// find the user
	user := models.FindUserByToken(values[1])

	// User Credential validation is developer responsibility,
	// framework gives you complete control on this.
	if user == nil {
		// No subject exists, return nil and error
		return nil, authc.ErrSubjectNotExists
	}

	if user.IsExpried || user.IsLocked {
		return nil, authc.ErrAuthenticationFailed
	}

	// User found and token validated.
	// Now create authentication info and return to the framework
	authcInfo := authc.NewAuthenticationInfo()
	authcInfo.Principals = append(authcInfo.Principals,
		&authc.Principal{
			Value:     user.Email,
			IsPrimary: true,
			Realm:     "inmemory",
		})

	return authcInfo, nil
}
