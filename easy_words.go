package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"

	"./action"
)

func main() {
	m := martini.Classic()

	store := sessions.NewCookieStore([]byte("words"))

	m.Use(sessions.Sessions("words_session", store))
	m.Use(render.Renderer())

	m.Post("/api/login", binding.Bind(action.LoginReq{}), action.Login)

	m.Post("/api/newlesson", binding.Bind(action.NewLessonReq{}), action.NewLesson)
	m.Post("/api/dellesson", binding.Bind(action.DelLessonReq{}), action.DelLesson)
	m.Get("/api/lessons", binding.Bind(action.LessonsReq{}), action.Lessons)

	m.Post("/api/study", binding.Bind(action.StudyReq{}), action.Study)
	m.Post("/api/report", binding.Bind(action.ReportReq{}), action.Report)

	m.Run()
}
