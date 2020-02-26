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
  if r.Method == "GET"{
    session, _ := store.Get(r, "cookie-name")
    session.Values["authenticated"] = false
    session.Save(r, w)
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
      err := rows.Scan(&user.Username, &user.Password, &user.Mail, &user.Index, &user.Class)
      if err != nil{
        fmt.Println("can't load pupils")
      }

      if (user.Username == r.FormValue("username") || user.Mail == r.FormValue("mail")){
        t := template.Must(template.ParseFiles("pages/register.html"))
        t.Execute(w, &Page{Bad: "username or mail is already use",})
        return
      }
    }

    _, err = db.Exec("insert into school_users (username, mail, password, index, class) values ('"+r.FormValue("username")+"','"+r.FormValue("mail")+"','"+GetMd5(r.FormValue("password"))+"','"+GetMd5(r.FormValue("username"))+"','"+r.FormValue("class")+"');")
    if err != nil {
      fmt.Println("can not insert into the table")
    }
    go sendMail(r.FormValue("mail"), "Confirm mail" , "To confirm your email follow the link 188.120.244.137:8080/index?index="+GetMd5(r.FormValue("username")))
    http.Redirect(w,r, "/login", 301)

  }
}

/*func sendMail(to string, index string){
  from := "somecitee@gmail.com"
	pass := "pidorasi"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Confirm mail\n\n" +
    "To confirm your email follow the link 188.120.244.137:8080/index?index=" + index

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
}*/
