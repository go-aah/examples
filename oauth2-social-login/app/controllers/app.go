package controllers

import (
	"aahframework.org/aah.v0"

	"aahframework.org/examples/oauth2-social-login/app/models"
)

// AppController struct application controller
type AppController struct {
	*aah.Context
}

// Index method is application home page.
func (a *AppController) Index() {
	data := aah.Data{
		"Greet": models.Greet{
			Message: "aah framework - OAuth2 Social Login",
		},
	}

	a.Reply().Ok().HTML(data)
}

// Success method is to display successful sign-in.
func (a *AppController) Success() {
}

// Logout method does logout subject.
func (a *AppController) Logout() {
	a.Subject().Logout()
	a.Reply().Redirect(a.RouteURL("index"))
}
