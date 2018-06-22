// Copyright (c) Jeevanandam M. (https://github.com/jeevatkm)
// aahframework.org/examples source code and usage is governed by a MIT style
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
		Token       string   `json:"-"`
		IsLocked    bool     `json:"is_locked"`
		IsExpried   bool     `json:"is_expried"`
		Roles       []string `json:"roles,omitempty"`
		Permissions []string `json:"permission,omitempty"`
	}
)

// FindUserByToken returns the user information for given token.
func FindUserByToken(token string) *User {
	if u, found := users[token]; found {
		uf := *u
		return &uf
	}
	return nil
}

// FindUserByEmail returns the user information for given email address.
func FindUserByEmail(email string) *User {
	for _, u := range users {
		if u.Email == email {
			ui := *u
			return &ui
		}
	}
	return nil
}

func init() {
	/*
		   Creating User Information
			 Learn about permission: http://docs.aahframework.org/security-permissions.html
	*/
	users = make(map[string]*User)

	users["2F405AF4FE65632ABFD1BCD584A447902FC77556CBEDF64CDADB9A289F736609"] = &User{
		FirstName:   "East",
		LastName:    "Corner",
		Token:       "2F405AF4FE65632ABFD1BCD584A447902FC77556CBEDF64CDADB9A289F736609",
		Email:       "user1@aahframework.org",
		Roles:       []string{"employee", "manager"},
		Permissions: []string{"user:read,edit:reportee"},
	}

	users["AE23F1211A28F480036BF425575FE08EE4183D3F64BF967FA39A33022F1C2E7E"] = &User{
		FirstName:   "West",
		LastName:    "Corner",
		Token:       "AE23F1211A28F480036BF425575FE08EE4183D3F64BF967FA39A33022F1C2E7E",
		Email:       "user2@aahframework.org",
		Roles:       []string{"employee"},
		Permissions: []string{},
	}

	users["68BEE0D2D97980E1337114713C3C26B49A45F5595F0E93FEAEAAC75ECCC4981C"] = &User{
		FirstName: "South",
		LastName:  "Corner",
		Token:     "68BEE0D2D97980E1337114713C3C26B49A45F5595F0E93FEAEAAC75ECCC4981C",
		Email:     "user3@aahframework.org",
		IsLocked:  true,
	}

	users["832273819090D7D036E2D529A47A554C41A37C8E060C5F7FFBC4CEBEFCDECF64"] = &User{
		FirstName:   "Admin",
		LastName:    "Corner",
		Token:       "832273819090D7D036E2D529A47A554C41A37C8E060C5F7FFBC4CEBEFCDECF64",
		Email:       "admin@aahframework.org",
		Roles:       []string{"employee", "manager", "admin"},
		Permissions: []string{"user:read,edit,delete:reportee"},
	}
}
