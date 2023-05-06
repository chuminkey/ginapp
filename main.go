package main

import (
	"ginapp/middlewares"
	"ginapp/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	app := gin.Default()
	//err := app.SetTrustedProxies([]string{"192.168.0.102"})
	//if err != nil {
	//	fmt.Println("err", err)
	//	return
	//}
	//fmt.Println("app set proxies ok")
	app.GET("/test", func(ctx *gin.Context) {
		ctx.String(200, "Test ok %s", "zzz")
	})
	app.GET("/redirect", routers.RedirectToBaidu)

	app.GET("/redirect2", routers.RedirectToInternalWeb(app))

	app.GET("/three", routers.ReadDataFromThree)

	app.GET("/json", routers.GetJson)
	app.GET("/pure", routers.GetPure)
	app.GET("/yaml", routers.GetYAML)
	app.GET("/file", routers.GetFile)

	app.POST("/upload", routers.UploadFile)

	app.GET("/login", middlewares.AuthMiddleware(), routers.Login)

	app.POST("/add-user", routers.BindData)

	app.POST("/from-other-api", routers.FromOtherApi)

	app.GET("/cookie", routers.CookieHandler)

	app.GET("/jwt-login", routers.LoginToGetJwt)
	app.POST("/jwt-login", routers.LoginToCheckToken)

	_ = app.Run(":7894")
}
