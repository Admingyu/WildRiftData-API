package serialization

type News struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Thumbnail string `json:"thumbnail" gorm:"column:thumb_nail"`
	Date      string `json:"date"`
	Category  string `json:"category"`
	Link      string `json:"link"`
}

type NewsDetail struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Thumbnail   string `json:"thumbnail" gorm:"column:thumb_nail"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Category    string `json:"category"`
	Link        string `json:"link"`
	Content     string `json:"content"`
}
