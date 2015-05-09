package action

import (
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"

	"log"
)

func Login(req LoginReq, session sessions.Session, r render.Render) {
	log.Println("login: email=", req.Email, ", token=", req.Token)
	r.JSON(200, map[string]interface{}{"code": 0})
}

func Lessons(req LessonsReq, session sessions.Session, r render.Render) {
	log.Println("lessons: page=", req.Page)
	var resp LessonsResp
	resp.Code = 100
	resp.Body = "test"
	r.JSON(200, resp)
}

func NewLesson(req NewLessonReq, session sessions.Session, r render.Render) {
}

func DelLesson(req DelLessonReq, session sessions.Session, r render.Render) {
}

func Study(req StudyReq, session sessions.Session, r render.Render) {
}

func Report(req ReportReq, session sessions.Session, r render.Render) {
}
