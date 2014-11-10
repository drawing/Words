package main

import (
	"./basic"
)

import (
	"encoding/json"
	"log"
	"net/http"
)

func articles(w http.ResponseWriter, r *http.Request) {
	res := map[string]interface{}{}

	student := basic.GetLoginStudent(w, r)
	if student == nil {
		log.Println("articles get student error")
		res["code"] = 2
		res["target"] = "/running/login.html"
	} else {
		u := student.SelectUser()
		if u == nil {
			log.Println("student get user error")
			res["code"] = 3
			res["target"] = "/running/login.html"
		} else {
			res["code"] = 0
			res["articles"] = u.Courses
			res["username"] = u.Name
			res["familiar_num"] = len(u.Familiars)
		}
	}

	data, err := json.Marshal(res)
	if err != nil {
		log.Println("Marshal", err)
	}
	w.Write(data)

	log.Println("LISP_RESP", string(data))
}

func login(w http.ResponseWriter, r *http.Request) {
	res := map[string]interface{}{}

	user := r.FormValue("username")
	pass := r.FormValue("password")

	var student basic.Student
	err := student.Login(user, pass)
	if err != nil {
		res["code"] = 1
	} else {
		var cookie http.Cookie
		cookie.Path = "/"
		cookie.Name = "username"
		cookie.Value = student.UserName
		w.Header().Add("Set-Cookie", cookie.String())
		cookie.Name = "identifier"
		cookie.Value = student.Identifier
		w.Header().Add("Set-Cookie", cookie.String())

		res["code"] = 0
		res["target"] = "/running/index.html"
	}

	data, err := json.Marshal(res)
	w.Write(data)

	log.Println("LOGIN_RESP", string(data))
}

func add_user(w http.ResponseWriter, r *http.Request) {
	res := map[string]interface{}{}

	name := r.FormValue("username")
	email := r.FormValue("email")
	pass := r.FormValue("password")

	err := basic.AddUser(name, email, pass)
	if err != nil {
		res["code"] = 1
		res["message"] = err.Error()
	} else {
		login(w, r)
		return
	}

	data, err := json.Marshal(res)
	if err != nil {
		log.Println("Marshal", err)
	}
	w.Write(data)
}

func add_article(w http.ResponseWriter, r *http.Request) {
	res := map[string]interface{}{}

	student := basic.GetLoginStudent(w, r)
	if student == nil {
		log.Println("articles get student error")
		res["code"] = 2
		res["target"] = "/running/login.html"
	} else {
		name := r.FormValue("title")
		content := r.FormValue("content")

		if name == "" || content == "" {
			res["code"] = 3
		} else {
			err := student.Study(name, content)
			if err != nil {
				res["code"] = 4
			} else {
				res["code"] = 0
			}
		}
	}

	data, err := json.Marshal(res)
	if err != nil {
		log.Println("Marshal", err)
	}
	w.Write(data)

	log.Println("ADD_RESP", string(data))
}

func vocabulary(w http.ResponseWriter, r *http.Request) {
	res := map[string]interface{}{}

	student := basic.GetLoginStudent(w, r)
	if student == nil {
		log.Println("articles get student error")
		res["code"] = 2
		res["target"] = "/running/login.html"
	} else {
		name := r.FormValue("course")

		if name == "" {
			res["code"] = 3
		} else {
			res["code"] = 0
			res["vocabulary"] = student.ReviewLesson(name)
		}
	}

	data, err := json.Marshal(res)
	if err != nil {
		log.Println("Marshal", err)
	}
	w.Write(data)

	log.Println("vocabulary_RESP", string(data))
}

func article_detail(w http.ResponseWriter, r *http.Request) {
	res := map[string]interface{}{}

	student := basic.GetLoginStudent(w, r)
	if student == nil {
		log.Println("articles get student error")
		res["code"] = 2
		res["target"] = "/running/login.html"
	} else {
		name := r.FormValue("course")
		id := r.FormValue("id")
		v := basic.GetArticleByID(id)

		if name == "" || id == "" || v == nil {
			res["code"] = 3
		} else {
			res["code"] = 0
			res["course"] = v
		}
	}

	data, err := json.Marshal(res)
	if err != nil {
		log.Println("Marshal", err)
	}
	w.Write(data)

	log.Println("detail_RESP", string(data))
}

func familiar_word(w http.ResponseWriter, r *http.Request) {
	res := map[string]interface{}{}

	student := basic.GetLoginStudent(w, r)
	if student == nil {
		log.Println("articles get student error")
		res["code"] = 2
		res["target"] = "/running/login.html"
	} else {
		word := r.FormValue("word")

		if word == "" {
			res["code"] = 3
		} else {
			student.ChangeWordStatus(word, "familiar")
			res["code"] = 0
		}
	}

	data, err := json.Marshal(res)
	if err != nil {
		log.Println("Marshal", err)
	}
	w.Write(data)

	log.Println("familiar_RESP", string(data))
}

func delete_article(w http.ResponseWriter, r *http.Request) {
	res := map[string]interface{}{}

	student := basic.GetLoginStudent(w, r)
	if student == nil {
		log.Println("articles get student error")
		res["code"] = 2
		res["target"] = "/running/login.html"
	} else {
		name := r.FormValue("course")
		id := r.FormValue("id")

		if name == "" || id == "" {
			res["code"] = 3
		} else {
			err := student.DeleteLesson(name, id)
			if err != nil {
				res["code"] = 4
			} else {
				res["code"] = 0
			}
		}
	}

	data, err := json.Marshal(res)
	if err != nil {
		log.Println("Marshal", err)
	}
	w.Write(data)

	log.Println("delete_RESP", string(data))
}

func add_familiar(w http.ResponseWriter, r *http.Request) {
	res := map[string]interface{}{}

	student := basic.GetLoginStudent(w, r)
	if student == nil {
		log.Println("articles get student error")
		res["code"] = 2
		res["target"] = "/running/login.html"
	} else {
		content := r.FormValue("content")
		err := student.AddFamiliar(content)
		if err != nil {
			res["code"] = 4
		} else {
			res["code"] = 0
		}
	}

	data, err := json.Marshal(res)
	if err != nil {
		log.Println("Marshal", err)
	}
	w.Write(data)

	log.Println("add_familiar_RESP", string(data))
}

func handler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/running/index.html", 302)
}

func main() {
	basic.InitStorage("paulo.mongohq.com:10061", "LearnEng",
		"user", "password")

	http.HandleFunc("/", handler)

	http.HandleFunc("/login", login)
	http.HandleFunc("/articles", articles)
	http.HandleFunc("/add_article", add_article)
	http.HandleFunc("/vocabulary", vocabulary)
	http.HandleFunc("/familiar_word", familiar_word)
	http.HandleFunc("/article_detail", article_detail)
	http.HandleFunc("/delete_article", delete_article)
	http.HandleFunc("/add_user", add_user)
	http.HandleFunc("/add_familiar", add_familiar)

	http.HandleFunc("/translate", basic.ActionTranslate)

	var thesaurus basic.Thesaurus
	thesaurus.LoadFromDB()
	thesaurus.SetHandleFunc()

	// log.Fatal(thesaurus.LoadFromTxt("六级词汇", "/home/cppbreak/workspace/xx/cet6.txt"))

	http.Handle("/running/", http.StripPrefix("/running/", http.FileServer(http.Dir("running"))))

	log.Println("running...")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
