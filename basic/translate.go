package basic

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Sentence struct {
	Orig  string `xml:"orig"`
	Trans string `xml:"trans"`
}

type DictElem struct {
	Key         string     `xml:"key"`
	Phonetic    []string   `xml:"ps"`
	Pron        []string   `xml:"pron"`
	Pos         []string   `xml:"pos"`
	Acceptation []string   `xml:"acceptation"`
	Sent        []Sentence `xml:"sent"`
}

type Pron struct {
	Phonetic string
	Pron     string
}
type Explan struct {
	Pos         string
	Acceptation string
}

type JsonDict struct {
	Key     string
	Prons   []Pron
	Explans []Explan
	Sent    []Sentence
}

func ActionTranslate(w http.ResponseWriter, r *http.Request) {
	res := map[string]interface{}{}

	r.ParseForm()
	word := r.FormValue("word")

	if word == "" {
		res["code"] = 1
	} else {
		var v DictElem
		resp, err := http.Get("http://dict-co.iciba.com/api/dictionary.php?key=D2E4A6A15C9DC43E214171AC0C6339B0&&w=" + strings.ToLower(word))
		if err != nil {
			log.Println("Get", err)
		} else {
			robots, err := ioutil.ReadAll(resp.Body)
			resp.Body.Close()

			err = xml.Unmarshal(robots, &v)
			if err != nil {
				res["code"] = 2
			} else {
				res["code"] = 0
				var j JsonDict
				j.Key = v.Key
				for ik, iv := range v.Phonetic {
					var p Pron
					p.Phonetic = iv
					if ik < len(v.Pron) {
						p.Pron = v.Pron[ik]
					}
					j.Prons = append(j.Prons, p)
				}
				for ik, iv := range v.Pos {
					var p Explan
					p.Pos = iv
					if ik < len(v.Acceptation) {
						p.Acceptation = v.Acceptation[ik]
					}
					j.Explans = append(j.Explans, p)
				}
				j.Sent = v.Sent
				res["dict"] = j
			}
		}
	}

	data, err := json.Marshal(res)
	if err != nil {
		log.Println("Marshal", err)
	}
	w.Write(data)

	log.Println("dict_RESP", string(data))
}
