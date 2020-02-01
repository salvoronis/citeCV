package main

import(
  "fmt"
  "net/http"
  _ "github.com/lib/pq"
  "html/template"
)

func register(w http.ResponseWriter, r *http.Request){
  if r.Method == "GET"{
    t := template.Must(template.ParseFiles("pages/register.html"))
    t.Execute(w,"")
  }else if r.Method == "POST" {
    t := template.Must(template.ParseFiles("pages/register.html"))
    type Page struct{
      Bad string
    }
    user := pupil{}

    err := r.ParseMultipartForm(1024)
    if err != nil {
      fmt.Println("can not parse form")
    }

    rows, err := db.Query("select * from school_users")
    if err != nil {
      fmt.Println("can not load rows")
    }
    for rows.Next(){
      err := rows.Scan(&user.username, &user.password, &user.mail, &user.session)
      if err != nil{
        fmt.Println("can't load pupils")
      }

      if (user.username == r.FormValue("username") || user.mail == r.FormValue("mail")){
        t.Execute(w,&Page{Bad: "username or password is already use",})
        return
      }
    }

    _, err = db.Exec("insert into school_users (username, mail, password, session) values ('"+r.FormValue("username")+"','"+r.FormValue("mail")+"','"+GetMd5(r.FormValue("password"))+"','hui');")
    if err != nil {
      fmt.Println("can not insert into table")
    }
    http.Redirect(w,r, "/login", 301)

  }
}
