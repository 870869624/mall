package main

import (
	"github.com/gin-gonic/gin"
	"jinghaijun.com/mall/cart"
	"jinghaijun.com/mall/password"
	"jinghaijun.com/mall/product"
	"jinghaijun.com/mall/user"
)

func main() {
	r := gin.Default()
	r.GET("/ping", user.Pin)
	r.POST("user/register", user.Register)
	r.POST("user/login", user.Login)
	r.PUT("user/nameChange", user.NameChange)
	r.POST("password/reset", password.ChangePassword)
	r.POST("user/addtocart", user.AddtoCart)
	r.GET("product/List", product.List)
	r.POST("product", product.Creat)
	r.PUT("product/:name", product.Update) // /cart?id=12&name=13 /cart（这里是根据ID更新产品信息）
	r.DELETE("product/:id", product.Delete)
	r.GET("cart/List/:id", cart.List)
	r.POST("", cart.Creat)
	//r.DELETE("cart/:name", product.Delete)(具体根据什么删除可以从这里入手)
	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	r.POST("")
}
