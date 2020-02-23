package main

import(
  "net/http"
  "fmt"
)

func index(w http.ResponseWriter, r *http.Request){
  user := pupil{}
  index := r.FormValue("index")

  if index != ""{

    rows, err := db.Query("select * from school_users")
    if err != nil {
      fmt.Println("can not load rows")
    }
    for rows.Next(){
      err := rows.Scan(&user.Username, &user.Password, &user.Mail, &user.Index, &user.Class)
      if err != nil{
        fmt.Println("can't load pupils")
      }

      if (user.Index == index){
        _, err = db.Exec("update school_users set index = 'ok' where username = '"+user.Username+"';")
        fmt.Fprint(w, "Подтверждено")
        if err != nil {
          fmt.Println("can not insert into the table")
        }
        return
      }
    }
    fmt.Fprint(w, "Не найден такой пользователь")
  } else {
    fmt.Fprint(w, "Нет индекса")
  }
}
