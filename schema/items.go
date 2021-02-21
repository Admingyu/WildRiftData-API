package schema

type ItemSearchSchema struct {
	Input string `form:"input" json:"input" validate:"max=16"`
	Type  string `form:"type" json:"type" validate:"max=16"`
}
