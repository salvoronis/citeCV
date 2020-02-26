package main

import (
  //"fmt"
  "net/http"
  "database/sql"
  _ "github.com/lib/pq"
  "crypto/md5"
  "encoding/hex"

  "github.com/gorilla/sessions"
  "encoding/gob"
  "net/smtp"
  "log"
)

type pupil struct{
  Username string
  Mail string
  Password string
  Index string
  Class string
}

var (
  key = []byte("super-secret-key")
  store = sessions.NewCookieStore(key)

  connStr = "host=localhost port=5432 user=postgres password=qwerty dbname=postgres sslmode=disable"
  db, err = sql.Open("postgres", connStr)
)

func init(){
  gob.Register(&pupil{})
}

func sendMail(to, subject, body string){
  from := "somecitee@gmail.com"
	pass := "pidorasi"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: "+subject+"\n\n" +
    body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
}

func GetMd5(text string) string {
  h := md5.New()
  h.Write([]byte(text))
  return hex.EncodeToString(h.Sum(nil))
}

func main() {
  http.HandleFunc("/profile", profile)
  http.HandleFunc("/login", login)
  http.HandleFunc("/logout", logout)
  http.HandleFunc("/register", register)
  http.HandleFunc("/index", index)
  http.HandleFunc("/gettimetable", gettimetable)
  http.HandleFunc("/recovery",recovery)
  http.HandleFunc("/", root)
  http.Handle("/js/",http.StripPrefix("/js/", http.FileServer(http.Dir("./scripts"))))
  http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))
  http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./images"))))

  http.ListenAndServe(":8080", nil)
}
