package main

import (
    "fmt"
    "net/http"
    "database/sql"
    _ "github.com/lib/pq"
    "html/template"

    "github.com/gorilla/sessions"
)

type pupil struct{
  username string
  password string
  mail string
  session string
}

var (
    // key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
    key = []byte("super-secret-key")
    store = sessions.NewCookieStore(key)

    connStr = "host=localhost port=5432 user=postgres password=qwerty dbname=postgres sslmode=disable"
    db, err = sql.Open("postgres", connStr)
)

func secret(w http.ResponseWriter, r *http.Request) {
    session, _ := store.Get(r, "cookie-name")

    // Check if user is authenticated
    if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
        http.Error(w, "Forbidden", http.StatusForbidden)
        return
    }

    // Print secret message
    fmt.Fprintln(w, "The cake is a lie!")
}

func login(w http.ResponseWriter, r *http.Request) {
  if r.Method == "POST"{
    user := pupil{}

    session, _ := store.Get(r, "cookie-name")

    err := r.ParseMultipartForm(1024)
    if err != nil {
      fmt.Println(err)
    }
    //fmt.Println(r.FormValue("password"))

    //_, err = db.Exec("insert into school_users (username, mail, password, session) values ('penis','hui','hui','hui');")
    rows, err := db.Query("select * from school_users")
    if err != nil {
      fmt.Println("can not load rows")
    }
    for rows.Next(){
      err := rows.Scan(&user.username, &user.password, &user.mail, &user.session)
      if err != nil{
        fmt.Println("can't load pupils")
      }

      //fmt.Println(user.username)
      if (user.username == r.FormValue("username") && user.password == r.FormValue("password")){
        session.Values["authenticated"] = true
        session.Save(r, w)
        fmt.Fprint(w, "ok")
        break
      }
    }
    //fmt.Println(user.username)

    session.Save(r, w)
  } else if r.Method == "GET"{
    t := template.Must(template.ParseFiles("pages/login.html"))
    t.Execute(w, "")
  }
}

func logout(w http.ResponseWriter, r *http.Request) {
    session, _ := store.Get(r, "cookie-name")

    // Revoke users authentication
    session.Values["authenticated"] = false
    session.Save(r, w)
}

func main() {
    http.HandleFunc("/secret", secret)
    http.HandleFunc("/login", login)
    http.HandleFunc("/logout", logout)
    http.Handle("/js/",http.StripPrefix("/js/", http.FileServer(http.Dir("./scripts"))))

    http.ListenAndServe(":8080", nil)
}
