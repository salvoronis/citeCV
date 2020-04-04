package main

import(
  "net/http"
  "html/template"
  "io"
  "fmt"
  "os"
  "mime/multipart"
  "strings"
  "time"
)

type newsPage struct {
  Username string
}

func addNews (w http.ResponseWriter, r *http.Request){
  session, _ := store.Get(r, "cookie-name")
  if session.Values["authenticated"] == true && session.Values["user"].(*pupil).Role == "teacher" {
    username := newsPage{}
    username.Username = session.Values["user"].(*pupil).Username
    if r.Method == "GET" {
      t := template.Must(template.ParseFiles("pages/addNews.html"))
      t.Execute(w,username)
    } else if r.Method == "POST" {
      err := r.ParseMultipartForm(16 << 20)
      checkErr(err, w)
      data := r.MultipartForm
      files := data.File["images"]
      images := getImg(files, w)
      imgsql := strings.Join(images, "','")
      tagsql := "'"+strings.Join(strings.Split(r.FormValue("tags"), " "), "','")+"'"
      sql := "insert into news(title, subtitle, body, author, date, tags, images) values ('"+r.FormValue("title")+"', '"+r.FormValue("subtitle")+"', '"+r.FormValue("body")+"', '"+username.Username+"', '"+time.Now().Format(time.RFC3339)+"', ARRAY ["+tagsql+"], '"+imgsql+"');"
      _, err = db.Exec(sql)
      if err != nil {
        fmt.Println(err)
      }
      http.Redirect(w,r, "/", 301)
    }
  }
}

func checkErr(err error, w http.ResponseWriter){
  if err != nil {
    fmt.Fprint(w, err)
    panic(err)
  }
}

func getImg(files []*multipart.FileHeader, w http.ResponseWriter) (images []string){
  for _, fh := range files {
    f, err := fh.Open()
    defer f.Close()
    checkErr(err, w)
    out, err := os.Create("./images/newsimg/"+fh.Filename)
    defer out.Close()
    checkErr(err, w)
    _, err = io.Copy(out, f)
    checkErr(err, w)
    images = append(images, fh.Filename)
  }
  return
}
