package root

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"models"

	"github.com/gorilla/sessions"
)

var (
	store *sessions.CookieStore
	templ *template.Template
)

func init(){
	templ = template.Must(template.ParseFiles("./pages/root.html"))
}

func root (res http.ResponseWriter, req *http.Request) {
	session, err := store.Get(req, "user-data")
	if err != nil {
		log.Printf("Can't get cookies %v\n", err)
	}
	if session.Values["authenticated"] == true {
		user := session.Values["user"].(models.Student)
		var buf bytes.Buffer
		err := templ.Execute(&buf, user)
		if err != nil {
			log.Printf("Can not execute template root.\n%v\n", err)
			return
		}
		buf.WriteTo(res)
	} else {
		http.Redirect(res, req, "/login", 301)
	}
}

func NewRoot(storeS *sessions.CookieStore) func(http.ResponseWriter, *http.Request) {
	store = storeS
	return root
}
