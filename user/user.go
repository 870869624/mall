package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"
	"crypto/sha256"

	"jinghaijun.com/mall/db"
)

type User struct {
	Username string
	Password string
}

// 注册用user
type RegisterUser struct {
	User
	Name string
}

func Pin(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
} // 基础user
func Register(c *gin.Context) {
	var user RegisterUser
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	connection := db.Get_db()
	h := sha256.New()
	h.Write([]byte(user.Password))
	cryptoPassword := fmt.Sprintf("%x", h.Sum(nil))
	sql := "insert into users(name, username, password, gender, headpicture, age) values ('" + user.Name + "', '" + user.Username + "', '" + cryptoPassword + "')"
	e := connection.Exec(sql)
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

func Login(c *gin.Context) {
	// 1 从请求体(json)中取出传过来的用户名和密码
	var user User
	var check User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	connection := db.Get_db()
	h := sha256.New()
	h.Write([]byte(user.Password))
	cryptoPassword := fmt.Sprintf("%x", h.Sum(nil))
	connection.Raw("select username, password from users where username = ? and password = ?", user.Username, cryptoPassword).Scan(&check)
	fmt.Println("result", check.Username, check.Password)
	if check.Username == "" {
		c.JSON(400, gin.H{
			"message": "登录失败",
		})
	} else {
		c.JSON(200, gin.H{
			"message":  "登录成功",
			"username": check.Username,
			"password": check.Password,
		})
	}
}

//添加进购物车的函数，需要获取用户ID和商品ID，商品所属类目的ID，商品的图片和商品的总价
func AddtoCart(c *gin.Context) {
	var addTocart ADD
	err := c.ShouldBindJSON(&addTocart)
	if err != nil {
		c.JSON(400, gin.H{"message": "参数错误"})
	}
	fmt.Println("------", addTocart.UserID, addTocart.ProductID, addTocart.Price, addTocart.Picture)
	connection := db.Get_db()
	connection.Exec("insert into cart (users_id, catalogue_id, product_id, priceINtotal, picture) values (?, ?, ?, ?, ?)", addTocart.UserID, addTocart.CatalogueID, addTocart.ProductID, addTocart.Price, addTocart.Picture)
}

type ChangeName struct {
	ID int
	User
	Name string
}

//请求体中获取用户的ID和Name
//利用登陆函数已经确认的情况下进行用户昵称的修改
func NameChange(c *gin.Context) {
	var NewName ChangeName
	if err := c.ShouldBindJSON(&NewName); err != nil {
		c.JSON(400, gin.H{"message": "参数错误"})
		return
	}
	fmt.Println(NewName.ID, NewName.Username, NewName.Password, NewName.Name)
	if NewName.ID == 0 || NewName.Name == "" {
		c.JSON(400, gin.H{"message": "参数错误"})
		return
	}
	d := db.Get_db()
	d.Exec("update users set name = ? where id = ?", NewName.Name, NewName.ID)
}
