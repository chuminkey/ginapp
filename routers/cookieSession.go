package routers

import "github.com/gin-gonic/gin"

func CookieHandler(c *gin.Context) {
	cookie, _ := c.Cookie("name")
	if cookie == "" {
		c.SetCookie("name", "nks", 36000, "/", "192.168.2.3", false, false)
		c.String(200, "第一次访问")
	} else {
		c.String(200, "第er次访问")
	}
}
