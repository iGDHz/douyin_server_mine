package service

type User struct {
	User_id       uint   `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	User_name     string `gorm:"size:64"`
	User_password string `gorm:"size:256"`
}
