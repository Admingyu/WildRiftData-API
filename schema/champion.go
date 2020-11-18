package schema

type ChampionSearchSchema struct {
	Input           string `form:"input" json:"input" validate:"max=16"`
	Role            string `form:"role" json:"role" validate:"max=16"`
	DifficultyLevel string `form:"difficulty_level" json:"difficulty_level" validate:"max=16"`
}
