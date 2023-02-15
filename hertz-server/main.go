// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/simplecolding/douyin/hertz-server/biz/mw"
)

func main() {
	mw.InitJwt()
	h := server.Default()

	register(h)
	h.Spin()
}