// Copyright (c) Jeevanandam M. (https://github.com/jeevatkm)
// aahframework.org/examples source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package controllers

import (
	"time"

	"aahframe.work"
	"aahframe.work/essentials"
	"aahframework.org/examples/rest-api-jwt-auth/app/models"
	"aahframework.org/examples/rest-api-jwt-auth/app/security"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// AppController struct application controller
type AppController struct {
	*aah.Context
}

// Index method is application root API endpoint.
func (c *AppController) Index() {
	c.Reply().Ok().JSON(models.Greet{
		Message: "Welcome to aah framework - REST API JWT Auth using Generic Auth scheme",
	})
}

// Token method validates the given username and password the
// generates the JWT token.
func (c *AppController) Token(userToken *models.UserToken) {
	if ess.IsStrEmpty(userToken.Username) || ess.IsStrEmpty(userToken.Password) {
		c.Reply().BadRequest().JSON(aah.Data{
			"message": "bad request",
		})
		return
	}

	// get the user details by username
	user := models.FindUserByEmail(userToken.Username)
	if user == nil || user.IsExpried || user.IsLocked {
		c.Reply().Unauthorized().JSON(aah.Data{
			"message": "invalid credentials",
		})
		return
	}

	// validate password
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(userToken.Password)); err != nil {
		c.Reply().Unauthorized().JSON(aah.Data{
			"message": "invalid credentials",
		})
		return
	}

	// create JWT claims
	claims := jwt.MapClaims{}
	claims["username"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate JWT token string
	token, err := security.GenerateToken(claims)
	if err != nil {
		c.Log().Error(err)
		c.Reply().InternalServerError().JSON(aah.Data{
			"message": "Whoops! something went wrong...",
		})
		return
	}

	// everything went good, respond token
	c.Reply().Ok().JSON(aah.Data{
		"token_type": "bearer",
		"token":      token,
	})
}
