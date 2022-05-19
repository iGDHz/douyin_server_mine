package service

type Comment_praise struct {
	Praise_user_id    int `gorm:"PRIMARY_KEY,AUTO_INCREMENT"`
	Praise_comment_id int
	Praise_ispraised  bool
}
