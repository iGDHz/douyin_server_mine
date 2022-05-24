package controller

import (
	"douyin_mine/config"
	"douyin_mine/service"
	"douyin_mine/utils"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"time"
)

type FeedController struct {
}

func (fc *FeedController) Get(context iris.Context) mvc.Result {
	latest_time := context.URLParam("latest_time") //时间戳
	token := context.URLParam("token")
	userid, err := utils.CheckToken(token)
	if err == redis.Nil {
		userid = -1
	}
	var videos []service.Video
	err = config.Database.Where("UNIX_TIMESTAMP(`video_latest_time`) < ?", latest_time).Limit(30).Find(&videos).Error
	if err != nil {
		fmt.Printf("%v", err)
	}
	videolist := make([]service.VideoJSON, 0, 30)
	t := time.Now()
	for i, video := range videos {
		if i == 0 {
			t = video.Video_latest_time
		}
		videolist = append(videolist, service.GetVideoJSON(userid, video))
	}

	return mvc.Response{
		Object: feedResponse{
			statusResponse: statusResponse{
				Status_Code: 0,
				Status_Msg:  "ok",
			},
			Video_list: videolist,
			Next_time:  t.Unix() * 1000,
		},
	}
}
