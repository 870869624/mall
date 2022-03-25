package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
  	"gorm.io/driver/mysql"
  	"gorm.io/gorm"
)

// 基础user
type User struct {
	Username string
	Password string
}

// 注册用user
type RegisterUser struct {
	User
	Name string
}

func pin (c *gin.Context){
		c.JSON(200, gin.H{
			"message": "pong",
		})
}
func register(c *gin.Context) {
	var user RegisterUser
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dsn := "root:123456@tcp(127.0.0.1:3306)/mall?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("db connect error")
	}

	sql := "insert into users(name, username, password, gender, headpicture, age) values ('" + user.Name + "', '" + user.Username + "', '" + user.Password +"', 0, '', 0)";
	e := db.Exec(sql);
	if e == nil {
		c.JSON(400, gin.H{
			"message": "注册失败",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "注册成功",
		})
	}	
}
func login(c *gin.Context){
	var user User
	var check User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dsn := "root:123456@tcp(127.0.0.1:3306)/mall?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		panic("db connect error")
	}
	query := "select username, password from users where username = ? and password = ?"
	db.Raw(query, user.Username, user.Password).Scan(&check)
	fmt.Println("result", check.Username, check.Password);
	// if check == nil {
	// 	c.JSON(400, gin.H{
	// 		"message": "登录失败",
	// 	})
	// } else {
		c.JSON(200, gin.H{
			"message": "登录成功",
			"username": check.Username,
			"password": check.Password,
		})
	// }
}

func main() {

	r := gin.Default()
	r.GET("/ping", pin)

	r.POST("/register", register)
	r.POST("/login", login)

	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}