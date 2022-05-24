package controller

import "douyin_mine/service"

type commentResponse struct {
	statusResponse
	Comment service.CommentJSON `json:"comment"`
}

type commentListResponse struct {
	statusResponse
	Comment_list []service.CommentJSON `json:"comment_list"`
}

type videolistResponse struct {
	statusResponse
	Video_list []service.VideoJSON `json:"video_list,omitempty"`
}

type feedResponse struct {
	statusResponse
	Video_list []service.VideoJSON `json:"video_list,omitempty"`
	Next_time  int64               `json:"next_time"`
}

type listResponse struct {
	statusResponse
	Video_list []service.VideoJSON `json:"video_list"`
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

type GetUserResponse struct {
	statusResponse
	service.UserJSON
}

type GetUserListResponse struct {
	statusResponse
	User_list []service.UserJSON `json:"user_list"`
}

type statusResponse struct {
	Status_Code int    `json:"status_code"`
	Status_Msg  string `json:"status_msg,omitempty"`
}
