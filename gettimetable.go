package main

import(
  "net/http"
  "fmt"
  "strconv"
)

type timetable struct{
  subject, time, teacher, weekday, class string
  kyoushitsu int
}
type grade struct{
  Grade int
}

func gettimetable(w http.ResponseWriter, r *http.Request){
  table := timetable{}
  session, _ := store.Get(r, "cookie-name")
  student := session.Values["user"].(*pupil)
  var tabl string
  if r.Method == "GET"{
    week := r.FormValue("week")
    fmt.Println(week)
    rows, err := db.Query("select * from \"timetable\" where class='"+student.Class+"';")
    if err != nil {
      fmt.Println("can not load class")
    }
    for rows.Next(){
      mark := &grade{}
      err := rows.Scan(&table.subject, &table.time, &table.kyoushitsu, &table.teacher, &table.weekday, &table.class)
      if err != nil {
        fmt.Println(err)
      }
      gradeRow, err := db.Query("select grade from grades where username = '"+student.Username+"' and week = "+week+" and subject = '"+table.subject+"' and time = '"+table.time+"';")
      if err != nil {
        fmt.Println("can not load class")
      } else{
        for gradeRow.Next(){
          err := gradeRow.Scan(&mark.Grade)
          fmt.Println(err)
        }
      }
      var result string = ""
      gr := strconv.Itoa(mark.Grade)
      if gr != "0"{
        result = gr
      }
      tabl += `<tr>
        <td>`+table.subject+`</td>
        <td>`+table.time+`</td>
        <td>`+strconv.Itoa(table.kyoushitsu)+`</td>
        <td>`+table.teacher+`</td>
        <td>`+table.weekday+`</td>
        <td>`+result+`</td>
      </tr>
      `
    }
    fmt.Fprint(w,tabl)
  }
}
