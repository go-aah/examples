package controllers

import "aahframework.org/aah.v0"

// App struct application controller
type App struct {
	*aah.Context
}

// Index method is to redirect root directory to locaized page.
func (a *App) Index() {
	a.Reply().Redirect(a.ReverseURL("index_lang", "en"))
}

// IndexLang method is application home page.
func (a *App) IndexLang() {
	a.Reply().Ok().HTMLf("index.html", nil)
}
