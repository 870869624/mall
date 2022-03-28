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
// 1 检查用户表中是否存在是否存在用户, 如果不存在提示错误， select passwod， username from user where username  = 传过来的用户名;
// 2 如果存在在用户名， 则对比密码是否正确 数据库的密码是否等于你传过来的密码
// 3 根据用户名或者id生成一个令牌返回给前段用户
// 4 前段所有请求都会带上令牌，如果没有带上令牌/或者令牌错误，我们就认为它没有登录

// client -> request- >  ｜ 取出request中的json数据
//  					 ｜ server
// client <- response <- ｜ 返回生成的令牌

func login(c *gin.Context){
	// 1 从请求体(json)中取出传过来的用户名和密码
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	


	
	

	var check User
	dsn := "root:123456@tcp(127.0.0.1:3306)/mall?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		panic("db connect error")
	}
	db.Raw(
		"select username, password from users where username = ? and password = ?", 
		user.Username, 
		user.Password
	).Scan(&check)
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
// func getting(c *gin.Context){
// 	var user1 User
// 	if err := c.ShouldBindJSON(&user1); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	dsn := "root:123456@tcp(127.0.0.1:3306)/mall?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil{
// 		panic("db connect error")
// 	}
	
// }

func main() {

	r := gin.Default()
	r.GET("/ping", pin)

	r.POST("/register", register)
	r.POST("/login", login)
	r.GET("/someGet", getting)
	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}