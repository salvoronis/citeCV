package main

import(
  "fmt"
  "net/http"
  _ "github.com/lib/pq"
  "html/template"
)

type Classes struct{
  Class string
}

type Page struct{
  Bad string
  Abc []string
}

func register(w http.ResponseWriter, r *http.Request){
  var cl []string
  classes := Classes{}
  rows, err := db.Query("select class from \"timetable\" group by class;")
  if err != nil {
    fmt.Println(err)
  }
  for rows.Next(){
    err := rows.Scan(&classes.Class)
    if err != nil {
      fmt.Println(err)
    }
    cl = append(cl, classes.Class)
  }
  if r.Method == "GET"{
    session, _ := store.Get(r, "cookie-name")
    session.Values["authenticated"] = false
    session.Save(r, w)
    t := template.Must(template.ParseFiles("pages/register.html"))
    t.Execute(w, &Page{Abc: cl})
  }else if r.Method == "POST" {
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
      err := rows.Scan(&user.Username, &user.Password, &user.Mail, &user.Index, &user.Class, &user.Role)
      if err != nil{
        fmt.Println("can't load pupils")
      }

      if (user.Username == r.FormValue("username") || user.Mail == r.FormValue("mail")){
        t := template.Must(template.ParseFiles("pages/register.html"))
        t.Execute(w, &Page{Bad: "username or mail is already use",Abc: cl})
        return
      }
    }

    index := randomPass()
    _, err = db.Exec("insert into school_users (username, mail, password, index, class) values ('"+r.FormValue("username")+"','"+r.FormValue("mail")+"','"+GetMd5(r.FormValue("password"))+"','"+GetMd5(index)+"','"+r.FormValue("class")+"');")
    if err != nil {
      fmt.Println("can not insert into the table")
    }
    go sendMail(r.FormValue("mail"), "Confirm mail" , "To confirm your email follow the link 188.120.244.137:8080/index?index="+GetMd5(index))
    http.Redirect(w,r, "/login", 301)

  }
}
