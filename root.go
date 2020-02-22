package main

import(
  "fmt"
  "net/http"
  "html/template"
)

func root(w http.ResponseWriter, r *http.Request){
  session, err := store.Get(r, "cookie-name")
  if err != nil {
    fmt.Println(err)
  }
  if session.Values["authenticated"] == true {
    t := template.Must(template.ParseFiles("pages/index.html"))
    t.Execute(w, &Page{Username: session.Values["user"].(string)})
  } else {
    http.Redirect(w,r, "/login", 301)
  }
}
