package main

import(
  "net/http"
  "html/template"
  "fmt"
  "os"
  "io"
)

type Profile struct {
  Username, Mail, File, Class, EditInfo string
  Confirm, Edit bool
}

func profile(w http.ResponseWriter, r *http.Request){
  session, _ := store.Get(r, "cookie-name")
  if r.Method == "GET"{
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
      err := rows.Scan(&object.Username, &object.Mail, &object.Password, &object.Index, &object.Class)
      if err != nil{
        fmt.Println("can't load pupils")
      }
    }
    if object.Index == "ok" {
      conf = true
    }
    Sfile, err := os.Open("./images/"+object.Username+".jpg")
    Sfile.Close();
    if err != nil {
      file = "_default"
    } else {
      file = object.Username
    }
    edit := object.Username == session.Values["user"].(*pupil).Username

    if forProfile != "" {
      t.Execute(w, &Profile{Username: object.Username, Mail: object.Mail, File: file, Confirm: conf, Edit: edit, Class: object.Class, EditInfo: forProfile})
      forProfile = ""
    } else {
      t.Execute(w, &Profile{Username: object.Username, Mail: object.Mail, File: file, Confirm: conf, Edit: edit, Class: object.Class})
    }

  } else {
    f, _, err := r.FormFile("file")
    if err != nil {
      fmt.Println("image shit")
    }
    defer f.Close()
    filename := "./images/" + session.Values["user"].(*pupil).Username + ".jpg"
    out, err := os.Create(filename)
    if err != nil {
      fmt.Println("another shit")
    }
    defer out.Close()
    io.Copy(out, f)
    fmt.Fprint(w,"ok");
  }
}
