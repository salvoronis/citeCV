package main

import(
  "net/http"
  "fmt"
  "strings"
)

type Pass struct{
  Password string
}

var forProfile string

func editUser(w http.ResponseWriter, r *http.Request){
  if r.Method == "POST" {
    session, _ := store.Get(r, "cookie-name")

    hash := &Pass{}
    rows, err := db.Query("select password from school_users where username = '"+session.Values["user"].(*pupil).Username+"';")
    if err != nil {
      fmt.Println("can not load rows")
    }
    for rows.Next(){
      err := rows.Scan(&hash.Password)
      if err != nil{
        fmt.Println("can't load pupils")
      }
    }

    if GetMd5(r.FormValue("old")) == hash.Password {
      var newdata string = "update school_users set "
      r.ParseForm()

      for key, val := range r.Form {
        if val[0] == "" || key == "old" {
          continue
        } else {
          if key == "password" {
              newdata += key + " = '" + GetMd5(val[0]) + "', "
          } else {
            newdata += key + " = '" + val[0] + "', "
          }
        }
      }

      newdata = strings.TrimRight(newdata,", ")
      newdata += " where username = '" + session.Values["user"].(*pupil).Username + "';"
      _, err := db.Query(newdata)
      if err != nil {
        fmt.Println("can not load rows")
      }

      user := &pupil{}
      rows, err := db.Query("select * from school_users where username = '"+r.FormValue("username")+"'")
      if err != nil {
        fmt.Println("can not load rows")
      }
      for rows.Next(){
        err := rows.Scan(&user.Username, &user.Mail, &user.Password, &user.Index, &user.Class)
        if err != nil{
          fmt.Println(err)
        }
          session.Values["user"] = user
          err = session.Save(r, w)
          if err != nil {
            fmt.Println(err)
          }
      }

      session, _ = store.Get(r, "cookie-name")
      forProfile = "ok"
      http.Redirect(w,r, "/profile?name="+session.Values["user"].(*pupil).Username, 301)
    } else {
      forProfile = "wrong password"
      http.Redirect(w,r, "/profile?name="+session.Values["user"].(*pupil).Username, 301)
    }
  }
}
