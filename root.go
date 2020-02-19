package main

import(
  //"fmt"
  "net/http"
  "html/template"
)

func root(w http.ResponseWriter, r *http.Request){
  session, _ := store.Get(r, "cookie-name")
  //user := pupil{}
  if session.Values["authenticated"] == true {
    t := template.Must(template.ParseFiles("pages/index.html"))
    /*user = *///fmt.Println(session.Values["user"])
    t.Execute(w, &Page{Username: session.Values["user"].(string)})
  } else {
    http.Redirect(w,r, "/login", 301)
  }
}
