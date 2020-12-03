package schema

type Settings struct {
	OpenID    string `form:"openID" json:"openID" validate:"max=64"`
	DarkTheme bool   `form:"darkTheme" json:"darkTheme"`
}
