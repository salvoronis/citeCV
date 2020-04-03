package main

import (
  "fmt"
  "net/http"
  "time"
  "text/template"
  "bytes"
  //"os"
  "github.com/lib/pq"
)

type news struct{
  Id int
  Title, Subtitle, Body, Author string
  Date time.Time
  Tags []string
}

func newss(w http.ResponseWriter, r *http.Request){
  t := template.New("news")
  var newOne = `
    <div class="blog-card">
      <div class="meta">
        <div class="photo" style="background-image: url(http://localhost:8080/img/salvoroni.jpg)"></div>
        <ul class="details">
          <li class="author"><a href="#">{{.Author}}</a></li>
          <li class="date">{{.Date}}</li>
          <li class="tags">
            <ul>
              {{range .Tags}}
              <li><a href="#">{{ . }}</a></li>
              {{end}}
            </ul>
          </li>
        </ul>
      </div>
      <div class="description">
        <h1>{{.Title}}</h1>
        <h2>{{.Subtitle}}</h2>
        <p>{{.Body}}</p>
        <p class="read-more">
          <a href="#">Read More</a>
        </p>
      </div>
    </div>
  `
  t.Parse(newOne)

  if r.Method == "POST"{
    //var result string
    result := new(bytes.Buffer)
    news := news{}
    rows, err := db.Query("select * from news offset 0 rows fetch next 10 rows only;")
    if err != nil {
      fmt.Println("can not load rows")
    }
    for rows.Next(){
      err := rows.Scan(&news.Id, &news.Title, &news.Subtitle, &news.Body, &news.Author, &news.Date, pq.Array(&news.Tags))
      t.Execute(result, news)
      if err != nil{
        fmt.Println(err)
      }
    }
    fmt.Fprint(w,result)
    //select * from school_users offset 0 rows fetch next 10 rows only;
  }
}
