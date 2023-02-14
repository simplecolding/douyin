package main
// 微服务用的，暂时不用管
import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/connpool"
	"github.com/simplecolding/douyin/kitex_gen/user"
	"github.com/simplecolding/douyin/kitex_gen/user/userservice"
	"log"
	"time"
)

func main()  {
	c, err := userservice.NewClient("user",
		client.WithHostPorts("127.0.0.1:8810"),
		// client.WithShortConnection(), // 短连接
		client.WithMuxConnection(1), //开启多路复用
		client.WithLongConnection(
			connpool.IdleConfig{10, 1000, 100, time.Minute}),
	)
	if err != nil {
		log.Fatal(err)
	}
	req := &user.DouyinUserLoginRequest{Username: "yanjiang_test",Password: "123456"}
	resp, err := c.LoginUser(context.Background(), req,
		callopt.WithRPCTimeout(3*time.Second), // 连接超时
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)

	reqereg := &user.DouyinUserRegisterRequest{Username: "yanjiang_test",Password: "123456"}
	respreg, err := c.RegisterUser(context.Background(), reqereg, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(respreg)
}