// Copyright (c) Jeevanandam M. (https://github.com/jeevatkm)
// aahframework.org/examples source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package controllers

import (
	"net/http"

	"aahframe.work/essentials"

	"aahframe.work"

	"aahframework.org/examples/rest-api-basic-auth/app/models"
)

// InfoController is demo purpose to send data for Authentication & Authoriization.
type InfoController struct {
	AppController
}

// ReporteeInfo returns the reportee info for who access of,
// role as `manager` and permission as `user:read:reportee`.
// Look at above Interceptor for authorization check.
func (c *InfoController) ReporteeInfo(email string) {
	userInfo := models.FindUserByEmail(email)
	if userInfo == nil {
		c.Reply().NotFound().JSON(aah.Data{
			"message": "repotee not found",
		})
		return
	}

	// Specfic case: admin user check
	if ess.IsSliceContainsString(userInfo.Roles, "admin") && !c.Subject().HasRole("admin") {
		c.Reply().Forbidden().JSON(aah.Data{"message": "access denied"})
		return
	}

	c.Reply().Ok().JSON(userInfo)
}

// HandleError method invoked by aah error handling flow
// Doc: https://aahframework.org/error-handling.html
func (c *InfoController) HandleError(err *aah.Error) bool {
	switch err.Code {
	case http.StatusForbidden:
		c.Reply().Forbidden().JSON(aah.Data{
			"message": "access denied",
		})
	}
	return true
}
