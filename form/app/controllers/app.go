// Copyright (c) Jeevanandam M. (https://github.com/jeevatkm)
// aahframework.org/examples source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package controllers

import (
	"html"
	"html/template"

	"aahframe.work"
	"aahframework.org/examples/form/app/models"
)

// AppController struct application controller
type AppController struct {
	*aah.Context
}

// Index method is application home page.
func (c *AppController) Index() {
	data := aah.Data{
		"Greet": models.Greet{
			Message: "Welcome to aah framework - Form Submission Example",
		},
	}

	c.Reply().Ok().HTML(data)
}

// UserProfile method displays the user profile info page.
func (c *AppController) UserProfile() {
	c.Reply().Ok()
}

// UserProfileSubmit method receives Form submission
func (c *AppController) UserProfileSubmit(user *models.User) {
	c.Reply().HTMLf("userprofile.html", aah.Data{
		"DataDisplay": true,
		"User":        user,
		"AboutYou":    template.HTML(html.UnescapeString(user.AboutYou)),
	})
}
