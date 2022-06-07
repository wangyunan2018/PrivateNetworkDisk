package disk

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DiskController struct {
}

func (d DiskController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "view/home.html", gin.H{
		"title": "首页",
	})
}

func (d DiskController) Home(c *gin.Context) {
	c.HTML(http.StatusOK, "view/home.html", gin.H{
		"title": "首页",
	})
}
