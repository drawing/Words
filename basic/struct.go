package basic

import (
	"labix.org/v2/mgo/bson"
)

type CourseDetail struct {
	ID      bson.ObjectId `bson:"_id"`
	Name    string
	Content string
}

type CourseBrief struct {
	ID   bson.ObjectId
	Name string
}

type Word struct {
	Name       string
	Status     string
	CourseName string
}

type User struct {
	ID        bson.ObjectId `bson:"_id"`
	Name      string
	Email     string
	Password  string
	Courses   []CourseBrief
	Familiars map[string]bool
	Words     map[string]Word
}

type TranPron struct {
}

type Translation struct {
	ID   bson.ObjectId `bson:"_id"`
	Name string
}
