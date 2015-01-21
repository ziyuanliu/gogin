package main

import (
	"bytes"
	"time"

	"github.com/gin-gonic/gin"
)

type MessageJSON struct {
	Channel string `json:"channel" binding:"required"`
	Message string `json:"message" binding:"required"`
}

type SubscribeJSON struct {
	Channel string `json:"channel" binding:"required"`
	Time    string `json:"_" binding:"required"`
}

func Subscribe(c *gin.Context) {

	// var json SubscribeJSON
	// c.Bind(&json)
	// c_cp := c.Copy()
	// go func() {
	var t time.Time
	var buffer bytes.Buffer
	buffer.WriteString("/ws?_=")
	buffer.WriteString(string(t.UnixNano()))
	buffer.WriteString("&tag=&time=&eventid=&channels=1")
	// buffer.WriteString(json.Time)
	// buffer.WriteString("&tag=&time=&eventid=&channels=")
	// buffer.WriteString(json.Channel)

	c.Writer.Header().Set("X-Accel-Redirect", buffer.String())
	// }()

}

func Publish(c *gin.Context) {

	var json MessageJSON
	okay := c.Bind(&json)
	if !okay {
		c.JSON(403, gin.H{"error": "malformed JSON"})
		return
	}

	var buffer bytes.Buffer
	buffer.WriteString("/pub?id=")
	buffer.WriteString(json.Channel)
	// fmt.Println("redirecting to", buffer.String())
	c.Redirect(307, buffer.String())

}

func main() {
	router := gin.Default()
	router.Use(gin.Recovery())
	router.GET("/subscribe", Subscribe)
	router.POST("/publish", Publish)
	router.Run(":8000")
}
