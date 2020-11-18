package model

type Champion struct {
	BaseModel

	Name           string `json:"name" gorm:"type:varchar(20)"`
	NameEn         string `json:"name_en" gorm:"type:varchar(20)"`
	Title          string `json:"title" gorm:"type:varchar(20)"`
	Image          string `json:"image" gorm:"type:varchar(1024)"`
	HeroImage      string `json:"hero_mage" gorm:"type:varchar(1024)"`
	DifﬁcultyLevel string `json:"difficulty_level" gorm:"type:varchar(20)"`
	Video          string `json:"video" gorm:"type:varchar(1024)"`
	PosterImage    string `json:"poster_image" gorm:"type:varchar(1024)"`
	Description    string `json:"description" gorm:"type:varchar(1024)"`
	Story          string `json:"story" gorm:"type:varchar(2048)"`
	AliasTip       string `json:"alias_tip" gorm:"type:varchar(1024)"`
	EnemyTip       string `json:"enemy_tip" gorm:"type:varchar(1024)"`
	PosterImagePC  string `json:"poster_image_pc" gorm:"type:varchar(1024)"`
}

type Role struct {
	BaseModel

	Name        string `json:"name" gorm:"type:varchar(20)"`
	MachineName string `json:"machine_name" gorm:"type:varchar(20)"`
	Icon        string `json:"icon" gorm:"type:varchar(1024)"`
}

type ChampionRole struct {
	BaseModel

	Champion Champion `gorm:"foreignKey:ChampionID"`
	Roles    Role     `gorm:"foreignKey:RoleID"`

	ChampionID int `json:"champion_id" gorm:"type:int;not null"`
	RoleID     int `json:"role_id" gorm:"type:int;not null"`
}

type ChampionAbilities struct {
	BaseModel

	Champion Champion `gorm:"foreignKey:ChampionID"`

	ChampionID  int    `json:"champion_id" gorm:"type:int;not null"`
	Type        string `json:"type" gorm:"type:varchar(10)"`
	Title       string `json:"title" gorm:"type:varchar(20)"`
	Thumbnail   string `json:"thumbnail" gorm:"type:varchar(1024);comment:'缩略图'"`
	Description string `json:"description" gorm:"type:varchar(500);comment:'描述'"`
	Video       string `json:"video" gorm:"type:varchar(1024);comment:'视频'"`
	PosterImage string `json:"poster_image" gorm:"type:varchar(1024);comment:'海报地址'"`
}
