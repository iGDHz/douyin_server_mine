package controller

import (
	. "douyin_mine/config"
	"douyin_mine/service"
	"github.com/go-redis/redis/v8"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"strconv"
)

type CommentController struct {
}

func (cc *CommentController) GetList(context iris.Context) mvc.Result {
	token := context.URLParam("token")
	value, err := Rdb.Get(RdbContext, token).Result()
	if err == redis.Nil {
		return mvc.Response{Object: statusResponse{
			Status_Code: 100,
			Status_Msg:  "请先登录",
		}}
	}
	userid, _ := strconv.Atoi(value)
	videoid, _ := strconv.Atoi(context.URLParam("video_id"))
	return mvc.Response{
		Object: commentListResponse{
			statusResponse: statusResponse{
				Status_Code: 0,
			},
			Comment_list: service.GetCommentList(userid, videoid),
		},
	}
}

func (cc *CommentController) PostAction(context iris.Context) mvc.Result {
	token := context.URLParam("token")
	err := Rdb.Get(RdbContext, token).Err()
	if err == redis.Nil {
		return mvc.Response{Object: statusResponse{
			Status_Code: 100,
			Status_Msg:  "请先登录",
		}}
	}
	userid, _ := strconv.Atoi(context.URLParam("user_id"))
	videoid, _ := strconv.Atoi(context.URLParam("video_id"))
	bool := context.URLParam("action_type") == "1" //bool为true时为发表评论操作
	if bool {
		content := context.URLParam("comment_text")
		comments, err := service.CreateComment(content, userid, videoid)
		if err != nil {
			return mvc.Response{
				Object: statusResponse{
					Status_Code: 500,
					Status_Msg:  "数据库创建评论出错",
				},
			}
		}
		return mvc.Response{
			Object: commentResponse{
				statusResponse: statusResponse{
					Status_Code: 0,
					Status_Msg:  "发表成功",
				},
				Comment: comments,
			},
		}
	} else {
		comment_id, _ := strconv.Atoi(context.URLParam("comment_id"))
		if service.DeleteComment(comment_id) {
			return mvc.Response{
				Object: commentResponse{
					statusResponse: statusResponse{
						Status_Code: 0,
						Status_Msg:  "删除成功",
					},
					Comment: service.CommentJSON{}, //删除要返回什么数据？
				},
			}
		} else {
			return mvc.Response{
				Object: commentResponse{
					statusResponse: statusResponse{
						Status_Code: 502,
						Status_Msg:  "从数据库删除失败",
					},
				},
			}
		}
	}
}
