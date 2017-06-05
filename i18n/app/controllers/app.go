package controllers

import "aahframework.org/aah.v0"

// App struct application controller
type App struct {
	*aah.Context
}

// Index method is application home page.
func (a *App) Index() {
	a.Reply().Ok()
}
