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
func (c *AppController) Index() {
	data := aah.Data{
		"Greet": models.Greet{
			Message: "Welcome to aah framework - Form Based Auth Example",
		},
	}

	c.Reply().Ok().HTML(data)
}

// BeforeLogin method action is interceptor of Login.
func (c *AppController) BeforeLogin() {
	if c.Subject().IsAuthenticated() {
		c.Reply().Redirect(c.RouteURL("index"))
		c.Abort()
	}
}

// Login method presents the login page.
func (c *AppController) Login() {
	c.Reply().Ok()
}

// Logout method does logout currently logged in subject (aka user).
func (c *AppController) Logout() {
	c.Subject().Logout()

	// Send it to login page or whatever the page you have to send the user
	// after logout
	c.Reply().Redirect(c.RouteURL("login"))
}

// BeforeManageUsers method is action interceptor of ManageUsers.
// func (c *AppController) BeforeManageUsers() {
// 	// Checking roles and permissions
// 	if !c.Subject().HasAnyRole("manager", "administrator") ||
// 		!c.Subject().IsPermitted("users:manage:view") {
// 		c.Reply().Forbidden().HTMLf("/access-denied.html", nil)
// 		c.Abort()
// 	}
// }

// ManageUsers method presents the manage user page afer verifying
// Authorization
func (c *AppController) ManageUsers() {
	// looks okay, present the page
	c.Reply().Ok().HTML(nil)
}
