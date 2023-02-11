package db

import (
	"context"
	"github.com/simplecolding/douyin/pkg/constants"
	"gorm.io/gorm"
)

type User struct {
	//UserId int `gorm:"primaryKey;AUTO_INCREMENT"`
	//UserName string
	//FollowCount int
	//FollowerCount int
	//IsFollow bool
	UserName string `gorm:"NOT NULL;UNIQUE",json:"user_name"`
	Password string `gorm:"NOT NULL;",json:"password"`
	gorm.Model
}

func (u User) TableName() string {
	return constants.UserTableName
}

func preTable() error {
	if !DB.Migrator().HasTable(&User{}) {
		err := DB.Migrator().CreateTable(&User{})
		if err != nil {
			panic("Create Table Failed")
			return err
		}
	}
	return nil
}
func RegisterUser(ctx context.Context, user []*User) (uid int, err error) {
	err = preTable()
	if err != nil {
		return -1, err
	}
	err = DB.Create(user).Error
	uid = 1
	return uid, err
}

//func QueryByUserName(username string) (uid int ){
//	var user User
//	err := DB.
//	return uid
//}
