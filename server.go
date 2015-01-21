package main

import (
	"bytes"
	"time"
"fmt"
"io/ioutil"
 "net/http"
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
//	c.Redirect(307, buffer.String())
        //c.Writer.Header().Set("X-Accel-Redirect", buffer.String())
	//c.Writer.Header().Set("Content-Type","text/plain")
	//c.Writer.Header().Set("Cache-Control","no-cache")
	defer c.Request.Body.Close()
	client := &http.Client{}
	req,_ := http.NewRequest("POST","http://0.0.0.0/pub",c.Request.Body)
	req.Header.Add("X-Accel-Redirect",buffer.String())
	_,err:=client.Do(req)
	if err!=nil{
		fmt.Println("there's an error :(")
	}else{
		body, _ := ioutil.ReadAll(c.Request.Body)
		fmt.Println("posted with",body)
	}
	c.String(200,"yay")

}


func main() {
	router := gin.Default()
	router.Use(gin.Recovery())
	router.GET("/subscribe", Subscribe)
	router.POST("/publish", Publish)
	router.Run(":8000")
}
