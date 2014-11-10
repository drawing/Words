package basic

import (
	"encoding/json"
	"io/ioutil"
	"labix.org/v2/mgo/bson"
	"log"
	"net/http"
	"strings"
)

type Dictionary struct {
	ID      bson.ObjectId `bson:"_id"`
	Name    string
	Content []string
}

type Thesaurus struct {
	Dict []Dictionary
}

func (t *Thesaurus) LoadFromDB() error {
	err := dictionary_collection.Find(nil).All(&t.Dict)
	if err != nil {
		return err
	}

	return nil
}

func (t *Thesaurus) LoadFromTxt(name string, filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	cpfunc := func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			return r
		}
		return ' '
	}

	res := []string{}
	sWord := strings.Map(cpfunc, string(data))
	arrWord := strings.Split(sWord, " ")
	for _, v := range arrWord {
		if len(v) <= 2 {
			continue
		}

		res = append(res, strings.ToLower(v))
	}

	var dict Dictionary
	dict.Name = name
	dict.Content = res
	dict.ID = bson.NewObjectId()

	err = dictionary_collection.Insert(&dict)
	if err != nil {
		return err
	}

	return nil
}

func (t *Thesaurus) dictionary_list(w http.ResponseWriter, r *http.Request) {
	res := map[string]interface{}{}

	res["code"] = 0
	res["dict"] = t.Dict

	data, err := json.Marshal(res)
	if err != nil {
		log.Println("Marshal", err)
	}

	w.Write(data)
	// log.Println("dictionary_list_resp", string(data))
}

func (t *Thesaurus) SetHandleFunc() {
	http.HandleFunc("/dictionary_list", t.dictionary_list)
}
