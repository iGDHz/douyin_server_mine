package controller

import (
	. "douyin_mine/config"
	"douyin_mine/service"
	. "douyin_mine/service"
	"douyin_mine/utils"
	"github.com/go-redis/redis/v8"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"log"
	"strconv"
	"time"
)

type UserController struct {
}

func (uc *UserController) PostRegister(ctx iris.Context) mvc.Result {
	username := ctx.URLParam("username")
	password := ctx.URLParam("password")
	if username == "" || password == "" {
		return mvc.Response{
			Object: registerResponse{
				statusResponse: statusResponse{
					Status_Code: 100,
					Status_Msg:  "用户名和密码不能为空",
				},
				User_Id: -1,
				Token:   "",
			},
		}
	}
	//Database.AutoMigrate(&User{}) //若用户表不存在则初始化用户表
	var user *User
	Log.Println("查找用户 : " + username)
	Database.Where("user_name = ?", username).First(&user)
	if user.User_Name != "" {
		log.Println("用户已存在")
		return mvc.Response{
			Object: registerResponse{
				statusResponse: statusResponse{
					Status_Code: 200,
					Status_Msg:  "用户已存在",
				},
				User_Id: -1,
				Token:   "",
			},
		}
	}
	user = &User{
		User_Name:     username,
		User_Password: utils.EncryptString(password),
	}

	Database.Create(&user)
	token := utils.CreateToken(user.User_Id)
	Log.Println("创建用户 : " + username + " 用户token:" + token)
	err := Rdb.Set(RdbContext, token, user.User_Id, 30*time.Minute).Err()
	if err != nil {
		Log.Println("Redis数据库出错")
		panic(err)
	}
	return mvc.Response{
		Object: registerResponse{
			statusResponse: statusResponse{
				Status_Code: 0,
				Status_Msg:  "注册成功",
			},
			User_Id: user.User_Id,
			Token:   token,
		},
	}
}

func (uc *UserController) PostLogin(ctx iris.Context) mvc.Result {
	username := ctx.URLParam("username")
	password := ctx.URLParam("password")

	//查找用户
	var user *User
	rowcount := Database.Where("user_name = ?", username).Find(&user).RowsAffected
	if rowcount == 0 {
		return mvc.Response{
			Object: loginResponse{
				statusResponse: statusResponse{
					Status_Code: 100,
					Status_Msg:  "用户不存在",
				},
			},
		}
	}

	//验证用户密码
	pwdMd5 := utils.EncryptString(password)
	if pwdMd5 != user.User_Password {
		return mvc.Response{
			Object: loginResponse{
				statusResponse: statusResponse{
					Status_Code: 200,
					Status_Msg:  "用户密码错误",
				},
			},
		}
	}

	//生成token并存入redis
	token := utils.CreateToken(user.User_Id)
	err := Rdb.Set(RdbContext, token, user.User_Id, 30*time.Minute).Err()
	if err != nil {
		Log.Println("Redis数据库出错")
		panic(err)
	}
	return mvc.Response{
		Object: loginResponse{
			statusResponse: statusResponse{
				Status_Code: 0,
				Status_Msg:  "登录成功",
			},
			User_Id: user.User_Id,
			Token:   token,
		},
	}
}

func (uc *UserController) Get(ctx iris.Context) mvc.Result {
	token := ctx.URLParam("token") //获取token
	user_id := ctx.URLParam("user_id")
	fromuserid, err := utils.CheckToken(token)
	Log.Println("请求用户（ID）信息" + user_id)
	if err == redis.Nil { //验证token是否有效
		return mvc.Response{
			Object: GetUserResponse{
				statusResponse: statusResponse{
					Status_Code: 100,
					Status_Msg:  "请先登录",
				},
				UserJSON: service.UserJSON{},
			},
		}
	}

	userIntId, _ := strconv.Atoi(user_id) //将参数转化为字符串类型
	usermsg := GetUser(fromuserid, userIntId)
	return mvc.Response{
		Object: GetUserResponse{
			statusResponse: statusResponse{
				Status_Code: 0,
			},
			UserJSON: usermsg,
		},
	}
}
