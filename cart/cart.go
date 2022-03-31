package cart
import(
	"fmt"
	"gorm.io/driver/mysql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)
type cart struct{
	ID int
	Users_id int
	Catalogue_id int
	Product_id int
	PriceINtotal int
	Picture string
}
func Creat(c *gin.Context){
	var add cart
	err := c.ShouldBindJSON(&add)
	if err != nil{
		c.JSON(400, gin.H{"message": "参数错误"})
		return
	}
	dsn := "root:123456@tcp(127.0.0.1:3306)/mall?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
  	if err != nil {
	  panic("db connect error")
  	}
}
func Delete(){

}
func Update(){

}
func List(){

}
