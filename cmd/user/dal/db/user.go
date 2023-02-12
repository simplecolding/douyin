package db

import (
	"context"
	"github.com/simplecolding/douyin/pkg/constants"
	"gorm.io/gorm"
)

type User struct {
	//UserId int `gorm:"primaryKey;AUTO_INCREMENT"`
	//UserName string
	ID int `gorm:"NOT NULL;PRIMARY KEY;AUTO_INCREMENT=1"`
	UserName string `gorm:"NOT NULL;UNIQUE" json:"user_name"`
	Password string `gorm:"NOT NULL;" json:"password"`
	FollowCount int `gorm:"default:0;"`
	FollowerCount int `gorm:"default:0;"`
	Avatar string `gorm:"default:"`
	IsFollow bool `gorm:"default:false"`
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
func RegisterUser(ctx context.Context, user *User) (uid int, err error) {
	err = preTable()
	if err != nil {
		return -1, err
	}
	err = DB.Create(user).Error
	uid = QueryByUserName(user.UserName)
	return uid, err
}

func QueryByUserName(username string) (uid int){
	var user User
	err := DB.Where("UserName = ?", username).Take(&user)
	if err != nil {
		return -1
	}
	uid = int(user.Model.ID)
	return
}
