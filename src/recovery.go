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
    if r.FormValue("password") != "" {
      password := r.FormValue("password")
      session, err := store.Get(r, password)
      if err != nil {
        fmt.Fprint(w,"stop right there your criminal scum")
        fmt.Println(err)
      } else {
        login := session.Values["username"].(string)
        _, err = db.Query("update school_users set password = '"+GetMd5(password)+"' where username = '"+login+"';")
        if err != nil {
          fmt.Println(err)
        }
        session.Options.MaxAge = 0
        session.Save(r,w)
        fmt.Fprint(w,"your new password is "+password)
      }
    } else {
      t := template.Must(template.ParseFiles("pages/recovery.html"))
      t.Execute(w,"")
    }
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
      go sendMail(user.Mail, "Password recover", "This message is enable only for 30 minutes\nto recover you password follow the link 188.120.244.137:8080/recovery?password="+password)
      session, err := store.Get(r, password)
      if err != nil {
        fmt.Println(err)
      }
      session.Values["username"] = username
      session.Options.MaxAge = 1800 //30 minutes
      session.Save(r,w)
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
  for i := 0;i < 15;i++ {
    rand := randomInt()
    Char := string('a' + byte(rand))
    password += Char
  }
  return password
}
