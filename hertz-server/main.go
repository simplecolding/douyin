// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	h := server.Default()

	// 静态资源访问
	h.StaticFile("/public","./biz")
	h.Static("/public","./biz")
	register(h)
	h.Spin()
}
