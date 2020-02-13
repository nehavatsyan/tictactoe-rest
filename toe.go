package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Player database structure
type player struct {
	gorm.Model
	Username string `gorm:"type:varchar(100)" `
	Name     string `gorm:"type:varchar(100)" `
	Win      int    `gorm:"type:number" `
	Loss     int    `gorm:"type:number" `
	Draw     int    `gorm:"type:number" `
	Email    string `gorm:"type:varchar(100)" `
	Password string `gorm:"type:varchar(100)" `
}

// main function render data
func main() {
	router := gin.Default()                         //start router engine
	router.LoadHTMLGlob("view/*")                   // load html files of view folder
	router.StaticFile("/toe.png", "./view/toe.png") //load static png file
	router.GET("/", getpage)                        // execute get request
	router.GET("/tac", loadgame)
	router.POST("/submit", postData) // loadgame on submit
	router.Run(":8080")              // Run engine on localhost 8080
}

// Load Game if username is correct

func loadgame(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	db, err := gorm.Open("sqlite3", "view/test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	db.AutoMigrate(&player{})
	var dp, da player
	//check if username exists
	if db.Where("Username = ?", name).First(&player{}).RecordNotFound() == true {
		c.HTML(http.StatusOK, "signup.html", gin.H{})
	} else if db.Where("Username = ? AND Password = ?", name, password).First(&player{}).RecordNotFound() {
		c.HTML(http.StatusOK, "toe.html", gin.H{
			"title": "Wrong Password",
		})
	} else {

		db.Where("name = ?", name).First(&dp)
		db.Where("name = ?", "ai").First(&da)
		c.HTML(http.StatusOK, "loadgame.html", gin.H{
			"title": name,
			"win":   dp.Win,
			"loss":  dp.Loss,
			"draw":  dp.Draw,
			"wina":  da.Win,
			"lossa": da.Loss,
			"drawa": da.Draw,
		})
	}

}

//Get Signin Page
func getpage(c *gin.Context) {
	c.HTML(http.StatusOK, "toe.html", gin.H{})

}

//Get Postdata from signup form
func postData(c *gin.Context) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	db.AutoMigrate(&player{})
	name := c.PostForm("name")
	username := c.PostForm("username")
	password := c.PostForm("psw")
	email := c.PostForm("email")
	win := 0
	loss := 0
	draw := 0
	fmt.Print(name, username, password, email)
	var p = player{Username: username, Name: name, Win: win, Loss: loss, Draw: draw, Email: email, Password: password}
	db.Create(&p)
	c.HTML(http.StatusOK, "toe.html", gin.H{})
}
