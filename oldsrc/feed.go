package main

import(
  "fmt"
  "net/http"
  "html/template"
  "github.com/lib/pq"
)

func feed(w http.ResponseWriter, r *http.Request){
  if r.Method == "GET"{
    t := template.Must(template.ParseFiles("pages/feed.html"))
    id := r.FormValue("id")
    rows, err := db.Query("select * from news where id = '"+id+"';")
    if err != nil {
      fmt.Println("can not load rows")
    }
    for rows.Next(){
      news := news{}
      err := rows.Scan(&news.Id, &news.Title, &news.Subtitle, &news.Body, &news.Author, &news.Date, pq.Array(&news.Tags), &news.Image)
      t.Execute(w, news)
      if err != nil{
        fmt.Println(err)
      }
    }
  }
}
