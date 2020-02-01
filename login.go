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
    t := template.Must(template.ParseFiles("pages/index.html"))
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
      err := rows.Scan(&user.username, &user.mail, &user.password, &user.session)
      if err != nil{
        fmt.Println("can't load pupils")
      }

      if (user.username == r.FormValue("username") && user.password == GetMd5(r.FormValue("password"))){
        session.Values["authenticated"] = true
        session.Save(r, w)
        t.Execute(w, &Page{Username: user.username})
        break
      }
    }

    session.Save(r, w)
  } else if r.Method == "GET"{
    t := template.Must(template.ParseFiles("pages/login.html"))
    t.Execute(w, "")
  }
}

func logout(w http.ResponseWriter, r *http.Request) {
  session, _ := store.Get(r, "cookie-name")

  session.Values["authenticated"] = false
  session.Save(r, w)
}
