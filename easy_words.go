package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"

	"./app"

	"log"
)

func login(req app.LoginReq, session sessions.Session) string {
	log.Println("login: email=", req.Email, ", token=", req.Token)
	r.JSON(200, map[string]interface{}{"code", 0})
}

func main() {
	m := martini.Classic()

	store := sessions.NewCookieStore([]byte("words"))

	m.Use(session.Sessions("words_session", store))
	m.Use(render.Renderer())

	m.Post("/api/login", binding.Bind(app.LoginReq{}), login)

	m.Run()
}
