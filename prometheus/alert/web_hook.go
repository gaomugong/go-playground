package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/webhook", func(c *gin.Context) {
		data, _ := ioutil.ReadAll(c.Request.Body)
		_ = hookHandler(string(data))
		c.JSON(http.StatusOK, gin.H{
			"message": "alert received.",
		})
	})
	_ = r.Run(":5001")
}

func hookHandler(data string) (err error) {
	log.Printf("Hook start, data:\n %v", data)
	return nil
}
