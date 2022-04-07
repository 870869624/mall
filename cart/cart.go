package cart

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"jinghaijun.com/mall/db"
)

type Cart struct {
	ID           int
	Users_id     int
	Catalogue_id int
	Product      []Product
	TotalCount   int //商品总价格
	TotalAmount  int //商品总数量
}
type Product struct {
	ID     int
	Name   string
	CartID string //属于哪一个购物车
	Price  int    //总价格
	Total  int    //总数量
	Size   string //尺寸
}

//取得总数量的方法
func (c *Cart) GetTotalAmount() int {
	var totalCount int

}

func Creat(c *gin.Context) {
	var addTocart Cart
	err := c.ShouldBindJSON(&addTocart)
	if err != nil {
		c.JSON(400, gin.H{"message": "参数错误"})
	}
	fmt.Println("------", addTocart.ID, addTocart.Users_id, addTocart.Price, addTocart.Picture)
	connection := db.Get_db()
	connection.Exec("insert into cart (users_id, catalogue_id, product_id, priceINtotal, picture) values (?, ?, ?, ?, ?)", addTocart.UserID, addTocart.CatalogueID, addTocart.ProductID, addTocart.Price, addTocart.Picture)
}

// func Delete() {

// }
// func Update() {

// }

func List(c *gin.Context) {
	id := c.Param("id")
	connection := db.Get_db()
	var ProductsInCart []Cart
	fmt.Println("id", id)
	i, e := strconv.Atoi(id)
	fmt.Println(i)
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "参数错误"})
		return
	}
	connection.Table("cart").Where("users_id = ?", i).Scan(&ProductsInCart)
	if ProductsInCart == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "查找用户失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": ProductsInCart})
}
