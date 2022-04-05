package product
import(
	// "gorm.io/driver/mysql"
	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"
	"net/http"
	"fmt"
	"jinghaijun.com/mall/db"
	"strconv"
	"strings"
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
type Change struct{
	ID int
	Name string
}
//从前端获得数据并且计入到商品信息中中。此处属于商品信息录入
func Creat(c *gin.Context){
	var add Product
	err := c.ShouldBindJSON(&add)
	if err != nil{
		c.JSON(400, gin.H{"message": "参数错误"})
		return
	}
	connection := db.Get_db()
	fmt.Println(add.Price, add.Total, add.ID, add.Name, add.Image, add.Size)
	connection.Exec("insert into product (id, name, image, price, total, size) values (?, ?, ?, ?, ?, ?)", add.ID, add.Name, add.Image, add.Price, add.Total, add.Size)
}
//从前端获取ID或者是商品名称进行匹配然后删除数据
func Delete(c *gin.Context){
	var delete Change
	err := c.ShouldBindJSON(&delete)
	if err != nil{
		c.JSON(400, gin.H{"message": "参数错误"})
		return
	}
	connection := db.Get_db()
	fmt.Println("删除的是编号为：",delete.ID)
	connection.Exec("Delete from product where id = ?", delete.ID)
	if err == nil{
		c.JSON(400, gin.H{"message": "删除失败"})
	}
}
// 从前端获取到产品的各种信息根据要求进行修改(此处只是更新单个信息，例如产品名字)
func Update(c *gin.Context){
	var New Change
	if err := c.ShouldBindJSON(&New); err != nil{
		c.JSON(400, gin.H{"message": "参数错误"})
		return
	}
	connection := db.Get_db()
	fmt.Println(New.Name)
	connection.Exec("Update product set name = ? where id = ?", New.Name, New.ID)
}


type ProductQuery struct {
	name string
	price_gt float64
	price_lt float64
}
//从前端获取到请求要求查看商品信息，此处返回（后期修改为用ID或者商品名查看或者根据价格区间查找）
func List(c *gin.Context){
	var ShowProduct []Product
	connection := db.Get_db()
	name := c.DefaultQuery("name", "")
	price_gt := c.DefaultQuery("price_gt", "")
	price_lt := c.DefaultQuery("price_lt", "")
	query := []interface{}{}
	queryString := []string{}
	if (name != "") {
		query = append(query, name)
		queryString = append(queryString, "name = ?");
	}
	if (price_gt != "") {
		v, _ := strconv.Atoi(price_gt)
		query = append(query, v)
		queryString = append(queryString, "price >= ?");
	}

	if (price_lt != "") {
		v, _ := strconv.Atoi(price_lt)
		query = append(query, v)
		queryString = append(queryString, "price < ?");
	}
	connection.Table("product").Where(strings.Join(queryString, " and "), query...).Find(&ShowProduct);
	c.JSON(http.StatusOK, gin.H{
		"result": ShowProduct,
	})
}
