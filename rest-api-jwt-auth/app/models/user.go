// Copyright (c) Jeevanandam M. (https://github.com/jeevatkm)
// aahframework.org/examples source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package models

// This is for demo purpose, set of users
// Typically you will be using Database, API calls, LDAP, etc to get the Authentication
// Information.

// Key is email and value is User
var users map[string]*User

type (
	// User struct hold the user details
	User struct {
		FirstName   string   `json:"first_name"`
		LastName    string   `json:"last_name"`
		Email       string   `json:"email"`
		Password    []byte   `json:"-"`
		IsLocked    bool     `json:"is_locked"`
		IsExpried   bool     `json:"is_expried"`
		Roles       []string `json:"roles,omitempty"`
		Permissions []string `json:"permission,omitempty"`
	}

	// UserToken struct used for token creation method.
	UserToken struct {
		Username string `json:"username"`
		Password string `json:"password"`
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

	users["user1@aahframework.org"] = &User{
		FirstName:   "East",
		LastName:    "Corner",
		Password:    []byte("$2y$10$2A4GsJ6SmLAMvDe8XmTam.MSkKojdobBVJfIU7GiyoM.lWt.XV3H6"), // welcome123
		Email:       "user1@aahframework.org",
		Roles:       []string{"employee", "manager"},
		Permissions: []string{"user:read,edit:reportee"},
	}

	users["user2@aahframework.org"] = &User{
		FirstName:   "West",
		LastName:    "Corner",
		Password:    []byte("$2y$10$2A4GsJ6SmLAMvDe8XmTam.MSkKojdobBVJfIU7GiyoM.lWt.XV3H6"), // welcome123
		Email:       "user2@aahframework.org",
		Roles:       []string{"employee"},
		Permissions: []string{},
	}

	users["user3@aahframework.org"] = &User{
		FirstName: "South",
		LastName:  "Corner",
		Password:  []byte("$2y$10$2A4GsJ6SmLAMvDe8XmTam.MSkKojdobBVJfIU7GiyoM.lWt.XV3H6"), // welcome123
		Email:     "user3@aahframework.org",
		IsLocked:  true,
	}

	users["admin@aahframework.org"] = &User{
		FirstName:   "Admin",
		LastName:    "Corner",
		Password:    []byte("$2y$10$2A4GsJ6SmLAMvDe8XmTam.MSkKojdobBVJfIU7GiyoM.lWt.XV3H6"), // welcome123
		Email:       "admin@aahframework.org",
		Roles:       []string{"employee", "manager", "admin"},
		Permissions: []string{"user:read,edit,delete:reportee"},
	}

}
