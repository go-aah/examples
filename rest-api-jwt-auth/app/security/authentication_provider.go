// Copyright (c) Jeevanandam M. (https://github.com/jeevatkm)
// aahframework.org/examples source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package security

import (
	"strings"

	"aahframework.org/config.v0"
	"aahframework.org/essentials.v0"
	"aahframework.org/log.v0"
	"aahframework.org/security.v0/authc"
	"github.com/dgrijalva/jwt-go"
)

var _ authc.Authenticator = (*AuthenticationProvider)(nil)

// AuthenticationProvider struct implements `authc.Authenticator` interface.
type AuthenticationProvider struct{}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// AuthenticationProvider methods
//___________________________________

// Init method initializes the AuthenticationProvider, this method gets called
// during server start up.
func (a *AuthenticationProvider) Init(cfg *config.Config) error {
	// NOTE: Init is called on application startup
	return nil
}

// GetAuthenticationInfo method is `authc.Authenticator` interface. It is called by Security Manager.
func (a *AuthenticationProvider) GetAuthenticationInfo(authcToken *authc.AuthenticationToken) (*authc.AuthenticationInfo, error) {
	// authcToken.Identity is complete value of "Authorization" HTTP header or
	// configured `security.auth_schemes.jwt_auth.header.identity` header value
	//
	// values[0] ==> Bearer
	// values[1] ==> JWT token
	values := strings.Fields(authcToken.Identity)
	if len(values) != 2 || ess.IsStrEmpty(values[1]) {
		log.Error("Token is not provided")
		return nil, authc.ErrAuthenticationFailed
	}

	// Validate the JWT token
	token, err := jwt.Parse(values[1], jwtKeyFunc)
	if err != nil || !token.Valid {
		log.Errorf("Error occurred '%v' or token is not valid", err)
		return nil, authc.ErrAuthenticationFailed
	}

	claims := token.Claims.(jwt.MapClaims)

	// User found and token validated.
	// Now create authentication info and return to the framework
	authcInfo := authc.NewAuthenticationInfo()
	authcInfo.Principals = append(authcInfo.Principals,
		&authc.Principal{
			Realm:     "inmemory",
			Claim:     "Username",
			Value:     claims["username"].(string),
			IsPrimary: true,
		})

	return authcInfo, nil
}
