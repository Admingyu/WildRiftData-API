package model

type User struct {
	BaseModel
	NickName  string `json:"nick_name" gorm:"type:varchar(50);comment:'昵称'"`
	AvatarUrl string `json:"avatar_url" gorm:"type:varchar(1024);comment:'头像'"`
	OpenID    string `json:"open_id" gorm:"type:varchar(512);uniqueIndex:OpenIDUniq;comment:'OpenID'"`
	Phone     string `json:"phone" gorm:"type:varchar(20);comment:'手机号'"`
	Gender    int    `json:"gender" gorm:"type:int;comment:'性别'"`
	Province  string `json:"province" gorm:"type:varchar(50);comment:'省会'"`
	City      string `json:"city" gorm:"type:varchar(50);comment:'城市'"`
	Language  string `json:"language" gorm:"type:varchar(50);comment:'语言'"`
	Country   string `json:"country" gorm:"type:varchar(50);comment:'国家'"`
	Darktheme bool   `json:"dark_theme" gorm:"default:0;type:tinyint(1);comment:'是否暗黑主题'"`
	TimeModel
}

type Device struct {
	BaseModel
	User   User `gorm:"foreignKey:UserID"`
	UserID int  `json:"user_id" gorm:"uniqueIndex:DeviceUserUniq;comment:'用户id'"`

	Brand    string `json:"brand" gorm:"type:varchar(50);comment:'厂商'"`
	Model    string `json:"model" gorm:"type:varchar(50);comment:'型号'"`
	Language string `json:"language" gorm:"type:varchar(50);comment:'语言'"`
	Platform string `json:"platform" gorm:"type:varchar(50);comment:'平台'"`
	System   string `json:"system" gorm:"type:varchar(50);comment:'系统'"`
	Version  string `json:"version" gorm:"type:varchar(50);comment:'版本'"`
	TimeModel
}

type WiFiNetWork struct {
	BaseModel
	User   User `gorm:"foreignKey:UserID"`
	UserID int  `json:"user_id" gorm:"comment:'用户id'"`

	SSID      string `json:"ssid" gorm:"type:varchar(50);comment:'ssid'"`
	BSSID     string `json:"bssid" gorm:"type:varchar(50);comment:'bssid'"`
	Signal    int    `json:"signal" gorm:"type:int;comment:'signal'"`
	Frequency int    `json:"frequency" gorm:"type:int;comment:'frequency'"`
	TimeModel
}

type ClickBoard struct {
	BaseModel
	User   User `gorm:"foreignKey:UserID"`
	UserID int  `json:"user_id" gorm:"comment:'用户id'"`

	Content string `json:"content" gorm:"type:text;comment:'内容'"`
	TimeModel
}
