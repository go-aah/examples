// Copyright (c) Jeevanandam M. (https://github.com/jeevatkm)
// aahframework.org/examples source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package controllers

import "aahframework.org/aah.v0"

// AppController struct application controller
type AppController struct {
	*aah.Context
}

// Index method is application home page.
func (a *AppController) Index() {
	a.Reply().Ok()
}
