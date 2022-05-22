package service

type Favorite struct {
	Favorite_user_id int `gorm:"PRIMARY_KEY"`
	Favorite_fan_id  int
}
