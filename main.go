package main

import (
	"github.com/gin-gonic/gin"
	// "fmt"
	// "net/http"
  	// "gorm.io/driver/mysql"
  	// "gorm.io/gorm"
	// "crypto/sha256"
	"jinghaijun.com/mall/user"
	"jinghaijun.com/mall/password"
	"jinghaijun.com/mall/username"
	"jinghaijun.com/mall/product"
)



func main(){
	r := gin.Default()
	r.GET("/ping", user.Pin)
	r.POST("user/register", user.Register)
	r.POST("user/login", user.Login)
	r.POST("username/nameChange", username.NameChange)
	r.POST("password/reset", password.ChangePassword)
	r.POST("user/addtocart", user.AddtoCart)
	r.GET("cart", product.List)
	r.POST("cart", product.Creat)
	// r.PUT("cart/:id", product.Update) // /cart?id=12&name=13 /cart
	// r.DELETE("cart/:id", product.Delete)
	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}