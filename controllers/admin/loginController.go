package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LonginController struct {
	BaseController
}

func (u LonginController) Login(c *gin.Context) {
	c.HTML(http.StatusOK, "view/login.html", gin.H{
		"title": "登录",
	})

	u.Success(c)
	println("Login SUCCESS")
}
