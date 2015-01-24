package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/apcera/nats"
	"github.com/gin-gonic/gin"
)

type Notification struct {
	Channel int    `json:"channel"`
	Message string `json:"message"`
}

func Subscribe(c *gin.Context) {

	t := time.Now()
	var buffer bytes.Buffer
	buffer.WriteString("/ws?_=")
	// buffer.WriteString("string(t.Unix())")
	buffer.WriteString("&tag=&time=&eventid=&channels=1")

	fmt.Println("buffer", buffer.String(), t.Unix())
	c.Writer.Header().Set("X-Accel-Redirect", buffer.String())
	// c.String(200, "")
}

func initCommunication() {
	nc, _ := nats.Connect(nats.DefaultURL)
	c, _ := nats.NewEncodedConn(nc, "json")

	c.Subscribe("notification", func(s string) {
		note := &Notification{}
		err := json.Unmarshal([]byte(s), note)
		if err == nil {
			fmt.Printf("Received a message: %v\n", note)
			go postRequest(note.Channel, []byte(note.Message))
		} else {
			panic(err)
		}
	})
}

func postRequest(channel int, msg []byte) {
	url := fmt.Sprintf("http://0.0.0.0:5000/pub?id=%d", channel)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(msg))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}

func main() {
	router := gin.Default()
	router.Use(gin.Recovery())
	// router.GET("/ws/*params", Subscribe)
	router.GET("/subscribe", Subscribe)
	router.GET("/subscribe/*asd", Subscribe)

	initCommunication()
	router.Run(":8888")
}
