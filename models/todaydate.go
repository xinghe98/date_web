package models

import "time"

// Todaydate 定义一个今日数据的模型
type Todaydate struct {
	Id uint `gorm:"primaryKey"`
	// 时间自动生成，避免用户随意修改时间数据
	DataTime time.Time `json:"datatime" gorm:"column:datatime"`
	// 今日扫码次数
	ScanTimes *int `json:"scantimes" gorm:"column:scantimes" binding:"required,gte=0,lt=200"`
	// 今日新增客户数
	NewUsers *int `json:"newUsers" gorm:"column:newUsers" binding:"required,gte=0,lt=200"`
}
