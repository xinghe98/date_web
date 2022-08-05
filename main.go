package main

import (
	"datevisual/V2/dao"
	"datevisual/V2/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	_ "gorm.io/driver/mysql"
	"os"
)

func InitConfig() {
	workdir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workdir + "/conf")
	err := viper.ReadInConfig()
	if err != nil {

		fmt.Println(err)
		panic(err)
	}
}
func main() {
	InitConfig()
	dao.ConnectMysql()
	dao.InitModel()
	r := gin.Default()
	r = router.RegistRoutes(r)
	r.Run(":3000")
}
