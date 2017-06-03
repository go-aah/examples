package controllers

import (
	"aahframework.org/aah.v0"
	"github.com/go-aah/tutorials/html-minify/app/models"

	// Import the aah framework HTML minifier
	// Note: Minifier gets applied only to `prod` environment profile.
	_ "github.com/aah-cb/minify"
)

// App struct application controller
type App struct {
	*aah.Context
}

// Index method is application home page.
func (a *App) Index() {
	data := aah.Data{
		"Greet": models.Greet{
			Message: "Welcome to aah framework - Web Application",
		},
	}

	a.Reply().Ok().HTML(data)
}
