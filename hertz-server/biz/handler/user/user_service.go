// Code generated by hertz generator.

package user

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	user "github.com/simplecolding/douyin/hertz-server/biz/model/hertz/user"
	"github.com/simplecolding/douyin/hertz-server/biz/mw"
	_ "github.com/simplecolding/douyin/hertz-server/biz/mw"
	_ "github.com/simplecolding/douyin/hertz-server/biz/orm" //
	"github.com/simplecolding/douyin/hertz-server/biz/orm/dal"
	"github.com/simplecolding/douyin/hertz-server/biz/orm/model"
)

// UserLogin .
// @router /douyin/user/login [POST]
func UserLogin(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.DouyinUserLoginRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	mw.JwtMiddleware.LoginHandler(ctx, c)
	//redisCtx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	//defer cancel()
	//resp := new(user.DouyinUserLoginResponse)
	//tk := "hisadfgh"
	//resp.Token = tk
	//redis.RD.Set(ctx, req.Username, tk, 10*time.Minute)
	//println(redis.RD.Get(ctx, req.Username).Val())
	//c.JSON(consts.StatusOK, resp)
}

// UserRegister .
// @router /douyin/user/register [POST]
func UserRegister(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.DouyinUserRegisterRequest
	resp := new(user.DouyinUserRegisterResponse)
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	username := req.Username
	password := req.Password
	// query
	u, err := dal.UserAuth.Where(dal.UserAuth.UserName.Eq(username)).First()
	if u != nil {
		resp.UserId = -1
		resp.StatusCode = 1
		resp.StatusMsg = "failed!  username已存在！！！" + username
		c.JSON(consts.StatusOK, resp)
		return
		//panic(fmt.Errorf("username已存在"))
	}
	userinfo := model.UserAuth{UserName: username, Password: password}
	err = dal.UserAuth.Create(&userinfo)
	if err != nil {
		resp.UserId = -1
		resp.StatusCode = 1
		resp.StatusMsg = "failed!  插入失败！！！"
		c.JSON(consts.StatusOK, resp)
		return
		//panic(fmt.Errorf("插入失败！！！"))
	}
	u, err = dal.UserAuth.Where(dal.UserAuth.UserName.Eq(username)).First()
	if err != nil {
		resp.UserId = -1
		resp.StatusCode = 1
		resp.StatusMsg = "failed!  插入失败！！！"
		c.JSON(consts.StatusOK, resp)
		return
	}
	resp.UserId = u.UID
	resp.StatusCode = 0
	resp.StatusMsg = "success"
	c.JSON(consts.StatusOK, resp)
}

// UserInfo .
// @router /douyin/user [POST]
func UserInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.DouyinUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	u, b := c.Get(mw.JwtMiddleware.IdentityKey)
	println(b)
	println("hhh",u)
	resp := new(user.DouyinUserResponse)
	//a := mw.JwtMiddleware.Authorizator(ctx,c)

	c.JSON(consts.StatusOK, resp)
}
