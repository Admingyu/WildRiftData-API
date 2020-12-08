package utils

import (
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

//去除填充数据
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// AES CBC模式解密
//key的长度必须是16、24或者32字节，分别用于选择AES-128, AES-192, or AES-256
func AESDecryptCBC(encryptedData, key, iv []byte) (plainData []byte) {
	block, _ := aes.NewCipher(key)
	//AES分组长度为128位，所以blockSize=16，单位字节
	// blockSize := block.BlockSize()
	// blockMode := cipher.NewCBCDecrypter(block, key[:blockSize]) //初始向量的长度必须等于块block的长度16字节
	blockMode := cipher.NewCBCDecrypter(block, iv)
	plainData = make([]byte, len(encryptedData))
	blockMode.CryptBlocks(plainData, encryptedData)
	plainData = PKCS5UnPadding(plainData)
	return plainData

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
