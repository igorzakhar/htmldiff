package main

import (
	"github.com/gin-gonic/gin"
	"htmldiff/diff"
	"net/http"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")
	router.StaticFile("/img/favicon.ico", "./static/img/favicon.ico")
	router.GET("/", index)
	router.POST("/api/v1/htmldiff", getDiff)
	router.Run(":8080")
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Diff service",
	})
}

func getDiff(c *gin.Context) {
	var message struct {
		Text1 string `json:"text1"`
		Text2 string `json:"text2`
	}

	if err := c.ShouldBindJSON(&message); err == nil {
		res, err := diff.DiffHTML(message.Text1, message.Text2)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{"result": res})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"result": err.Error()})
	}
}
