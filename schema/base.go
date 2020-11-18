package schema

type IdSchema struct {
	ID int `form:"id" json:"id" bind:"required,min=1"`
}
