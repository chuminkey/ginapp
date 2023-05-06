package routers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type fromData struct {
	Name       string `json:"name"`
	Password2  string `json:"password_2"`
	RePassword string `json:"repassword"`
	InviteCode string `json:"invite_code"`
	Action     string `json:"action"`
}

func FromOtherApi(c *gin.Context) {
	url := "https://api.uomg.com/api/rand.qinghua"
	client := &http.Client{}
	var fromD fromData
	c.Bind(&fromD)
	marshal, _ := json.Marshal(fromD)
	// 请求体添加
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(marshal))

	// 请求头添加
	req.Header.Add("Referer", "http://dongti233.com/")
	req.Header.Add("Origin", "http://dongti233.com/")
	req.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")

	resp, err := client.Do(req)
	if err != nil {
		c.String(500, "请求失败"+err.Error())
		return
	}
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.String(500, "读取失败"+err.Error())
		return
	}
	// ok 可以顺利返回情话api 调用第三方完成
	c.String(200, string(res))
}
