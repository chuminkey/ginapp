package routers

import (
	"ginapp/utils"
	"github.com/gin-gonic/gin"
	"strings"
)

type userInfo struct {
	Name string
	Age  string
}

func LoginToGetJwt(c *gin.Context) {
	name := c.DefaultQuery("name", "")
	age := c.DefaultQuery("age", "18")
	u := &userInfo{
		Name: name,
		Age:  age,
	}
	token := utils.SignDataToToken(u)
	c.JSON(200, gin.H{
		"code":  "200",
		"token": token,
	})
}

func LoginToCheckToken(c *gin.Context) {
	token := c.GetHeader("token")
	sl := strings.Split(token, " ")
	tt := sl[len(sl)-1]
	data, err := utils.DecodeSignStringToData(tt, "1234567aa")
	if err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"err":  err.Error(),
			"tt":   tt,
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": data,
	})
}
