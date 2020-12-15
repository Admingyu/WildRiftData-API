package schema

type NewsUpdateSchema struct {
	TimeStamp int    `json:"time_stamp" binding:"required"`
	AuthKey   string `json:"auth_key" binding:"required"`
	Content   string `json:"content" binding:"required"`
	NewsID    string `json:"news_id" binding:"required"`
}

type NewsUniqIDSchema struct {
	UniqID string `json:"uniq_id" form:"uniq_id" binding:"required"`
}
