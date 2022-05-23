package service

import (
	"douyin_mine/config"
	"time"
)

type CommentJSON struct {
	Id          int      `json:"id"`
	User        UserJSON `json:"user"`
	Content     string   `json:"content"`
	Create_date string   `json:"create_date"`
}

type Comment struct {
	Comment_id          int `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Comment_user_id     int
	Comment_video_id    int
	Comment_content     string
	Comment_latest_time time.Time
	Comment_returnid    int
}

func GetCommentList(userid, videoid int) []CommentJSON {
	var comments []Comment
	config.Database.Where("`comment_video_id` = ?", videoid).Find(&comments)
	commentlist := make([]CommentJSON, 0, 30)
	for _, comment := range comments {
		commentlist = append(commentlist, CommentJSON{
			Id:          comment.Comment_id,
			User:        GetUser(userid, comment.Comment_user_id),
			Content:     comment.Comment_content,
			Create_date: comment.Comment_latest_time.Format("2006-01-02 15:04:05"),
		})
	}
	return commentlist
}

func CreateComment(content string, userid, videoid int) (CommentJSON, error) {
	comment := Comment{
		Comment_user_id:     userid,
		Comment_video_id:    videoid,
		Comment_content:     content,
		Comment_latest_time: time.Now(),
		Comment_returnid:    0,
	}
	err := config.Database.Create(&comment).Error
	return CommentJSON{
		Id:          comment.Comment_id,
		User:        GetUser(userid, comment.Comment_user_id),
		Content:     content,
		Create_date: comment.Comment_latest_time.Format("2006-01-02 15:04:05"),
	}, err
}

func DeleteComment(commentid int) bool {
	rowcount := config.Database.Where("`comment_id` = ?", commentid).Delete(Comment{}).RowsAffected
	return rowcount == 1
}
