package main

import(
  "net/http"
  "html/template"
  "fmt"
  "math/rand"
  "time"
)

type User struct{
  Mail string
}

type PageRecovery struct{
  Error string
}

func recovery(w http.ResponseWriter, r *http.Request){
  if r.Method == "GET" {
      t := template.Must(template.ParseFiles("pages/recovery.html"))
      t.Execute(w,"")
  } else if r.Method == "POST" {
    user := User{}
    username := r.FormValue("username")
    rows, err := db.Query("select mail from school_users where username = '"+username+"'")
    if err != nil {
      fmt.Println(err)
    }
    for rows.Next(){
      err = rows.Scan(&user.Mail)
      if err != nil{
        fmt.Println(err)
      }
    }
    if user.Mail == "" {
      t := template.Must(template.ParseFiles("pages/recovery.html"))
      t.Execute(w,&PageRecovery{Error: "No such username"})
    } else {
      password := randomPass()
      _, err := db.Query("update school_users set password = '"+GetMd5(password)+"' where username = '"+username+"';")
      if err != nil {
        fmt.Println(err)
      }
      go sendMail(user.Mail, "Password recover", "Your new password is "+password)
      http.Redirect(w,r, "/login", 301)
    }
  }
}

func randomInt() int{
  rand.Seed(time.Now().UTC().UnixNano())
  bytes := rand.Intn(26)
  return bytes
}

func randomPass() string {
  var password string
  for i := 0;i < 30;i++ {
    rand := randomInt()
    Char := string('a' + byte(rand))
    password += Char
  }
  return password
}
