// Copyright (c) Jeevanandam M. (https://github.com/jeevatkm)
// aahframework.org/examples source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package controllers

import (
	"fmt"
	"net/http"

	"aahframework.org/aah.v0"
	"aahframework.org/examples/rest-api-jwt-auth/app/models"
)

// InfoController is to demostrate the REST API endpoints for Authentication & Authorization.
type InfoController struct {
	AppController
}

// BeforeReporteeInfo Interceptor checks authority of
// role as `manager` and permission as `user:read:reportee`.
// func (i *InfoController) BeforeReporteeInfo() {
// 	if !i.Subject().HasRole("manager") ||
// 		!i.Subject().IsPermitted("user:read:reportee") {
// 		i.Reply().Forbidden().JSON(aah.Data{
// 			"message": "access denied",
// 		})

// 		// abort the flow
// 		i.Abort()
// 	}
// }

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

// HandleError method invoked by aah error handling flow
// Doc: https://aahframework.org/error-handling.html
func (i *InfoController) HandleError(err *aah.Error) bool {
	fmt.Println("err", err)
	switch err.Code {
	case http.StatusForbidden:
		i.Reply().Forbidden().JSON(aah.Data{
			"message": "access denied",
		})
	}
	return true
}
