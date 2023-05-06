package routers

import "github.com/gin-gonic/gin"

func GetFile(c *gin.Context) {
	c.File("D:\\tmp\\emojisdy\\emojisdy\\0.png")
}
