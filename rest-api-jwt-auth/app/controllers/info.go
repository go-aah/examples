// Copyright (c) Jeevanandam M. (https://github.com/jeevatkm)
// go-aah/tutorials source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package controllers

import (
	"aahframework.org/aah.v0"
	"github.com/go-aah/tutorials/rest-api-jwt-auth/app/models"
)

// InfoController is to demostrate the REST API endpoints for Authentication & Authorization.
type InfoController struct {
	AppController
}

// BeforeReporteeInfo Interceptor checks authority of
// role as `manager` and permission as `user:read:reportee`.
func (i *InfoController) BeforeReporteeInfo() {
	if !i.Subject().HasRole("manager") ||
		!i.Subject().IsPermitted("user:read:reportee") {
		i.Reply().Forbidden().JSON(aah.Data{
			"message": "access denied",
		})

		// abort the flow
		i.Abort()
	}
}

// ReporteeInfo returns the reportee info for who access of,
// role as `manager` and permission as `user:read:reportee`.
// Look at above Interceptor for authorization check.
func (i *InfoController) ReporteeInfo(email string) {
	userInfo := models.FindUserByEmail(email)
	if userInfo == nil {
		i.Reply().NotFound().JSON(aah.Data{
			"message": "repotee not found",
		})
		return
	}

	i.Reply().Ok().JSON(userInfo)
}
