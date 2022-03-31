package username
import (
	"github.com/gin-gonic/gin"
	"fmt"
	// "net/http"
  	// "gorm.io/driver/mysql"
  	// "gorm.io/gorm"
	// "crypto/sha256"
	"jinghaijun.com/mall/db"
)
type User struct {
	Username string
	Password string
}
type ChangeName struct{
	ID int
	User
	Name string
}
//请求体中获取用户的ID和Name
//利用登陆函数已经确认的情况下进行用户昵称的修改
func NameChange(c *gin.Context){
	var NewName ChangeName
	if err := c.ShouldBindJSON(&NewName); err != nil{
		c.JSON(400, gin.H{"message": "参数错误"})
		return
	}
	fmt.Println(NewName.ID, NewName.Username, NewName.Password, NewName.Name)
	if NewName.ID == 0 || NewName.Name == ""{
		c.JSON(400, gin.H{"message": "参数错误"})
		return 
	}
	fmt.Println(NewName.ID, NewName.Username, NewName.Password, NewName.Name)
	if NewName.ID == 0 || NewName.Name == ""{
		c.JSON(400, gin.H{"message": "参数错误"})
		return 
	}
	d := db.Get_db()
	d.Exec("update users set name = ? where id = ?", NewName.Name, NewName.ID)
}