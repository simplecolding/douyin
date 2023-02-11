package main

import (
	"context"
	"github.com/simplecolding/douyin/cmd/user/service"
	user "github.com/simplecolding/douyin/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// LoginUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) LoginUser(ctx context.Context, req *user.DouyinUserLoginRequest) (resp *user.DouyinUserLoginResponse, err error) {
	// TODO: Your code here...
	resp = new(user.DouyinUserLoginResponse)
	var id int64 = 1
	var stateid int32 = 1
	var token string = "jkjkkj"
	resp.UserId = id
	resp.StatusCode = stateid
	resp.Token = token
	return resp, err
}

// RegisterUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) RegisterUser(ctx context.Context, req *user.DouyinUserRegisterRequest) (resp *user.DouyinUserRegisterResponse, err error) {
	// TODO: Your code here...
	resp = new(user.DouyinUserRegisterResponse)
	resp.UserId = 1
	resp.StatusMsg = "sucess"
	resp.Token = "gjhjfg"
	resp.StatusCode = 0

	uid, err := service.NewRegisterUserService(ctx).RegisterUser(req)
	if err != nil {
		panic(err)
	}
	resp.UserId = int64(uid)

	return resp, nil
}
