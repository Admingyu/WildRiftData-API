package serialization

type DevLogSer struct {
	Text    string   `json:"text" gorm:"column:title"`
	Desc    jsonDate `json:"desc" gorm:"column:date"`
	Content string   `json:"content"`
	Version string   `json:"version"`
}

type About struct {
	Content string `json:"content"`
}
