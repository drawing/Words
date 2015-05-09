package action

import (
// "github.com/martini-contrib/binding"
)

// use http://mholt.github.io/json-to-go/

type RSPHeader struct {
	Code int
}

type LoginReq struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type LoginResp struct {
	RSPHeader
}

type LessonsReq struct {
	Page int `form:"page"`
}

type LessonsResp struct {
	RSPHeader
	Body string
}

type NewLessonReq struct {
}

type NewLessonResp struct {
	RSPHeader
}

type DelLessonReq struct {
}

type DelLessonResp struct {
	RSPHeader
}

type StudyReq struct {
}

type StudyResp struct {
	RSPHeader
}

type ReportReq struct {
}

type ReportResp struct {
	RSPHeader
}
