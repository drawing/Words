package basic

import (
	"errors"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
	"net/http"
)

var user_collection *mgo.Collection
var course_collection *mgo.Collection
var session_collection *mgo.Collection
var dictionary_collection *mgo.Collection

func InitStorage(host string, dbname string, user string, pass string) {
	session, err := mgo.Dial(host)
	if err != nil {
		log.Fatalln("mgo dial:", err)
	}
	db := session.DB(dbname)
	err = db.Login(user, pass)
	if err != nil {
		log.Fatalln("mgo login:", err)
	}
	course_collection = db.C("Course")
	user_collection = db.C("User")
	session_collection = db.C("Session")
	dictionary_collection = db.C("Dictionary")
}

func GetLoginStudent(w http.ResponseWriter, r *http.Request) *Student {
	name, _ := r.Cookie("username")
	identifier, _ := r.Cookie("identifier")

	if name == nil || identifier == nil {
		return nil
	}

	log.Println("access", name, identifier)

	var s Student
	err := session_collection.Find(bson.M{"username": name.Value}).One(&s)
	if err != nil {
		log.Println("QUERY_SESSION_ERROR:", err)
		return nil
	}

	return &s
}

func AddUser(name string, email string, pass string) error {
	var nUser User
	nUser.ID = bson.NewObjectId()
	nUser.Familiars = map[string]bool{}
	nUser.Name = name
	nUser.Email = email
	nUser.Password = pass
	nUser.Words = map[string]Word{}

	adduser_error := errors.New("add user failed")
	user_exist_error := errors.New("user already exist")
	email_exist_error := errors.New("email already exist")

	num, err := user_collection.Find(bson.M{"name": name}).Count()
	if err != nil || num != 0 {
		return user_exist_error
	}

	num, err = user_collection.Find(bson.M{"email": email}).Count()
	if err != nil || num != 0 {
		return email_exist_error
	}

	err = user_collection.Insert(&nUser)
	if err != nil {
		log.Println("CREATE_USER_ERROR:", err)
		return adduser_error
	}

	log.Println("CREATE_USER_SUCC:", name, email, pass)
	return nil
}

func GetArticleByID(id string) *CourseDetail {
	var course CourseDetail

	err := course_collection.FindId(bson.ObjectIdHex(id)).One(&course)
	if err != nil {
		log.Println("course find:", err)
		return nil
	}

	return &course
}
