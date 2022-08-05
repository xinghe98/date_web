package service

import (
	"datevisual/V2/dao"
	"datevisual/V2/models"
	HttpResp "datevisual/V2/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Create(ctx *gin.Context) {
	var data *models.Todaydate
	ctx.ShouldBindJSON(&data)
	dao.DB.Create(data)
	HttpResp.Res(ctx, http.StatusOK, 200, nil, "创建成功")
}
