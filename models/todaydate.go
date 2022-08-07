package models

import "time"

// Todaydate 定义一个今日数据的模型
type Todaydate struct {
	Id uint `gorm:"primaryKey"`
	// 时间自动生成，避免用户随意修改时间数据
	DataTime time.Time `json:"datatime" gorm:"column:datatime;not null;unique"`
	// 今日扫码次数
	ScanTimes *int `json:"scantimes" gorm:"column:scantimes;not null" binding:"required,gte=0,lt=200"`
	// 今日新增客户数
	NewUsers *int `json:"newUsers" gorm:"column:newUsers;not null" binding:"required,gte=0,lt=200"`
	// 关联今日成功的用户列表
	SuccUser []*UsersInfo `json:"succuser" gorm:"foreignKey:CreatedTime;references:DataTime;constraint:OnUpdate:CASCADE"`
}
