package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterController struct {
}

func (u RegisterController) Register(c *gin.Context) {
	c.HTML(http.StatusOK, "view/register.html", gin.H{
		"title": "注册",
	})
}
