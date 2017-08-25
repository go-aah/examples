// Copyright (c) Jeevanandam M. (https://github.com/jeevatkm)
// go-aah/tutorials source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

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
