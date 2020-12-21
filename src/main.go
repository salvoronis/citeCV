package main

import (
  "net/http"
  "database/sql"
  _ "github.com/lib/pq"
  "crypto/md5"
  "encoding/hex"

  "github.com/gorilla/sessions"
  "encoding/gob"
  "net/smtp"
  "log"

  "github.com/gorilla/websocket"
  "time"
  "bytes"
)

type pupil struct{
  Username string
  Mail string
  Password string
  Index string
  Class string
  Role string
}

var (
  key = []byte("super-secret-key")
  store = sessions.NewCookieStore(key)

  connStr = "host=localhost port=5432 user=postgres password=qwerty dbname=postgres sslmode=disable"
  db, err = sql.Open("postgres", connStr)
)

func init(){
  gob.Register(&pupil{})
}

func sendMail(to, subject, body string){
  from := "somecitee@gmail.com"
	pass := "pidorasi"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: "+subject+"\n\n" +
    body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
}

func GetMd5(text string) string {
  h := md5.New()
  h.Write([]byte(text))
  return hex.EncodeToString(h.Sum(nil))
}

func main() {
  hub := newHub()
	go hub.run()

  http.HandleFunc("/profile", profile)
  http.HandleFunc("/news", newss)
  http.HandleFunc("/profile/edit", editUser)
  http.HandleFunc("/login", login)
  http.HandleFunc("/logout", logout)
  http.HandleFunc("/register", register)
  http.HandleFunc("/index", index)
  http.HandleFunc("/gettimetable", gettimetable)
  http.HandleFunc("/recovery",recovery)
  http.HandleFunc("/addNews",addNews)
  http.HandleFunc("/feed",feed)
  http.HandleFunc("/", root)


	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})


  http.Handle("/js/",http.StripPrefix("/js/", http.FileServer(http.Dir("./scripts"))))
  http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))
  http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./images"))))

  http.ListenAndServe(":8080", nil)
}





//I've copypasted it here because I want to know how does it work
// Hub
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

//cli


const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub *Hub

	// The websocket connection.
	conn *websocket.Conn

  user *pupil

	// Buffered channel of outbound messages.
	send chan []byte
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
    message = append([]byte(c.user.Username + ": "), message...)
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		c.hub.broadcast <- message
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// serveWs handles websocket requests from the peer.
func serveWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
  session, _ := store.Get(r, "cookie-name")
	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256), user: session.Values["user"].(*pupil)}
	client.hub.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()
}
