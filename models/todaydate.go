package models

import "time"

type Todaydate struct {
	DataTime  time.Time `json:"datatime"`
	ScanTimes int       `json:"scantimes"`
}
