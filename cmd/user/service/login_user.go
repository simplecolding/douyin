package service

import (
	"context"
	"github.com/simplecolding/douyin/kitex_gen/user"
)

type LoginUserService struct {
	ctx context.Context
}

func NewLoginUserService(ctx context.Context) *LoginUserService {
	return &LoginUserService{ctx: ctx}
}

func (s *LoginUserService) RegisterUser(req *user.DouyinUserRegisterRequest) error {

	return nil
}