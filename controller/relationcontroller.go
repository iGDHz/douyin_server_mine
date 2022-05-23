package controller

import (
	. "douyin_mine/config"
	"douyin_mine/service"
	"douyin_mine/utils"
	"github.com/go-redis/redis/v8"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"strconv"
)

type RelationController struct {
}

func (rc *RelationController) PostAction(context iris.Context) mvc.Result {
	token := context.URLParam("token")
	userid, err := utils.CheckToken(token) //用户id
	if err == redis.Nil {
		return mvc.Response{
			Object: statusResponse{
				Status_Code: 100,
				Status_Msg:  "请先登录",
			},
		}
	}
	//操作用户id
	touserid := context.URLParam("to_user_id")     //关注用户id
	action_type := context.URLParam("action_type") //1为关注 2为取消关注
	var praise bool
	switch action_type {
	case "1":
		{
			praise = true //关注操作
		}
	case "2":
		{
			praise = false //取消关注操作
		}
	}

	var favorite *service.Favorite
	var exists = false
	rowcount := Database.Where("`favorite_user_id` = ? and `favorite_fan_id` = ?", touserid, userid).First(&favorite).RowsAffected
	if rowcount != 0 {
		exists = true
	}
	//当用户已经关注/取消关注用户
	if exists == praise {
		return mvc.Response{Object: statusResponse{
			Status_Code: 400,
			Status_Msg:  "你已经关注/取关了",
		}}
	}
	if praise {
		touid, _ := strconv.Atoi(touserid)
		Database.Create(service.Favorite{
			Favorite_user_id: touid,
			Favorite_fan_id:  userid,
		})
	} else {
		Database.Delete(&favorite)
	}
	return mvc.Response{
		Object: statusResponse{
			Status_Code: 0,
			Status_Msg:  "关注/取消关注成功",
		},
	}
}

func (rc *RelationController) GetFollowList(context iris.Context) mvc.Result {
	token := context.URLParam("token")
	userid, err := utils.CheckToken(token)
	if err == redis.Nil {
		return mvc.Response{
			Object: statusResponse{
				Status_Code: 100,
				Status_Msg:  "请先登录",
			},
		}
	}
	favoritelist := make([]service.Favorite, 0, 30)
	Database.Where("`favorite_fan_id` = ?", userid).Find(&favoritelist)
	userlist := make([]service.UserJSON, 0, 30)
	for _, favorite := range favoritelist {
		user := service.GetUser(userid, favorite.Favorite_user_id)
		userlist = append(userlist, user)
	}
	return mvc.Response{
		Object: GetUserListResponse{
			statusResponse: statusResponse{
				Status_Code: 0,
			},
			User_list: userlist,
		},
	}
}

func (rc *RelationController) GetFollowerList(context iris.Context) mvc.Result {
	token := context.URLParam("token")
	userid, err := utils.CheckToken(token)
	if err == redis.Nil {
		return mvc.Response{
			Object: statusResponse{
				Status_Code: 100,
				Status_Msg:  "请先登录",
			},
		}
	}
	favoritelist := make([]service.Favorite, 0, 30)
	Database.Where("`favorite_user_id` = ?", userid).Find(&favoritelist)
	userlist := make([]service.UserJSON, 0, 30)
	for _, favorite := range favoritelist {
		user := service.GetUser(userid, favorite.Favorite_fan_id)
		userlist = append(userlist, user)
	}
	return mvc.Response{
		Object: GetUserListResponse{
			statusResponse: statusResponse{
				Status_Code: 0,
			},
			User_list: userlist,
		},
	}
}
