package routers

import (
	"my-network-disk/controllers/admin"

	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(r *gin.Engine) {
	// 新建admin路由组
	adminRouter := r.Group("/admin")
	{
		// 注册路由
		adminRouter.GET("/register", admin.RegisterController{}.Register)
		// 登录路由
		adminRouter.GET("/login", admin.LonginController{}.Login)
		// 文件上传路由
		adminRouter.POST("/upload", admin.UploadController{}.AddFile)
	}
}
