// Copyright (c) Jeevanandam M. (https://github.com/jeevatkm)
// aahframework.org/examples source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package admin

import (
	"net/http"

	"aahframe.work"
)

// DashboardController struct demo implementation of admin scope.
type DashboardController struct {
	*aah.Context
}

// Index method dispay the admin Dashboard page.
func (c *DashboardController) Index() {
	c.Reply().Ok()
}

// HandleError method invoked by aah error handling flow
// Doc: https://aahframework.org/error-handling.html
func (c *DashboardController) HandleError(err *aah.Error) bool {
	switch err.Code {
	case http.StatusForbidden:
		c.Reply().Forbidden().HTMLf("/access-denied.html", nil)
	}
	return true
}
