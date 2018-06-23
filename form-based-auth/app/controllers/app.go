// Copyright (c) Jeevanandam M. (https://github.com/jeevatkm)
// aahframework.org/examples source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package controllers

import (
	"aahframework.org/aah.v0"
	"aahframework.org/examples/form-based-auth/app/models"
)

// AppController struct application controller
type AppController struct {
	*aah.Context
}

// Index method is application home page.
func (a *AppController) Index() {
	data := aah.Data{
		"Greet": models.Greet{
			Message: "Welcome to aah framework - Form Based Auth Example",
		},
	}

	a.Reply().Ok().HTML(data)
}

// BeforeLogin method action is interceptor of Login.
func (a *AppController) BeforeLogin() {
	if a.Subject().IsAuthenticated() {
		a.Reply().Redirect(a.RouteURL("index"))
		a.Abort()
	}
}

// Login method presents the login page.
func (a *AppController) Login() {
	a.Reply().Ok()
}

// Logout method does logout currently logged in subject (aka user).
func (a *AppController) Logout() {
	a.Subject().Logout()

	// Send it to login page or whatever the page you have to send the user
	// after logout
	a.Reply().Redirect(a.RouteURL("login"))
}

// BeforeManageUsers method is action interceptor of ManageUsers.
func (a *AppController) BeforeManageUsers() {
	// Checking roles and permissions
	if !a.Subject().HasAnyRole("manager", "administrator") ||
		!a.Subject().IsPermitted("users:manage:view") {
		a.Reply().Forbidden().HTMLf("/access-denied.html", nil)
		a.Abort()
	}
}

// ManageUsers method presents the manage user page afer verifying
// Authorization
func (a *AppController) ManageUsers() {
	// looks okay, present the page
	a.Reply().Ok().HTML(nil)
}
