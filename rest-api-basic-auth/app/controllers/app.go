package controllers

import (
	"time"

	"aahframework.org/aah.v0"
	"github.com/go-aah/tutorials/rest-api-basic-auth/app/models"
)

// AppController struct application controller
type AppController struct {
	*aah.Context
}

// Index method is application root API endpoint.
func (a *AppController) Index() {
	a.Reply().Ok().JSON(models.Greet{
		Message: "Welcome to aah framework - API application",
	})
}

// UserInfo method returns current subject Principal Value and access time.
func (a *AppController) UserInfo() {
	a.Reply().Ok().JSON(models.UserInfo{
		PrincipalValue: a.Subject().PrimaryPrincipal().Value,
		AccessTime:     time.Now(),
	})
}

// BeforeUserInfoWithMsg method to check authorization.
// Interceptor is best spot for checking authorization.
func (a *AppController) BeforeUserInfoWithMsg() {
	if !a.Subject().HasRole("manager") {
		a.Reply().Forbidden().JSON(aah.Data{
			"code":    403,
			"message": "you don't have permission to access this endpoint",
		})

		// Abort if subject does not have proper authorization
		a.Abort()
	}
}

// UserInfoWithMsg method just to demo permission, nothing fancy here.
func (a *AppController) UserInfoWithMsg() {
	a.Reply().Ok().JSON(models.UserInfo{
		PrincipalValue: a.Subject().PrimaryPrincipal().Value,
		AccessTime:     time.Now(),
		Message:        "Shown to user who has role 'manager'",
	})
}
