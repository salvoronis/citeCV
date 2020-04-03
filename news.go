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
  Title, Subtitle, Body, Author string
  Date time.Time
  Tags, Images []string
}

func newss(w http.ResponseWriter, r *http.Request){
  t := template.New("news")
  var newOne = `
    <div class="blog-card">
      <div class="meta">
        <div class="photo" style="background-image: url(http://localhost:8080/img/newsimg/{{index .Images 0}})"></div>
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
    rows, err := db.Query("select * from news offset 0 rows fetch next 10 rows only;")
    if err != nil {
      fmt.Println("can not load rows")
    }
    for rows.Next(){
      news := news{}
      err := rows.Scan(&news.Id, &news.Title, &news.Subtitle, &news.Body, &news.Author, &news.Date, pq.Array(&news.Tags), pq.Array(&news.Images))
      news.Body = cut(news.Body, 250)
      t.Execute(result, news)
      if err != nil{
        fmt.Println(err)
      }
    }
    fmt.Fprint(w,result)
    //267
  }
}

func cut(text string, limit int) string {
  runes := []rune(text)
  if len(runes) >= limit {
    return string(runes[:limit])
  }
  return text
}
