package routers

import "github.com/gin-gonic/gin"

func GetJson(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"code": "200",
		"html": "<h1>this is h1</h1>",
	})
}

func GetPure(ctx *gin.Context) {
	ctx.PureJSON(200, gin.H{
		"code": 200,
		"html": "<h1>this is h1</h1>",
	})
}

func GetYAML(ctx *gin.Context) {
	ctx.YAML(200, gin.H{
		"code": 200,
		"data": "yaml data",
	})
}
