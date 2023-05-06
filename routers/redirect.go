package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RedirectToBaidu(ctx *gin.Context) {
	ctx.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
}

func RedirectToInternalWeb(app *gin.Engine) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.Request.URL.Path = "/test"
		app.HandleContext(ctx)
	}
}
