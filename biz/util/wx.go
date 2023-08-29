package util

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
)

type Result struct {
	OpenId    string    `json:"openId"`
	NickName  string    `json:"nickName"`
	Gender    int64     `json:"gender"`
	City      string    `json:"city"`
	Province  string    `json:"province"`
	Country   string    `json:"country"`
	AvatarUrl string    `json:"avatarUrl"`
	UnionId   string    `json:"unionId"`
	Watermark Watermark `json:"watermark"`
}

type Watermark struct {
	AppId     string `json:"appid"`
	Timestamp int64  `json:"timestamp"`
}

// 官方文档：https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/signature.html#%E5%8A%A0%E5%AF%86%E6%95%B0%E6%8D%AE%E8%A7%A3%E5%AF%86%E7%AE%97%E6%B3%95
// 微信小程序解密算法 AES-128-CBC
func DecryptWXOpenData(app_id, sessionKey, encryptData, iv string) (*Result, error) {
	decodeBytes, err := base64.StdEncoding.DecodeString(encryptData)
	if err != nil {
		return nil, err
	}
	sessionKeyBytes, errKey := base64.StdEncoding.DecodeString(sessionKey)
	if errKey != nil {
		return nil, errKey
	}
	ivBytes, errIv := base64.StdEncoding.DecodeString(iv)
	if errIv != nil {
		return nil, errIv
	}
	dataBytes, errData := aesDecrypt(decodeBytes, sessionKeyBytes, ivBytes)
	fmt.Printf("dataBytes: %v\n", dataBytes)
	if errData != nil {
		return nil, errData
	}
	result := &Result{}

	errResult := json.Unmarshal(dataBytes, result)

	return result, errResult
}

// AES 解密
func aesDecrypt(crypted, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	// 原始数据
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)

	// 去除填充  --- 数据尾端有'/x0e'占位符,去除它
	length := len(origData)
	unp := int(origData[length-1])
	return origData[:(length - unp)], nil
}
