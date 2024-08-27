package mysql

import "gorm.io/gorm"

var (
	DB *gorm.DB
)

// Init 初始化mysql连接
func Init() {
	/*
		初始化db 如果初始化不成功则直接panic
	*/
}
