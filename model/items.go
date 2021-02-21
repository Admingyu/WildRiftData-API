package model

type Items struct {
	BaseModel
	Name        string `json:"name" gorm:"type:varchar(255);comment:'名称'"`
	Icon        string `json:"icon" gorm:"type:varchar(1024);comment:'图标'"`
	Price       int    `json:"price" gorm:"type:int(10);comment:'合成价'"`
	Description string `json:"description" gorm:"type:text;comment:'描述'"`
	PlainText   string `json:"plaintext" gorm:"type:varchar(255);comment:'简述'"`
	Sell        int    `json:"sell" gorm:"type:int(255);comment:'卖出价格'"`
	Total       int    `json:"total" gorm:"type:int(255);comment:'总价'"`
	Tag         string `json:"tag" gorm:"type:varchar(255);comment:'标签'"`
	Keywords    string `json:"keywords" gorm:"type:varchar(255);comment:'关键词'"`
}

type ItemFroms struct {
	BaseModel
	FatherID int `json:"father_id"`
	ChildID  int `json:"child_id"`
}

type ItemTypes struct {
	BaseModel
	ItemID   int    `json:"item_id"`
	ItemType string `json:"item_type" gorm:"type:varchar(255);comment:'类型'"`
}
