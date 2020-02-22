package main

import(
  "net/http"
  "fmt"
  "strconv"
)

type timetable struct{
  subject, time, teacher, weekday string
  kyoushitsu int
}

func gettimetable(w http.ResponseWriter, r *http.Request){
  table := timetable{}
  var tabl string
  if r.Method == "GET"{
    rows, err := db.Query("select * from \"11И\"")
    if err != nil {
      fmt.Println("can not load class")
    }
    for rows.Next(){
      err := rows.Scan(&table.subject, &table.time, &table.kyoushitsu, &table.teacher, &table.weekday)
      if err != nil {
        fmt.Println("table problem")
      }
      tabl += `<tr>
        <td>`+table.subject+`</td>
        <td>`+table.time+`</td>
        <td>`+strconv.Itoa(table.kyoushitsu)+`</td>
        <td>`+table.teacher+`</td>
        <td>`+table.weekday+`</td>
      </tr>
      `
    }
    fmt.Fprint(w,tabl)
  }
}
