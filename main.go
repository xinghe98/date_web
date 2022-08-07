package main

import (
	middleware "datevisual/V2/Middleware"
	"datevisual/V2/dao"
	"datevisual/V2/router"
	"datevisual/V2/util"
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
	util.InitTrans()
	dao.ConnectMysql()
	dao.InitModel()
	r := gin.Default()
	r.Use(middleware.Cors())
	r = router.RegistRoutes(r)
	r.Run(":3000")
}
