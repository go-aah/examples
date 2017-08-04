// Copyright (c) Jeevanandam M. (https://github.com/jeevatkm)
// go-aah/tutorials source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package security

import (
	"aahframework.org/aah.v0"
	"aahframework.org/config.v0"
	"aahframework.org/security.v0/authc"
	"github.com/go-aah/tutorials/form-based-auth/app/models"
)

var _ authc.Authenticator = (*FormAuthenticationProvider)(nil)

// FormAuthenticationProvider struct implements `authc.Authenticator` interface.
type FormAuthenticationProvider struct {
}

// Init method initializes the FormAuthenticationProvider, this method gets called
// during server start up.
func (fa *FormAuthenticationProvider) Init(cfg *config.Config) error {

	// NOTE: Init is called on application startup

	return nil
}

// GetAuthenticationInfo method is `authc.Authenticator` interface
func (fa *FormAuthenticationProvider) GetAuthenticationInfo(authcToken *authc.AuthenticationToken) (*authc.AuthenticationInfo, error) {

	// Form Auth Values from authcToken
	// 		authcToken.Identity => username
	// 		authcToken.Credential => passowrd

	user := models.FindUserByEmail(authcToken.Identity)
	if user == nil {
		// No subject exists, return nil and error
		return nil, authc.ErrSubjectNotExists
	}

	// User found, now create authentication info and return to the framework
	authcInfo := authc.NewAuthenticationInfo()
	authcInfo.Principals = append(authcInfo.Principals,
		&authc.Principal{
			Value:     user.Email,
			IsPrimary: true,
			Realm:     "inmemory",
		})
	authcInfo.Credential = []byte(user.Password)
	authcInfo.IsLocked = user.IsLocked
	authcInfo.IsExpired = user.IsExpried

	return authcInfo, nil
}

func postAuthEvent(e *aah.Event) {
	ctx := e.Data.(*aah.Context)

	// Populate session info after authentication
	user := models.FindUserByEmail(ctx.Subject().PrimaryPrincipal().Value)
	ctx.Session().Set("FirstName", user.FirstName)
	ctx.Session().Set("LastName", user.LastName)
}

func init() {
	aah.OnPostAuth(postAuthEvent)
}
