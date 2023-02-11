package service

import (
	"context"
	"github.com/simplecolding/douyin/cmd/user/dal"
	"github.com/simplecolding/douyin/kitex_gen/user"
	"testing"
)

func TestRegisterUserService_RegisterUser(t *testing.T) {
	dal.Init()
	type fields struct {
		ctx context.Context
	}
	type args struct {
		req *user.DouyinUserRegisterRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			args: args{req: &user.DouyinUserRegisterRequest{Username: "yanjaing",Password: "123456"}},
		},
		{
			name: "t2",
			args: args{req: &user.DouyinUserRegisterRequest{Username: "jack",Password: "123456"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &RegisterUserService{
				ctx: tt.fields.ctx,
			}
			println(s)
			s.RegisterUser(tt.args.req)
			//if err := s.RegisterUser(tt.args.req); (err != nil) != tt.wantErr {
			//	t.Errorf("RegisterUser() error = %v, wantErr %v", err, tt.wantErr)
			//}
		})
	}
}
