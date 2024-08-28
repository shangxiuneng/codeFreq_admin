package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB *gorm.DB
)

// Init 初始化mysql连接
func Init() {
	/*
		初始化db 如果初始化不成功则直接panic
	*/
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:my-secret-pw@tcp(127.0.0.1:3306)/mydatabase?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	DB = db
}

/*
docker run --name my-mysql-container -e MYSQL_ROOT_PASSWORD=my-secret-pw -e MYSQL_USER=myuser -e MYSQL_PASSWORD=mypassword -e MYSQL_DATABASE=mydatabase -p 3306:3306 -d mysql:latest

--name my-mysql-container：为容器指定名称为 my-mysql-container。
-e MYSQL_ROOT_PASSWORD=my-secret-pw：设置 MySQL root 用户的密码为 my-secret-pw。
-e MYSQL_USER=myuser：创建一个用户名为 myuser 的普通用户。
-e MYSQL_PASSWORD=mypassword：为 myuser 用户设置密码为 mypassword。
-e MYSQL_DATABASE=mydatabase：在启动时创建一个名为 mydatabase 的数据库。
-p 3306:3306：将容器内的 MySQL 端口 3306 映射到主机的 3306 端口。你可以根据需要更改主机端口。
-d mysql:latest：以守护进程模式运行 MySQL，并使用最新的 MySQL 官方镜像。
*/
