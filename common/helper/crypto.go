package helper

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func Hash256(str string) string {
	bt := []byte(str)
	hashByte := sha256.Sum256(bt)
	hashStr := fmt.Sprintf("%x", hashByte)
	return hashStr
}

func HmacSha256(data string, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(data))
	return hex.EncodeToString(mac.Sum(nil))
}

func Md5(str string) string {
	bt := []byte(str)
	md5Byte := md5.Sum(bt)
	md5Str := fmt.Sprintf("%x", md5Byte)
	return md5Str
}
