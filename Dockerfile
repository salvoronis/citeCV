FROM golang:1.14

MAINTAINER Yaremko Roman <salvoronis@gmail.com>

WORKDIR /home/salvoroni/someshittycite
COPY . .
#COPY ./pages ./pages

RUN go get github.com/gorilla/sessions github.com/gorilla/websocket github.com/lib/pq
RUN go build -o anime main.go login.go addNews.go editUser.go feed.go gettimetable.go index.go news.go profile.go recovery.go register.go root.go

ENTRYPOINT ./anime

EXPOSE 8080
