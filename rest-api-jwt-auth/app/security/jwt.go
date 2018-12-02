// Copyright (c) Jeevanandam M. (https://github.com/jeevatkm)
// aahframework.org/examples source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package security

import (
	"fmt"
	"log"

	"aahframe.work"

	"github.com/dgrijalva/jwt-go"
)

// JWT config values
var (
	JWTSigningMethod string
	JWTSigningKey    []byte
)

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Package methods
//___________________________________

// LoadJWTconfig JWT setting into variable on server startup.
func LoadJWTconfig(_ *aah.Event) {
	JWTSigningMethod = aah.App().Config().StringDefault("security.auth_schemes.jwt_auth.sign.method", "HS256")

	if signingKey, found := aah.App().Config().String("security.auth_schemes.jwt_auth.sign.key"); found {
		JWTSigningKey = []byte(signingKey)
	} else {
		log.Fatal("security.auth_schemes.jwt_auth.sign.key is required")
	}
}

func GenerateToken(claims jwt.MapClaims) (string, error) {
	token := createToken()
	token.Claims = claims

	// Generate encoded token and send it as response.
	signedToken, err := token.SignedString(JWTSigningKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// JWT Unexported methods
//___________________________________

// createToken return token based on config value.
// `security.auth_schemes.jwt_auth.sign.method` on security.conf
func createToken() *jwt.Token {
	switch JWTSigningMethod {
	case "HS256":
		return jwt.New(jwt.SigningMethodHS256)
	case "HS384":
		return jwt.New(jwt.SigningMethodHS384)
	case "HS512":
		return jwt.New(jwt.SigningMethodHS512)
	}
	return nil
}

// jwtKeyFunc methdo to check the signing method
func jwtKeyFunc(t *jwt.Token) (interface{}, error) {
	if t.Method.Alg() != JWTSigningMethod {
		return nil, fmt.Errorf("jwt: unexpected signing method '%v'", t.Header["alg"])
	}
	return JWTSigningKey, nil
}
