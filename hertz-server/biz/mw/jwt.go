/*
 * Copyright 2022 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package mw

import (
	"context"
	"errors"
	"github.com/simplecolding/douyin/hertz-server/biz/orm/dal"

	"github.com/simplecolding/douyin/hertz-server/biz/model/hertz/user"
	"github.com/simplecolding/douyin/hertz-server/biz/redis"
	"net/http"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/jwt"
)

var (
	JwtMiddleware *jwt.HertzJWTMiddleware
	IdentityKey   = "claim"
)
type Claim struct {
	ID       int64
	Username string
}
func Init() {
	var err error
	authMiddleware, err := jwt.New(&jwt.HertzJWTMiddleware{
		Realm:         "test zone",
		Key:           []byte("secret key"),
		Timeout:       10*time.Hour,
		MaxRefresh:    time.Hour,
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			var req user.DouyinUserLoginRequest
			if err := c.BindAndValidate(&req); err != nil {
				return
			}
			u, err := dal.UserAuth.Where(dal.UserAuth.UserName.Eq(req.Username)).Take()
			if err != nil {
				return
			}
			if code == 200 {
				// store to redis
				redis.RD.Set(ctx, token, u.UID,0)
				println(redis.RD.Get(ctx, token).Val())
			}
			c.JSON(http.StatusOK, user.DouyinUserLoginResponse{
				StatusCode: int32(code),
				Token: token,
				StatusMsg: "success",
				UserId: u.UID,
			})
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			println("asdofjhh")
			var req user.DouyinUserLoginRequest
			err := c.BindAndValidate(&req)
			println("jjkjkjk")
			if  err != nil {
				return nil, err
			}


			users, err := dal.UserAuth.Where(dal.UserAuth.UserName.Eq(req.Username)).Take()
			if err != nil {
				return nil, err
			}
			if users.Password != req.Password {
				return nil, errors.New("wrong password")
			}

			return Claim{ID: users.UID,Username: users.UserName}, nil
		},
		IdentityKey: IdentityKey,
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			claim := claims[IdentityKey].(map[string]interface{})
			return &Claim{
				ID:       int64(claim["ID"].(float64)),
				Username: claim["Username"].(string),
			}
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(Claim); ok {
				return jwt.MapClaims{
					IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			hlog.CtxErrorf(ctx, "jwt biz err = %+v", e.Error())
			hlog.CtxErrorf(ctx, "jwt biz err = %+v", e.Error())
			return e.Error()
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(http.StatusOK, utils.H{
				"code":    code,
				"message": message,
			})
		},
	})
	if err != nil {
		panic(err)
	}
	JwtMiddleware = authMiddleware
}
