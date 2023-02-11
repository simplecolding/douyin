package service

import (
	"context"
	"github.com/simplecolding/douyin/cmd/user/dal/db"
	"github.com/simplecolding/douyin/kitex_gen/user"
)

type RegisterUserService struct {
	ctx context.Context
}

func NewRegisterUserService(ctx context.Context) *RegisterUserService{
	return &RegisterUserService{ctx: ctx}
}

func (s *RegisterUserService) RegisterUser(req *user.DouyinUserRegisterRequest) (uid int,err error) {
	uid, err = db.RegisterUser(s.ctx, &db.User{UserName: req.Username, Password: req.Password})
	if err != nil {
		//
	}
	return uid, err
}

