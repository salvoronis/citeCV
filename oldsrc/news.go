package main

import (
  "fmt"
  "net/http"
  "time"
  "text/template"
  "bytes"
  "github.com/lib/pq"
)

type news struct{
  Id int
  Title, Subtitle, Body, Author, Image string
  Date time.Time
  Tags []string
}

func newss(w http.ResponseWriter, r *http.Request){
  t := template.New("news")
  var newOne = `
    <div class="blog-card">
      <div class="meta">
        <div class="photo" style="background-image: url(/img/newsimg/{{.Image}})"></div>
        <ul class="details">
          <li class="author"><a href="/profile?name={{.Author}}">{{.Author}}</a></li>
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
          <a href="/feed?id={{.Id}}">Read More</a>
        </p>
      </div>
    </div>
  `
  t.Parse(newOne)

  if r.Method == "POST"{
    result := new(bytes.Buffer)
    rows, err := db.Query("select * from news order by id asc offset "+r.FormValue("offset")+" rows fetch next 10 rows only;")
    if err != nil {
      fmt.Println("can not load rows")
    }
    for rows.Next(){
      news := news{}
      err := rows.Scan(&news.Id, &news.Title, &news.Subtitle, &news.Body, &news.Author, &news.Date, pq.Array(&news.Tags), &news.Image)
      news.Body = cut(news.Body, 250)
      t.Execute(result, news)
      if err != nil{
        fmt.Println(err)
      }
    }
    fmt.Fprint(w,result)
  }
}

func cut(text string, limit int) string {
  runes := []rune(text)
  if len(runes) >= limit {
    return string(runes[:limit])
  }
  return text
}
