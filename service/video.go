package service

import (
	"douyin_mine/config"
	"strings"
	"time"
)

type Video struct {
	Video_id               int `gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	Video_authorid         int
	Video_location         string
	Video_picture_location string
	Video_latest_time      time.Time
	Video_state            int
	Video_category         string
	Video_title            string
	Video_introduction     string
}

type VideoJSON struct {
	Id             int      `json:"id"`
	Author         UserJSON `json:"author"`
	Play_url       string   `json:"play_url"`
	Cover_url      string   `json:"cover_url"`
	Favorite_count int      `json:"favorite_count"`
	Comment_count  int      `json:"comment_count"`
	Is_favorite    bool     `json:"is_favorite"`
	Title          string   `json:"title"`
}

func GetVideoJSON(userid int, video Video) VideoJSON {
	var v VideoJSON
	v.Id = video.Video_id
	//查询用户信息
	v.Author = GetUser(userid, video.Video_authorid)
	v.Play_url = "http://139.9.43.9:8878/douyin/video/" + strings.ReplaceAll(video.Video_location, "\\", "/")
	v.Cover_url = "http://139.9.43.9:8878/douyin/picture/" + strings.ReplaceAll(video.Video_picture_location, "\\", "/")
	//获取视频点赞数
	row := config.Database.Raw("select count(*) from `watch_praises` where `praise_user_id`=? and `praise_video_id`=?", userid, video.Video_id).Row()
	row.Scan(&v.Favorite_count)
	//获取评论总数
	row = config.Database.Raw("select count(*) from `comments` where `comment_video_id`=?", video.Video_id).Row()
	row.Scan(&v.Comment_count)
	//获取是否点赞信息
	var ispraised bool
	row = config.Database.Raw("select count(*) from `watch_praises` where `praise_user_id`=? and `praise_video_id`=? and `praise_ispraised`"+
		"= TRUE", userid, video.Video_id).Row()
	row.Scan(&ispraised)
	v.Title = video.Video_title
	return v
}
