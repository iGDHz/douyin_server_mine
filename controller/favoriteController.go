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

type videolistResponse struct {
	statusResponse
	Video_list []service.VideoJSON `json:"video_list"`
}

type watch_praise struct {
	Praise_user_id  int
	Praise_video_id int
}

type FavoriteController struct {
}

func (fc *FavoriteController) PostAction(context iris.Context) mvc.Result {
	token := context.URLParam("token")
	_, err := Rdb.Get(RdbContext, token).Result() //用户id
	if err == redis.Nil {
		return mvc.Response{
			Object: statusResponse{
				Status_Code: 100,
				Status_Msg:  "请先登录",
			},
		}
	}
	userid := context.URLParam("user_id")          //用户id
	videoid := context.URLParam("video_id")        //视频id
	action_type := context.URLParam("action_type") //1为点赞 2为取消赞
	var praise bool
	switch action_type {
	case "1":
		{
			praise = true //点赞操作
		}
	case "2":
		{
			praise = false //取消赞操作
		}
	}

	var favorite *watch_praise
	var exists = false
	Database.Where("`praise_user_id` = ? and `praise_video_id` = ?", userid, videoid).First(&favorite)
	if favorite.Praise_user_id == 0 {
		exists = true
	}
	//但点赞操作已经存在点赞列表
	if exists == praise {
		return mvc.Response{Object: statusResponse{
			Status_Code: 400,
			Status_Msg:  "你已经点赞/取消赞了",
		}}
	}
	if praise {
		uid, _ := strconv.Atoi(userid)
		vid, _ := strconv.Atoi(videoid)
		Database.Create(watch_praise{
			Praise_user_id:  uid,
			Praise_video_id: vid,
		})
	} else {
		Database.Delete(&favorite)
	}
	return mvc.Response{
		Object: statusResponse{
			Status_Code: 0,
			Status_Msg:  "点赞/取消赞成功",
		},
	}
}

func (fc *FavoriteController) GetList(context iris.Context) mvc.Result {
	token := context.URLParam("token")
	userid, err := utils.CheckToken(token)
	if err == redis.Nil {
		return mvc.Response{
			Object: videolistResponse{
				statusResponse: statusResponse{
					Status_Code: 0,
					Status_Msg:  "请先登录",
				},
			},
		}
	}
	var videos []service.Video
	err = Database.Where("`video_authorid` = ?", userid).Find(&videos).Error
	if err != nil {
		Log.Printf("%v", err)
	}
	videolist := make([]service.VideoJSON, 0, 30)
	for _, video := range videos {
		videolist = append(videolist, service.GetVideoJSON(userid, video))
	}

	return mvc.Response{
		Object: videolistResponse{
			statusResponse: statusResponse{
				Status_Code: 0,
				Status_Msg:  "ok",
			},
			Video_list: videolist,
		},
	}
}