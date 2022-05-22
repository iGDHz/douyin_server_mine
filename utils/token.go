package utils

import (
	"crypto/md5"
	"douyin_mine/config"
	"encoding/hex"
	"strconv"
	"time"
)

func CreateToken(userid int) string {
	MD5 := md5.New()
	MD5.Write([]byte(strconv.Itoa(time.Now().Minute()) + strconv.Itoa(userid) + "myfirstgoapp"))
	return hex.EncodeToString(MD5.Sum(nil))
}

func EncryptString(str string) string {
	MD5 := md5.New()
	MD5.Write([]byte(str))
	return hex.EncodeToString(MD5.Sum(nil))
}

//检查token 返回id值和bool判断是否存在
func CheckToken(token string) (int, error) {
	target := config.Rdb.Get(config.RdbContext, token)
	id, _ := strconv.Atoi(target.Val())
	return id, target.Err()
}
