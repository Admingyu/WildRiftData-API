package model

import "time"

type BaseModel struct {
	ID int `json:"id" gorm:"primary_key"`
}

type TimeModel struct {
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime(6);not null;default:CURRENT_TIMESTAMP(6)"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime(6);not null;default:CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6)"`
}
