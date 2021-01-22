package auth

import (
	"bytes"
	"databaseutils"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

type getLoginWrapper struct {
	Messages []string
}

var (
	storeL *sessions.CookieStore
	temlp *template.Template
)

func init() {
	temlp = template.Must(template.ParseFiles("./pages/login.html"))
}

func login(res http.ResponseWriter, req *http.Request) {
	session, err := storeL.Get(req, "user-data")
	if err != nil {
		log.Printf("Can't get cookies.\n%v\n", err)
		return
	}

	if req.Method == "GET" {
		answer := getLoginWrapper{}

		flashes := session.Flashes()
		session.Save(req, res)
		if len(flashes) > 0 {
			for _, val := range flashes {
				answer.Messages = append(answer.Messages, fmt.Sprintf("%v", val))
			}
		}

		var buf bytes.Buffer
		err := temlp.Execute(&buf, answer)
		if err != nil {
			log.Printf("Can't execute login template.\n%v\n", err)
			return
		}
		buf.WriteTo(res)
	} else if req.Method == "POST" {
		if databaseutils.CheckStudent(req.FormValue("nickname"), req.FormValue("password")){
			student := databaseutils.GetStudentByNickname(req.FormValue("nickname"))

			session.Values["user"] = student
			session.Values["authenticated"] = true
			session.Save(req, res)

			http.Redirect(res, req, "/", 301)
		} else {
			session.AddFlash("Incorrect username or password")
			session.Save(req, res)
			http.Redirect(res, req, "/login", 301)
		}
	}
}

func NewLogin(storeS *sessions.CookieStore) func(http.ResponseWriter, *http.Request) {
	storeL = storeS
	return login
}
