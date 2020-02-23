package main

import(
  "fmt"
  "net/http"
  _ "github.com/lib/pq"
  "html/template"
)

type Page struct{
  Username string
}

func login(w http.ResponseWriter, r *http.Request) {
  if r.Method == "POST"{
    user := pupil{}

    session, _ := store.Get(r, "cookie-name")

    err := r.ParseMultipartForm(1024)
    if err != nil {
      fmt.Println(err)
    }

    rows, err := db.Query("select * from school_users")
    if err != nil {
      fmt.Println("can not load rows")
    }
    for rows.Next(){
      err := rows.Scan(&user.Username, &user.Mail, &user.Password, &user.Index, &user.Class)
      if err != nil{
        fmt.Println(err)
      }

      if (user.Username == r.FormValue("username") && user.Password == GetMd5(r.FormValue("password"))){
        session.Values["authenticated"] = true
        session.Values["user"] = &user
        session.Save(r, w)
        http.Redirect(w,r, "/", 301)
        break
      }
    }

    session.Save(r, w)
    http.Redirect(w,r, "/login", 301)
  } else if r.Method == "GET"{
    t := template.Must(template.ParseFiles("pages/login.html"))
    t.Execute(w, "")
  }
}

func logout(w http.ResponseWriter, r *http.Request) {
  session, _ := store.Get(r, "cookie-name")

  session.Values["authenticated"] = false
  session.Save(r, w)
  http.Redirect(w,r, "/login", 301)
}
