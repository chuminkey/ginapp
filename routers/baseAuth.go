package routers

import "github.com/gin-gonic/gin"

func Login(c *gin.Context) {
	userName := c.MustGet(gin.AuthUserKey).(string)
	c.String(200, "登录成功！欢迎您%s", userName)
}
