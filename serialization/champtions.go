package serialization

type ChampionSearch struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Title     string `json:"title"`
	Thumbnail string `json:"thumbnail"`
}

type ChampionInfo struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Title           string `json:"title"`
	HeroImage       string `json:"hero_image"`
	DifficultyLevel string `json:"difficulty_level"`
	Video           string `json:"video"`
}

type AbilityInfo struct {
	ID          int    `json:"id"`
	Type        string `json:"type" gorm:"type:varchar(10)"`
	Title       string `json:"title" gorm:"type:varchar(20)"`
	Thumbnail   string `json:"thumbnail" gorm:"type:varchar(1024);comment:'缩略图'"`
	Description string `json:"description" gorm:"type:varchar(500);comment:'描述'"`
	Video       string `json:"video" gorm:"type:varchar(1024);comment:'视频'"`
	PosterImage string `json:"poster_image" gorm:"type:varchar(1024);comment:'海报地址'"`
}

type Role struct {
	Name        string `json:"name" gorm:"type:varchar(20)"`
	MachineName string `json:"machine_name" gorm:"type:varchar(20)"`
	Icon        string `json:"icon" gorm:"type:varchar(1024)"`
}
