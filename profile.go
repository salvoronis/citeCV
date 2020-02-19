package main

import(
  "net/http"
  "html/template"
  "fmt"
  "os"
)

type Profile struct {
  Username, Mail, File string
  Confirm bool
}

func profile(w http.ResponseWriter, r *http.Request){
  object := pupil{}
  var file string
  var conf bool = false
  t := template.Must(template.ParseFiles("pages/profile.html"))
  name := r.FormValue("name")

  rows, err := db.Query("select * from school_users where username = '"+name+"';")
  if err != nil {
    fmt.Println("can not load rows")
  }
  for rows.Next(){
    err := rows.Scan(&object.username, &object.mail, &object.password, &object.index)
    if err != nil{
      fmt.Println("can't load pupils")
    }
  }
  if object.index == "" {
    conf = true
  }
  Sfile, err := os.Open("./images/"+object.username+".jpg")
  Sfile.Close();
  if err != nil {
    file = "_default"
  } else {
    file = object.username
  }

  t.Execute(w, &Profile{Username: object.username, Mail: object.mail, File: file, Confirm: conf})
}
