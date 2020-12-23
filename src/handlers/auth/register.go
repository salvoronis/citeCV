package auth

import (
	"bytes"
	"database/sql"
	"databaseutils"
	"fmt"
	"html/template"
	"log"
	"models"
	"net/http"
	"utils"

	"github.com/gorilla/sessions"
)

type getRegisterWrapper struct {
	Classes []models.Class
	Messages []string
}

var (
	storeR *sessions.CookieStore
	templ *template.Template
	dbReg *sql.DB
)

func init() {
	templ = template.Must(template.ParseFiles("./pages/register.html"))
	dbReg = databaseutils.GetDB()
}

func register(res http.ResponseWriter, req *http.Request){
	session, err := storeR.Get(req, "user-data")
	if err != nil {
		log.Printf("Can't get cookies \n%v\n", err)
		return
	}

	if req.Method == "GET" {
		answer := getRegisterWrapper{}
		answer.Classes = databaseutils.GetClasses()

		flashes := session.Flashes()
		session.Save(req, res)
		if len(flashes) > 0 {
			for _, val := range flashes {
				answer.Messages = append(answer.Messages, fmt.Sprintf("%v", val))
			}
		}

		var buf bytes.Buffer
		err := templ.Execute(&buf, answer)
		if err != nil {
			log.Printf("Can not execute structure to template %v\n", err)
			return
		}
		buf.WriteTo(res)
	} else if req.Method == "POST"{
		student := models.Student{
			Username: req.FormValue("nickname"),
			Name: req.FormValue("name"),
			Password: utils.GetSHA256(req.FormValue("password")),
			Class: req.FormValue("class")}
		if student.Password != utils.GetSHA256(req.FormValue("repeat")) || req.FormValue("terms-of-use") != "on" {
			session.AddFlash("Something went wrong!")
			session.Save(req, res)
			http.Redirect(res, req, "/register", 301)
		}

		databaseutils.SaveStudent(student)

		session.Values["user"] = student
		session.Values["authenticated"] = true
		session.Save(req, res)

		http.Redirect(res, req, "/", 301)
	}
}

func NewRegister(storeS *sessions.CookieStore) func(http.ResponseWriter, *http.Request) {
	storeR = storeS
	return register
}
