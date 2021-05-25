package schema

type TemperatureSchema struct {
	Name string `json:"name" binding:"required"`
	Temp string `json:"temp" binding:"required"`
}

type GetTemperatureSchema struct {
	Name string `json:"name" binding:"required"`
}
