package model

type User struct {
	BaseModel
	NickName  string `json:"nick_name" gorm:"type:varchar(20);comment:'昵称'"`
	AvatarUrl string `json:"head_image" gorm:"type:varchar(1024);comment:'头像'"`
	OpenID    string `json:"open_id" gorm:"type:varchar(512);comment:'OpenID'"`
	Phone     string `json:"phone" gorm:"type:varchar(20);comment:'手机号'"`
	Gender    string `json:"gender" gorm:"type:varchar(5);comment:'性别'"`
	Province  string `json:"province" gorm:"type:varchar(20);comment:'省会'"`
	City      string `json:"city" gorm:"type:varchar(20);comment:'城市'"`
	Country   string `json:"country" gorm:"type:varchar(20);comment:'国家'"`
}
