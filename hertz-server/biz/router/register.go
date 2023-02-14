// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	comment "github.com/simplecolding/douyin/hertz-server/biz/router/comment"
	favorite "github.com/simplecolding/douyin/hertz-server/biz/router/favorite"
	user "github.com/simplecolding/douyin/hertz-server/biz/router/user"
	video "github.com/simplecolding/douyin/hertz-server/biz/router/video"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	comment.Register(r)

	favorite.Register(r)

	video.Register(r)

	user.Register(r)
}
