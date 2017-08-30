// Copyright (c) Jeevanandam M. (https://github.com/jeevatkm)
// go-aah/tutorials source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package models

// User struct holds user profile info
type User struct {
	FirstName string   `bind:"first_name"`
	LastName  string   `bind:"last_name"`
	Email     string   `bind:"email"`
	AboutYou  string   `bind:"about_you"`
	Residence *Address `bind:"residence"`
}

// Address struct holds address information
type Address struct {
	Address1 string `bind:"address1"`
	Address2 string `bind:"address2"`
	City     string `bind:"city"`
	ZipCode  string `bind:"zip_code"`
}
