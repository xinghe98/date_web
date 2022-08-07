package service

import (
	"datevisual/V2/dao"
	"datevisual/V2/models"
	HttpResp "datevisual/V2/response"
	"datevisual/V2/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// CreateUser 创建今日新增的用户信息
func CreateUser(ctx *gin.Context) {
	var userinfos models.UsersInfo
	err := ctx.ShouldBindJSON(&userinfos)
	if err != nil {
		HttpResp.Res(ctx, http.StatusBadGateway, 502, util.TransLate(err), "数据输入不合法")
		return
	}
	date := time.Now().Format("2006-01-02")
	userinfos.CreatedTime, _ = time.ParseInLocation("2006-01-02", date, time.Local)
	err = dao.DB.Create(&userinfos).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	HttpResp.ResOK(ctx, nil)
}

// FindAllUser 查询所有的新增用户信息
func FindAllUser(ctx *gin.Context) {
	var allusers []models.UsersInfo
	dao.DB.Find(&allusers)
	HttpResp.ResOK(ctx, allusers)
}
