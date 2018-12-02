package controllers

import (
	"aahframe.work"
	"aahframework.org/examples/domain-subdomain/app/models"
)

// AppController struct application controller
type AppController struct {
	*aah.Context
}

// Index method is application home page.
func (c *AppController) Index() {
	data := aah.Data{
		"Greet": models.Greet{
			Message: "Example  - Domain, Subdomain, and Wildcard Subdomain",
		},
	}

	c.Reply().Ok().HTML(data)
}
