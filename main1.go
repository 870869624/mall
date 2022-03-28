package main

import(
	"fmt"
	"github.com/gin-gonic/gin"
)

type register struct{
	User string
	Password string
}

func getting(c *gin.Context){
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
func register(c *gin.Context){
	var user register
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dsn := "root:123456@tcp(127.0.0.1:3306)/mall?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		panic("you can`t do this")
	}
	sql := "insert into users(name, gender, age, headpicture, username,password) values ('" + users.name + "', '" + users.username +"', '" + users.password + "')"
	db.Exec(sql)
}


func main() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.GET("/someGet", getting)
	router.POST("/somePost", register)
	router.PUT("/somePut", putting)
	router.DELETE("/someDelete", deleting)
	router.PATCH("/somePatch", patching)
	router.HEAD("/someHead", head)
	router.OPTIONS("/someOptions", options)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()
	// router.Run(":3000") for a hard coded port
}