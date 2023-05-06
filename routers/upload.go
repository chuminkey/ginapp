package routers

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
)

func genMd5(bt []byte) (m5 string) {
	md5New := md5.New()
	md5New.Write(bt)
	return hex.EncodeToString(md5New.Sum(nil))
}

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	ftp, _ := file.Open()
	myByte := make([]byte, file.Size)
	ftp.Read(myByte)
	m5 := genMd5(myByte)
	if err != nil {
		c.String(500, "服务器错误")
		return
	}
	dst := "D:/testzone/golearn/ginapp/uploads/"
	// 此方法会过滤重名文件
	err = c.SaveUploadedFile(file, dst+m5)
	if err != nil {
		c.String(500, "失败")
	} else {
		// 说明已经转存
		c.String(200, "上传成功"+m5)
	}
}
