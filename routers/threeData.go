package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ReadDataFromThree(ctx *gin.Context) {
	// 似乎直接返回第三方网页有问题
	url := ctx.DefaultQuery("url", "https://cdn.1231218.xyz/uploads/2023/03/2023030115561866.png")
	response, err := http.Get(url)
	if err != nil || response.StatusCode != http.StatusOK {
		ctx.Status(http.StatusServiceUnavailable)
		return
	}
	body := response.Body
	defer body.Close()
	contentType := response.Header.Get("Content-Type")
	contentLength := response.ContentLength
	ctx.DataFromReader(200, contentLength, contentType, body, nil)
}
