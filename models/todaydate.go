package models

import "time"

type Todaydate struct {
	Id        uint      `gorm:"primaryKey"`
	DataTime  time.Time `json:"datatime" gorm:"column:datatime"`
	ScanTimes int       `json:"scantimes" gorm:"column:datatime"`
	NewUsers  int       `json:"newUsers" gorm:"column:newUsers"`
}
