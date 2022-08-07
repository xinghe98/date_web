package models

import "time"

// UsersInfo 定义一个今日数据的模型
type UsersInfo struct {
	Id uint `gorm:"primaryKey"`
	// 时间自动生成，避免用户随意修改时间数据
	CreatedTime time.Time `json:"created_time" gorm:"column:createdtime;not null"`
	// 用户名字
	UserName string `json:"username" gorm:"column:username;not null" binding:"required"`
	// 用户身份证号
	IDCard string `json:"idCard" gorm:"column:idCard;not null" binding:"required,max=18,min=18"`
	// 用户手机号
	TelPhone string `json:"telphone" gorm:"column:telphone;not null" binding:"required,max=11,min=11"`
	// 额度
	Amount int `json:"amount" gorm:"column:amount;not null" binding:"required,lte=100000"`
	// 二维码归属人
	Owner string `json:"owner" gorm:"column:owner;not null" binding:"required"`
}
