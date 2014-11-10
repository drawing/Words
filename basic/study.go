package basic

import (
	"errors"
	"log"
	"math/rand"
	"sort"
	"strconv"
	"strings"
)

import (
	"labix.org/v2/mgo/bson"
	// "labix.org/v2/mgo"
)

var find_error error = errors.New("find user")
var login_error error = errors.New("login")
var auth_error error = errors.New("auth")

type Student struct {
	UserName   string
	UserID     bson.ObjectId
	Identifier string
}

func (s *Student) Study(name string, content string) error {
	var c CourseDetail
	var brief CourseBrief

	c.ID = bson.NewObjectId()
	c.Name = name
	c.Content = content

	brief.ID = c.ID
	brief.Name = name

	u := s.SelectUser()
	if u == nil {
		return find_error
	}

	err := course_collection.Insert(&c)
	if err != nil {
		log.Println("INSERT_COURSE:", err)
		return err
	}

	cpfunc := func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			return r
		}
		return ' '
	}

	sWord := strings.Map(cpfunc, content)
	arrWord := strings.Split(sWord, " ")
	for _, v := range arrWord {
		// invalid word
		if len(v) <= 2 {
			continue
		}

		// familiar word
		if u.Familiars[strings.ToLower(v)] {
			continue
		}

		// learned word
		_, present := u.Words[strings.ToLower(v)]
		if present {
			continue
		}

		var w Word
		w.Name = v
		w.CourseName = c.Name
		w.Status = "unfamiliar"

		u.Words[strings.ToLower(v)] = w
	}

	u.Courses = append(u.Courses, brief)

	return s.UpdateUser(u)
}

func (s *Student) ReviewLesson(name string) []string {
	var res []string

	u := s.SelectUser()
	if u == nil {
		return res
	}

	for _, v := range u.Words {
		if v.CourseName != name {
			continue
		}
		res = append(res, v.Name)
	}

	sort.Strings(res)
	return res
}

func (s *Student) DeleteLesson(name string, id string) error {
	err := course_collection.RemoveId(bson.ObjectIdHex(id))
	if err != nil {
		log.Println("course delete", err)
	}
	u := s.SelectUser()
	if u == nil {
		return login_error
	}

	index := -1
	for k, v := range u.Courses {
		if v.ID == bson.ObjectIdHex(id) {
			log.Println("delete article:", id)
			index = k
		}
	}
	if index != -1 {
		u.Courses = append(u.Courses[:index], u.Courses[index+1:]...)
	}

	res := []string{}
	for k, v := range u.Words {
		if v.CourseName == name {
			res = append(res, k)
		}
	}
	for _, v := range res {
		delete(u.Words, v)
	}

	err = s.UpdateUser(u)
	if err != nil {
		return err
	}

	return nil
}

func (s *Student) SelectUser() *User {
	var u User
	err := user_collection.FindId(s.UserID).One(&u)
	if err != nil {
		log.Println("QUERY_USER_ERROR:", err)
		return nil
	}
	return &u
}

func (s *Student) UpdateUser(u *User) error {
	log.Println("UpdateUser", s.UserID, u.ID)
	err := user_collection.UpdateId(s.UserID, u)
	if err != nil {
		log.Println("UPDATE_USER_ERROR:", err)
		return err
	}
	return nil
}

func (s *Student) AddFamiliar(content string) error {
	u := s.SelectUser()
	if u == nil {
		return find_error
	}

	cpfunc := func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			return r
		}
		return ' '
	}

	sWord := strings.Map(cpfunc, content)
	arrWord := strings.Split(sWord, " ")
	for _, v := range arrWord {
		if len(v) <= 2 {
			continue
		}

		u.Familiars[strings.ToLower(v)] = true
		delete(u.Words, v)
	}

	return s.UpdateUser(u)
}

func (s *Student) ChangeWordStatus(name string, status string) error {
	name = strings.ToLower(name)

	u := s.SelectUser()
	if u == nil {
		return find_error
	}

	if status == "familiar" {
		_, present := u.Words[name]
		if !present {
			return nil
		}
		delete(u.Words, name)
		u.Familiars[strings.ToLower(name)] = true
	} else {
		v, present := u.Words[name]
		if present {
			v.Status = status
			u.Words[name] = v
		}
	}

	return s.UpdateUser(u)
}

func (s *Student) Login(name string, pass string) error {
	var u User
	err := user_collection.Find(bson.M{"name": name}).One(&u)
	if err != nil {
		log.Println("QUERY_USER_ERROR:", err)
		return login_error
	}

	log.Println("name", u)

	if pass != u.Password {
		return login_error
	}

	s.UserName = u.Name
	s.UserID = u.ID
	s.Identifier = strconv.Itoa(rand.Int())

	_, err = session_collection.Upsert(bson.M{"username": name}, s)
	if err != nil {
		log.Println("SESSION_UPSERT_ERROR:", err)
		return login_error
	}

	log.Println("up", s)

	return nil
}

/*
func (s * Student) CreateUser(name string, pass string) error {
	u := s.SelectUser()
	if u.Name != "fancymore" {
		return auth_error
	}

	var nUser User
	nUser.ID = bson.NewObjectId()
	nUser.Familiars = map[string]bool {}
	nUser.Name = name
	nUser.Password = pass
	nUser.Words = map[string]Word {}

	err := user_collection.Insert(&nUser)
	if err != nil {
		log.Println("CREATE_USER:", err)
		return err
	}
	return nil
}
*/
