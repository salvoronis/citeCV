package main

import(
  "fmt"
  "net/http"
  _ "github.com/lib/pq"
  "html/template"
  "net/smtp"
  "net"
  "crypto/tls"
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
    go sendMail(r.FormValue("mail"))
    if err != nil {
      fmt.Println("can not insert into the table")
    }
    http.Redirect(w,r, "/login", 301)

  }
}

func sendMail(to string){
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
  subj := "This is the email subject"
  body := "This is an example body.\n With two lines."

  // Setup headers
  headers := make(map[string]string)
  headers["From"] = from
  headers["To"] = to
  headers["Subject"] = subj

  // Setup message
  message := ""
  for k,v := range headers {
      message += fmt.Sprintf("%s: %s\r\n", k, v)
  }
  message += "\r\n" + body

  // Connect to the SMTP Server
  servername := "smtp.gmail.com:465"

  host, _, _ := net.SplitHostPort(servername)

  auth := smtp.PlainAuth("","somecitee@gmail.com", "pidorasi", host)

  // TLS config
  tlsconfig := &tls.Config {
      InsecureSkipVerify: true,
      ServerName: host,
  }

  conn, err := tls.Dial("tcp", servername, tlsconfig)
  if err != nil {
      log.Panic(err)
  }

  c, err := smtp.NewClient(conn, host)
  if err != nil {
      log.Panic(err)
  }

  // Auth
  if err = c.Auth(auth); err != nil {
      log.Panic(err)
  }

  // To && From
  if err = c.Mail(from); err != nil {
      log.Panic(err)
  }

  if err = c.Rcpt(to); err != nil {
      log.Panic(err)
  }

  // Data
  w, err := c.Data()
  if err != nil {
      log.Panic(err)
  }

  _, err = w.Write([]byte(message))
  if err != nil {
      log.Panic(err)
  }

  err = w.Close()
  if err != nil {
      log.Panic(err)
  }

  c.Quit()
}
