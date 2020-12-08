package model

import "github.com/braintree-go/braintree-go/date"

type DevelopLog struct {
	BaseModel
	Title   string    `json:"title" gorm:"type:varchar(64);comment:'标题'"`
	Content string    `json:"content" gorm:"type:text;comment:'内容'"`
	Date    date.Date `json:"date" gorm:"type:date;comment:'日期'"`
	Version string    `json:"version" gorm:"type:varchar(50);comment:'版本'"`
	TimeModel
}

type About struct {
	BaseModel
	Content string `json:"content" gorm:"type:text;comment:'内容'"`
}
