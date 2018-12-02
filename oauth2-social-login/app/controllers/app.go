package controllers

import (
	"aahframe.work"

	"aahframework.org/examples/oauth2-social-login/app/models"
)

// AppController struct application controller
type AppController struct {
	*aah.Context
}

// Index method is application home page.
func (c *AppController) Index() {
	data := aah.Data{
		"Greet": models.Greet{
			Message: "aah framework - OAuth2 Social Login",
		},
	}

	c.Reply().Ok().HTML(data)
}

// Success method is to display successful sign-in.
func (c *AppController) Success() {
}

// Logout method does logout subject.
func (c *AppController) Logout() {
	c.Subject().Logout()
	c.Reply().Redirect(c.RouteURL("index"))
}
