package main

import(
  "fmt"
  "net/http"
  _ "github.com/lib/pq"
  "html/template"
  "net/smtp"
  "log"
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
      err := rows.Scan(&user.username, &user.password, &user.mail, &user.index)
      if err != nil{
        fmt.Println("can't load pupils")
      }

      if (user.username == r.FormValue("username") || user.mail == r.FormValue("mail")){
        t.Execute(w,&Page{Bad: "username or mail is already use",})
        return
      }
    }

    _, err = db.Exec("insert into school_users (username, mail, password, index) values ('"+r.FormValue("username")+"','"+r.FormValue("mail")+"','"+GetMd5(r.FormValue("password"))+"','"+GetMd5(r.FormValue("username"))+"');")
    //тут нужно подтвердить по почте
    go sendMail(r.FormValue("mail"), GetMd5(r.FormValue("username")))
    if err != nil {
      fmt.Println("can not insert into the table")
    }
    http.Redirect(w,r, "/login", 301)

  }
}

func sendMail(to string, index string){
  // somecitee@gmail.com pidorasi
  /*auth := smtp.PlainAuth("", "somecitee@gmail.com", "pidorasi", "smtp.gmail.com")
  msg := []byte("To: "+mail+"\r\n" +
		"Subject: discount Gophers!\r\n" +
		"\r\n" +
		"This is the email body.\r\n")

  err := smtp.SendMail("smtp.gmail.com:465", auth, "somecitee@gmail.com", []string{mail}, msg)
	if err != nil {
		fmt.Println(err)
	}*/
  from := "somecitee@gmail.com"
	pass := "pidorasi"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there\n\n" +
    "To confirm your email follow the link localhost:8080/index?index=" + index

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
}
