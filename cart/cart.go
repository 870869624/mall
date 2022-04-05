package cart

import (
	"github.com/gin-gonic/gin"
	"jinghaijun.com/mall/db"
)

type cart struct {
	ID           int
	Users_id     int
	Catalogue_id int
	Product_id   int
	PriceINtotal int
	Picture      string
}

func Creat(c *gin.Context) {
	var add cart
	err := c.ShouldBindJSON(&add)
	if err != nil {
		c.JSON(400, gin.H{"message": "参数错误"})
		return
	}
	connection := db.Get_db()
}
func Delete() {

}
func Update() {

}
func List() {

}
