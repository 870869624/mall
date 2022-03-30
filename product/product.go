package product
import(
	"gorm.io/driver/mysql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"fmt"
)
//用户发送请求信息，调用数据库内容并且返回给前端
type Product struct{
	ID int
	Name string
	Image string
	Price int
	Total int
	Size string
}
//从前端获得数据并且计入到商品信息中中。
func Creat(c *gin.Context){
	var add Product
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
	fmt.Println(add.Price, add.Total, add.ID, add.Name, add.Image, add.Size)
	db.Exec("insert into prodouct (id, name, image, price, total, size) values (?, ?, ?, ?, ?, ?)", add.ID, add.Name, add.Image, add.Price, add.Total, add.Size)
}

// func Delete(){

// }

// func Update(){

// }
func List(c *gin.Context){
	var ShowProduct Product
	dsn := "root:123456@tcp(127.0.0.1:3306)/mall?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("db connect error")
	}
	db.Raw("select * from product").Scan(&ShowProduct)
	c.JSON(http.StatusOK, gin.H{
		"ID": ShowProduct.ID,
		"Name": ShowProduct.Name,
		"image": ShowProduct.Image,
		"Price": ShowProduct.Price,
		"Total": ShowProduct.Total,
		"SIze": ShowProduct.Size,
	})
}
