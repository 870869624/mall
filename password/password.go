package password
import (
	"github.com/gin-gonic/gin"
	"fmt"
	// "net/http"
  	"gorm.io/driver/mysql"
  	"gorm.io/gorm"
	"crypto/sha256"
)
type PasswordChange struct{
	Id int
	Password string
}
func ChangePassword(c *gin.Context){
	var data PasswordChange
	// 从请求体里获取用户ID和密码
	err := c.ShouldBindJSON(&data)
	if err != nil{
		c.JSON(400, gin.H{"message": "参数错误"})
		return
	}
	fmt.Println("----", data.Id, data.Password)
	if data.Id == 0 || data.Password == ""{ 
		c.JSON(400, gin.H{"message": "参数不全"})
	}
	dsn := "root:123456@tcp(127.0.0.1:3306)/mall?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		panic("db connect error")
	}
	var record PasswordChange
	db.Raw("select password, id from users where id = ?", data.Id).Scan(&record)
	fmt.Println("record:", record.Id, record.Password)
	if record.Id == 0{
		c.JSON(400, gin.H{"message": "没有找到该用户"})
	}
	// 加密data.Password，使用sha256
	// 构造加密因子
	h := sha256.New()
	h.Write([]byte(data.Password))
	cryptoPassword := fmt.Sprintf("%x", h.Sum(nil))
	result := db.Exec("update users set password = ? where id = ?", cryptoPassword, data.Id)
	if result.Error != nil {
		c.JSON(400, gin.H{"message": "操作失败"})
		return;
	}
		c.JSON(200, gin.H{"message": "修改成功"})
}