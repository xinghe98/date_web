package service

import (
	"datevisual/V2/dao"
	"datevisual/V2/models"
	HttpResp "datevisual/V2/response"
	"datevisual/V2/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Create 创建今日数据的逻辑层
func Create(ctx *gin.Context) {
	var data *models.Todaydate
	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		//errs, _ := err.(validator.ValidationErrors)
		HttpResp.Res(ctx, http.StatusBadGateway, 502, util.TransLate(err), "数据输入不合法")
		return
	}
	data.DataTime = time.Now()
	err = dao.DB.Create(&data).Error
	if err != nil {
		panic(err)
		return
	}
	HttpResp.ResOK(ctx, nil)
}
