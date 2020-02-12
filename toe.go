package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("view/*")
	router.StaticFile("/toe.png", "./view/toe.png")
	router.GET("/", getpage)
	router.GET("/tac", loadgame)
	router.Run(":8080")
}
func loadgame(c *gin.Context) {
	c.HTML(http.StatusOK, "loadgame.html", gin.H{
		"title": "Main website",
	})
}
func getpage(c *gin.Context) {
	c.HTML(http.StatusOK, "toe.html", gin.H{
		"title": "Main website",
	})

}
