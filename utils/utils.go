package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"wildrift-api/config"
	"wildrift-api/errors"
)

func LikeFormat(s string) string {
	return fmt.Sprintf("%%%s%%", s)
}

func PKCS7Padding(ciphertext []byte) []byte {
	padding := aes.BlockSize - len(ciphertext)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(plantText []byte) []byte {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	return plantText[:(length - unpadding)]
}

// Aes加密
func AesEncrypt(plainData, key, iv []byte) {
	plainData = PKCS7Padding(plainData)
	ciphertext := make([]byte, len(plainData))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, plainData)
	fmt.Printf("%x\n", ciphertext)
}

// Aes解密
func AesDecrypt(enData, key, iv []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(enData, enData)
	enData = PKCS7UnPadding(enData)
	return enData
}

// Code 换取openid和sessionKey
func CodeToJSession(code string) (openID, sessionKey string) {
	// code换取openid
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", config.WECHAT_APP_ID, config.WECHAT_APP_SECRET, code)
	res, err := http.Get(url)
	errors.HandleError("err request code", err)

	// json解析返回数据
	type JSession struct {
		OpenID     string `json:"openid"`
		SessionKey string `json:"session_key"`
		ErrCode    int    `json:"errcode"`
		ErrMsg     string `json:"errmsg"`
	}
	var content []byte
	var userInfo JSession
	defer res.Body.Close()
	content, err = ioutil.ReadAll(res.Body)
	err = json.Unmarshal(content, &userInfo)
	errors.HandleError("Err unmarshal JSession", err)
	log.Println(string(content))

	return userInfo.OpenID, userInfo.SessionKey
}
