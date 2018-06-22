// Copyright (c) Jeevanandam M. (https://github.com/jeevatkm)
// aahframework.org/examples source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package admin

import "aahframework.org/aah.v0"

// DashboardController struct demo implementation of admin scope.
type DashboardController struct {
	*aah.Context
}

// BeforeIndex method is action interceptor of Dashboard.
func (a *DashboardController) BeforeIndex() {
	if !a.Subject().HasRole("administrator") {
		a.Reply().Forbidden().HTMLf("/access-denied.html", nil)
		a.Abort()
	}
}

// Index method dispay the admin Dashboard page.
func (a *DashboardController) Index() {
	a.Reply().Ok()
}
