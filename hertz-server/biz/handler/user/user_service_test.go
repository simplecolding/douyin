package user

import (
	"encoding/base64"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/ut"
	"math/rand"
	"strings"
	"testing"
	"time"
)

func TestUserInfo(t *testing.T) {
	h := server.Default()
	h.GET("/douyin/user", UserInfo)
	url := "/douyin/user"
	method := "GET"

	reqBody := strings.NewReader("token=IiujEkqTWcmsEuJOLRKcpDQuasgDjernZMwDFuZhNspVJlbyCD&user_id=6")
	w := ut.PerformRequest(h.Engine,
		method,
		url,
		&ut.Body{reqBody,reqBody.Len()},
		ut.Header{"Content-Type", "application/x-www-form-urlencoded"},
	)
	resp := w.Result()
	//assert.DeepEqual(t, 200, resp.StatusCode())
	println(string(resp.Body()))
}

func TestUserLogin(t *testing.T) {
	// jwt鉴权
	//mw.Init()
	h := server.Default()
	h.POST("/douyin/user/login", UserLogin)
	//username := "aaaaaaaa"
	//reqBody := strings.NewReader("username=" + username + "&password=aaaaaaa")
	url := "/douyin/user/login?username=oDwxGVXO&password=aaaaaaa"
	method := "POST"
	w := ut.PerformRequest(h.Engine,
		method,
		url,
		//&ut.Body{reqBody,reqBody.Len()},
		nil,
		ut.Header{"Content-Type", "application/x-www-form-urlencoded"},
	)
	resp := w.Result()
	//assert.DeepEqual(t, 200, resp.StatusCode())
	println("resp", string(resp.Body()))
}

func TestUserRegister(t *testing.T) {
	h := server.Default()
	h.POST("/douyin/user/register", UserRegister)
	rand.Seed(time.Now().UnixNano())
	username, _ := GenerateRandomString(8)
	reqBody := strings.NewReader("username=" + username + "&password=aaaaaaa")
	url := "/douyin/user/register"
	method := "POST"
	w := ut.PerformRequest(h.Engine,
		method,
		url,
		&ut.Body{reqBody,reqBody.Len()},
		ut.Header{"Content-Type", "application/x-www-form-urlencoded"},
	)
	resp := w.Result()
	//assert.DeepEqual(t, 200, resp.StatusCode())
	println(string(resp.Body()))
	//assert.DeepEqual(t, "{\"message\":\"pong\"}", string(resp.Body()))
}

func GenerateRandomString(n int) (string, error) {
	// 生成随机字节数组
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	// 对字节数组进行Base64编码
	return base64.URLEncoding.EncodeToString(b)[:n], nil
}