// Copyright (c) Jeevanandam M. (https://github.com/jeevatkm)
// go-aah/tutorials source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package models

// This is for demo purpose, set of users
// Typically you will be using Database, API calls, LDAP, etc to get the Authentication
// Information.

// Key is Token and value is User
var users map[string]*User

type (
	// User struct hold the user details
	User struct {
		FirstName   string   `json:"first_name"`
		LastName    string   `json:"last_name"`
		Email       string   `json:"email"`
		Password    string   `json:"-"`
		IsLocked    bool     `json:"is_locked"`
		IsExpried   bool     `json:"is_expried"`
		Roles       []string `json:"roles,omitempty"`
		Permissions []string `json:"permission,omitempty"`
	}
)

// FindUserByEmail returns the user information for given email address.
func FindUserByEmail(email string) *User {
	if u, found := users[email]; found {
		uf := *u
		return &uf
	}
	return nil
}

func init() {
	/*
		   Creating User Information
			 Learn about permission: http://docs.aahframework.org/security-permissions.html
	*/
	users = make(map[string]*User)

	users["user1@example.com"] = &User{
		FirstName:   "East",
		LastName:    "Corner",
		Password:    "$2y$10$2A4GsJ6SmLAMvDe8XmTam.MSkKojdobBVJfIU7GiyoM.lWt.XV3H6", // welcome123
		Email:       "user1@example.com",
		Roles:       []string{"employee", "manager"},
		Permissions: []string{"user:read,edit:reportee"},
	}

	users["user2@example.com"] = &User{
		FirstName:   "West",
		LastName:    "Corner",
		Password:    "$2y$10$2A4GsJ6SmLAMvDe8XmTam.MSkKojdobBVJfIU7GiyoM.lWt.XV3H6", // welcome123
		Email:       "user2@example.com",
		Roles:       []string{"employee"},
		Permissions: []string{},
	}

	users["user3@example.com"] = &User{
		FirstName: "South",
		LastName:  "Corner",
		Password:  "$2y$10$2A4GsJ6SmLAMvDe8XmTam.MSkKojdobBVJfIU7GiyoM.lWt.XV3H6", // welcome123
		Email:     "user3@example.com",
		IsLocked:  true,
	}
}
