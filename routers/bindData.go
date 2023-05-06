package routers

import "github.com/gin-gonic/gin"

type useInfo struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Gender   string `json:"gender"`
}

func BindData(c *gin.Context) {
	var u useInfo
	err := c.Bind(&u)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "绑定失败",
			"err": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "200",
		"data": u,
	})
}
