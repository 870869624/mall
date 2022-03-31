package db

import(
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
func Get_db() *gorm.DB{
	dsn := "root:123456@tcp(127.0.0.1:3306)/mall?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		panic("db connect error")
	}
	return db
}