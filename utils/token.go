package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"time"
)

func CreateToken(userid int) string {
	MD5 := md5.New()
	MD5.Write([]byte(strconv.Itoa(time.Now().Minute()) + strconv.Itoa(userid) + "myfirstgoapp"))
	return hex.EncodeToString(MD5.Sum(nil))
}

func CreatePassword(password string) string {
	MD5 := md5.New()
	MD5.Write([]byte(password))
	return hex.EncodeToString(MD5.Sum(nil))
}
