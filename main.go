package main

import (
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/server"
	user "github.com/simplecolding/douyin/kitex_gen/user/userservice"
	"log"
	"net"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8810")
	if err != nil {
		panic(err)
	}
	svr := user.NewServer(new(UserServiceImpl),
		server.WithServiceAddr(addr),
		server.WithMuxTransport(), //开启多路复用
		server.WithLimit(&limit.Option{MaxConnections: 10000, MaxQPS: 1000}), // 限流
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
