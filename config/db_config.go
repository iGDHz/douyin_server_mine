package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"strings"
)

var Database *gorm.DB

func init() {
	db_driver, dsn := getDsn("datasource")
	log.Println("数据库驱动:" + db_driver)
	log.Println("DSN:" + dsn)
	switch db_driver {
	case "mysql":
		{
			var err error
			Database, err = gorm.Open(mysql.Open(dsn)) //连接数据库
			if err != nil {
				panic(err)
			}
		}

	}
}

func getDsn(pre string) (string, string) {
	var dsn string
	drivername := AppConfig.GetString(pre + ".drivername")
	user := AppConfig.GetString(pre + ".user")             //获取用户名
	password := ":" + AppConfig.GetString(pre+".password") //获取密码
	protocol := "@" + AppConfig.GetString(pre+".protocol") //或是使用的协议
	host := "(" + AppConfig.GetString(pre+".host")         //获取地址
	port := ":" + AppConfig.GetString(pre+".port") + ")"   //获取端口号
	dbname := "/" + AppConfig.GetString(pre+".name")       //获取数据库名
	args := AppConfig.GetStringSlice("datasource.args")    //获取配置参数
	dsn = strings.Join([]string{user, password, protocol, host, port, dbname}, "")
	if args != nil {
		dsn = dsn + "?" + strings.Join(args, "&&")
	}
	fmt.Println("%v", args)
	return drivername, dsn
	//args = args.(map[string]{})
}
