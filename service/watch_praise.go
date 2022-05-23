package service

import "douyin_mine/config"

type Watch_praise struct {
	Praise_user_id  int `gorm:"PRIMARY_KEY"`
	Praise_video_id int `gorm:"PRIMARY_KEY"`
}

func GetFavoriteVideoList(userid int) []Video {
	var favorites []Watch_praise
	config.Database.Where("`praise_user_id` = ?", userid).Find(&favorites)
	videolist := make([]Video, len(favorites))
	for i, favorite := range favorites {
		var video Video
		config.Database.Where("`video_id`=?", favorite.Praise_video_id).First(&video)
		videolist[i] = video
	}
	return videolist
}
