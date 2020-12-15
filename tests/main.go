package main

import (
	"encoding/base64"
	"log"
	"wildrift-api/utils"
)

func main() {

	en := "Rs5bsjau0alJF5//La8rmaJ1k3PNEuAdZxkYtHO45hUmagcpENhOYbvXdG4ZzVeH2jUILybZfchN+4x9KIVHFfF9cZKqFvbJ3ZB9CmipRPdX3kXbTi6tWpOEzNDbdduJm4OCiwXf60Lv7XQaDJDwzOQSKPeuLdJzdVHyCFzPPY4tZjGmBv91kV95fcjUj3HX5yfIw/MOWc4NdG7Qro56dydBUz56eRgbgTGOasQ9S0o3PfxuqbA5FWiNJKYMQt36j901cknX1TFzwMmZqBPnZzZX+4nT/ZOgWbUzzTr63pCJlIOKRAK79ThT0/eHB9J1xA7t5QnrbgdynvNSDUyOCT9vEwvVVabGA73Xnsx6wBvJmHpp/ual6ANPpwX/21WnNgpILqHPqyWDHfL06BEslML6tUaBD5ZM7FX39/0YQSyWmT8vLwVrPwoAnEDz+dmeFY/5TPg2gkGxe/yOK0BYBA=="
	iv := "EYOkMI2PUjMreKr08xAc5A=="
	sign := "96e4b954723a8454d098bb2e4d1fcb4e6a42dc17"
	enB, _ := base64.StdEncoding.DecodeString(en)
	enKey, _ := base64.StdEncoding.DecodeString(sign)
	eniv, _ := base64.StdEncoding.DecodeString(iv)
	res := utils.AesDecrypt(enB, enKey, eniv)
	log.Println(res)
}
