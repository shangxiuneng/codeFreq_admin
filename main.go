package main

import (
	"codeFreq_admin/dal/mysql"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	h := server.Default()

	// 初始化mysql连接
	mysql.Init()

	registerRouter(h)

	h.Spin()
}
