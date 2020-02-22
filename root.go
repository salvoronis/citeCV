package main

import(
  "fmt"
  "net/http"
  "html/template"
)


type RootP struct {
  Username, Class string
}

func root(w http.ResponseWriter, r *http.Request){
  session, err := store.Get(r, "cookie-name")
  if err != nil {
    fmt.Println(err)
  }
  if session.Values["authenticated"] == true {
    person := session.Values["user"].(*pupil)
    t := template.Must(template.ParseFiles("pages/index.html"))
    t.Execute(w, &RootP{Username: person.Username, Class: person.Class})
  } else {
    http.Redirect(w,r, "/login", 301)
  }
}
