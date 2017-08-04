// Copyright (c) Jeevanandam M. (https://github.com/jeevatkm)
// go-aah/tutorials source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package security

import (
	"aahframework.org/config.v0"
	"aahframework.org/security.v0/authc"
	"aahframework.org/security.v0/authz"
	"github.com/go-aah/tutorials/form-based-auth/app/models"
)

// FormAuthorizationProvider struct implements `authz.Authorizer` interface.
type FormAuthorizationProvider struct {
}

// Init method initializes the FormAuthoriationProvider, this method gets called
// during server start up.
func (fa *FormAuthorizationProvider) Init(cfg *config.Config) error {

	// NOTE: Init is called on application startup

	return nil
}

// GetAuthorizationInfo method is `authz.Authorizer` interface.
//
// GetAuthorizationInfo method gets called after authentication is successful
// to get Subject's (aka User) access control information such as roles and permissions.
func (fa *FormAuthorizationProvider) GetAuthorizationInfo(authcInfo *authc.AuthenticationInfo) *authz.AuthorizationInfo {
	authorities := models.FindUserByEmail(authcInfo.PrimaryPrincipal().Value)

	authzInfo := authz.NewAuthorizationInfo()
	authzInfo.AddRole(authorities.Roles...)
	authzInfo.AddPermissionString(authorities.Permissions...)

	return authzInfo
}
