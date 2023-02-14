package user

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"testing"
)

func TestUserLogin(t *testing.T) {
	type args struct {
		ctx context.Context
		c   *app.RequestContext
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{

		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestUserRegister(t *testing.T) {
	type args struct {
		ctx context.Context
		c   *app.RequestContext
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name:"t1",
			args: args{ctx: context.Background(), },
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}