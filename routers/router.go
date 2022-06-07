package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	r.LoadHTMLGlob("view/*")
	r.Static("/static", "static")

	// 后台路由
	AdminRoutersInit(r)

	// 前台路由
	DiskRoutersInit(r)

	r.Run(":8989")
}
