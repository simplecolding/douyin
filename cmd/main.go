package main

import (
	"fmt"
	_ "github.com/simplecolding/douyin/cmd/orm" //
	"github.com/simplecolding/douyin/cmd/orm/dal"
	"github.com/simplecolding/douyin/cmd/orm/model"
)

func main()  {
	//db, err := gorm.Open(mysql.Open(constants.MySQLDefaultDSN))
	//if err != nil {
	//	panic(fmt.Errorf("连接mysql失败：%w", err))
	//}
	//dal.SetDefault(db)
	//var user = model.FollowUser{UserID: "yanjiang",FollowID: "tingxia",Status: true}
	//err := dal.FollowUser.Create(&user)
	//if err != nil {
	//	panic(fmt.Errorf("创建user失败：%w", err))
	//}
	//acc, err := dal.FollowUser.Where(dal.FollowUser.UserID.Eq("yanjiang")).Take()
	//fmt.Println(acc)

	var user = model.UserAuth{UserName: "jiangyan", Password: "123456", Status: true}
	err := dal.UserAuth.Create(&user)
	if err != nil {
		panic(fmt.Errorf("创建user失败：%w", err))
	}

	var video = model.Video{UID: 1,PlayURL: "/public/a.mp5",CoverURL: "public/a.cover"}
	err = dal.Video.Create(&video)
	if err != nil {
		panic(fmt.Errorf("创建video失败：%w", err))
	}
}