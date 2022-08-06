package dao

import (
	"datevisual/V2/models"
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB 创建一个全局DB对象，后面直接用这个操作数据库
var DB *gorm.DB

func ConnectMysql() *gorm.DB {
	url := viper.GetString("mysql.url")
	username := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	var err error
	database := fmt.Sprintf("%s:%s@tcp%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, url)
	DB, err = gorm.Open(mysql.Open(database), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return DB
}

// InitModel 在数据库中建立对应的表结构
func InitModel() {
	DB.AutoMigrate(&models.Todaydate{})
}
