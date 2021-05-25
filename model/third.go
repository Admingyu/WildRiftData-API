package model

import "time"

type Temperature struct {
	ID          int       `json:"id" gorm:"primary_key"`
	Name        string    `json:"name" gorm:"type:varchar(20)"`
	Temperature string    `json:"temperature" gorm:"type:varchar(20)"`
	Time        time.Time `json:"time" gorm:"type:datetime(6);not null;default:CURRENT_TIMESTAMP(6)"`
}
