package controller

import (
	. "douyin_mine/config"
	"douyin_mine/utils"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"log"
	"time"
)

type User struct {
	User_Id       int    `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	User_Name     string `gorm:"size:32"`
	User_Password string `gorm:"size:256"`
}

type registerResponse struct {
	statusResponse
	User_Id int    `json:"user_id"`
	Token   string `json:"token"`
}

type loginResponse struct {
	statusResponse
	User_Id int    `json:"user_id,omitempty"`
	Token   string `json:"token,omitempty"`
}

type statusResponse struct {
	Status_Code int    `json:"Status_Code"`
	Status_Msg  string `json:"Status_Msg,omitempty"`
}

type UserController struct {
}

func (uc *UserController) PostRegister(ctx iris.Context) mvc.Result {
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")
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
	Database.AutoMigrate(&User{})
	var user *User
	log.Println("查找用户 : " + username)
	Database.Where("user_name = ?", username).First(&user)
	if user != nil {
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
		User_Password: utils.CreatePassword(password),
	}

	Database.Create(&user)
	token := utils.CreateToken(user.User_Id)
	log.Println("创建用户 : " + username + " 用户token:" + token)
	Rdb.Set(RdbContext, token, user.User_Id, 30*time.Minute)
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
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")

	//查找用户
	var user *User
	Database.Where("user_name = ?", username).First(&user)
	if user == nil {
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
	pwdMd5 := utils.CreatePassword(password)
	if pwdMd5 != password {
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
	Rdb.Set(RdbContext, token, user.User_Id, 30*time.Minute)
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
