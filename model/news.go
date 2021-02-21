package model

type News struct {
	BaseModel

	UniqID      string `json:"uniq_id" gorm:"type:varchar(50);uniqueIndex:NewsUniq;commnet:'唯一id'"`
	Title       string `json:"title" gorm:"type:varchar(50);commnet:'标题'"`
	ThumbNail   string `json:"thumb_nail" gorm:"type:varchar(1024);commnet:'缩略图'"`
	Description string `json:"description" gorm:"type:varchar(100);commnet:'简要'"`
	Content     string `json:"content" gorm:"type:mediumtext;commnet:'内容'"`
	Date        string `json:"date" gorm:"type:varchar(20);commnet:'日期'"`
	Link        string `json:"link" gorm:"type:varchar(500);commnet:'外链'"`
	Category    string `json:"category" gorm:"type:varchar(20);commnet:'默认分类'"`
	TimeModel
}

type NewsCategory struct {
	BaseModel
	News News `gorm:"foreignKey:NewsID"`

	NewsID   int    `json:"news_id" gorm:"type:foreign_key;commnet:'新闻id'"`
	Name     string `json:"name" gorm:"type:varchar(20);commnet:'分类名称'"`
	Category string `json:"category" gorm:"type:varchar(20);commnet:'分类'"`
	TimeModel
}
