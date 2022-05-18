package service

import (
	. "douyin_mine/config"
)

type User struct {
	User_Id       int    `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	User_Name     string `gorm:"size:64"`
	User_Password string `gorm:"size:256"`
}

type UserJSON struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Follow_count   int    `json:"follow_count"`   //关注总数
	Follower_count int    `json:"follower_count"` //粉丝总数
	Is_follow      bool   `json:"is_follow"`      //true-已关注 false-未关注
}

/*
	@parm:userid 使用用户id
	@parm:touserid 所查询的用户id
*/
func GetUser(userid int, targetuser User) UserJSON {
	var user UserJSON
	count_row := Database.Raw("select count(*) from `favorites` where `favorite_user_id`=?", targetuser.User_Id).Row()
	count_row.Scan(&user.Follower_count)
	count_row = Database.Raw("select count(*) from `favorites` where `favorite_fan_id`=?", targetuser.User_Id).Row()
	count_row.Scan(&user.Follow_count)
	count_row = Database.Raw("select count(*) from `favorites` where favorite_user_id`=? "+
		"abd `favorite_fan_id`=?", targetuser.User_Id, userid).Row()
	count_row.Scan(&user.Is_follow)
	user.Id = targetuser.User_Id
	user.Name = targetuser.User_Name
	return user
}
