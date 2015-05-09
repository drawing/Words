package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"

	"./action"

	"log"
)

func login(req action.LoginReq, session sessions.Session, r render.Render) {
	log.Println("login: email=", req.Email, ", token=", req.Token)
	r.JSON(200, map[string]interface{}{"code": 0})
}

func main() {
	m := martini.Classic()

	store := sessions.NewCookieStore([]byte("words"))

	m.Use(sessions.Sessions("words_session", store))
	m.Use(render.Renderer())

	m.Post("/api/login", binding.Bind(action.LoginReq{}), login)

	m.Run()
}
