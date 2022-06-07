package routers

import (
	"my-network-disk/controllers/disk"

	"github.com/gin-gonic/gin"
)

func DiskRoutersInit(r *gin.Engine) {

	// 根路由
	homeRouter := r.Group("/")
	{
		homeRouter.GET("/", disk.DiskController{}.Index)
	}
	// 前端版本分组路由
	groupRouter := r.Group("/v1")
	{
		groupRouter.GET("/home", disk.DiskController{}.Home)
	}
}
