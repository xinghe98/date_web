package service

import (
	"datevisual/V2/dao"
	"datevisual/V2/models"
	HttpResp "datevisual/V2/response"
	"datevisual/V2/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

// Create 创建今日数据的逻辑层
func Create(ctx *gin.Context) {
	var data *models.Todaydate
	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		errs, _ := err.(validator.ValidationErrors)
		HttpResp.Res(ctx, http.StatusBadGateway, 502, errs.Translate(util.Trans), "请按要求输入失败")
		return
	}
	data.DataTime = time.Now()
	err = dao.DB.Create(&data).Error
	if err != nil {
		panic(err)
	}
	HttpResp.Res(ctx, http.StatusOK, 200, nil, "创建成功")
}
