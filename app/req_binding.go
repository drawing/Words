package app

import (
// "github.com/martini-contrib/binding"
)

// use http://mholt.github.io/json-to-go/

type LoginReq struct {
	Email string `json:"email"`
	Token string `json:"token"`
}
