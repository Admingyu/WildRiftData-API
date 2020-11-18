package model

type BaseModel struct {
	ID int `json:"id" gorm:"primary_key"`
}
