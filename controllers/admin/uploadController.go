package admin

import (
	"fmt"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

type UploadController struct{}

func (u UploadController) AddFile(c *gin.Context) {
	username := c.PostForm("username")
	filename, err := c.FormFile("filename")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	// 上传文件到指定目录
	dst := path.Join("./static/upload", filename.Filename)
	c.SaveUploadedFile(filename, dst)
	fmt.Println(dst)

	c.HTML(http.StatusOK, "view/file.html", gin.H{
		"message":  fmt.Sprintf("%s upload done!", filename.Filename),
		"username": username,
		"filelist": "文件列表：",
	})
}
