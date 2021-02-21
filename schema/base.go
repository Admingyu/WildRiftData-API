package schema

type IdSchema struct {
	ID int `form:"id" json:"id" bind:"required,min=1"`
}

type CodeSchema struct {
	Code       string     `form:"code" json:"code" bind:"required,max=32"`
	DeviceInfo DeviceInfo `form:"device" json:"device"`
}

type WiFiInfo struct {
	SSID           string `form:"ssid" json:"ssid" bind:"max=32"`
	BSSID          string `form:"bssid" json:"bssid" bind:"max=32"`
	SignalStrength int    `form:"signalStrength" json:"signalStrength" bind:"max=32"`
	Frequency      int    `form:"frequency" json:"frequency" bind:"max=32"`
}

type DeviceInfo struct {
	BatteryLevel   int    `form:"batteryLevel" json:"ssid" bind:"max=32"`
	Brand          string `form:"brand" json:"brand" bind:"max=32"`
	Language       string `form:"language" json:"language" bind:"max=32"`
	Model          string `form:"model" json:"model" bind:"max=32"`
	Platform       string `form:"platform" json:"platform" bind:"max=32"`
	System         string `form:"system" json:"system" bind:"max=32"`
	Version        string `form:"version" json:"version" bind:"max=32"`
	BenchmarkLevel int    `form:"benchmarkLevel" json:"benchmarkLevel" bind:"max=32"`
}

type UserInfoSave struct {
	RawData    string   `form:"rawData" json:"rawData" bind:"required"`
	OpenID     string   `form:"openID" json:"openID" bind:"required"`
	ClickBoard string   `form:"clickBoard" json:"clickBoard"`
	WiFiInfo   WiFiInfo `form:"WiFi" json:"WiFi"`
}
