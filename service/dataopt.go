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
	scantime := data.ScanTimes
	NewUser := data.NewUsers
	if err != nil {
		//errs, _ := err.(validator.ValidationErrors)
		HttpResp.Res(ctx, http.StatusBadGateway, 502, util.TransLate(err), "数据输入不合法")
		return
	}
	date := time.Now().Format("2006-01-02")
	data.DataTime, _ = time.ParseInLocation("2006-01-02", date, time.Local)
	dao.DB.Where("datatime=?", data.DataTime).First(&data)
	if data.Id != 0 {
		//err = dao.DB.Model(&models.Todaydate{}).Where("datatime=?", data.DataTime).Update("scantimes", scantime).Update("newUsers", NewUser).Error
		dao.DB.Model(&data).Updates(models.Todaydate{ScanTimes: scantime, NewUsers: NewUser})
		HttpResp.ResOK(ctx, nil)
		return
	}
	err = dao.DB.Create(&data).Error
	if err != nil {
		panic(err)
		return
	}
	HttpResp.ResOK(ctx, nil)
}

// FindDateByTime 根据日期查询
func FindDateByTime(ctx *gin.Context) {

	var data []models.Todaydate
	dayStart, Ok := ctx.GetQuery("timestart")
	if !Ok {
		dao.DB.Preload("SuccUser").Find(&data)
		HttpResp.ResOK(ctx, data)
		return
	}
	dayEnd, ok := ctx.GetQuery("timeend")
	if !ok {
		dao.DB.Where("datatime = ?", dayStart).Preload("SuccUser").Find(&data)
		HttpResp.ResOK(ctx, data)
		return
	} else {
		dao.DB.Where("datatime BETWEEN ? AND ? ", dayStart, dayEnd).Find(&data)
		HttpResp.ResOK(ctx, data)
	}

}
