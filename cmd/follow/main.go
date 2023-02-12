package main

import (
	"fmt"
	_ "github.com/simplecolding/douyin/cmd/follow/orm" //
	"github.com/simplecolding/douyin/cmd/follow/orm/dal"
	"github.com/simplecolding/douyin/cmd/follow/orm/model"
)

func main()  {
	//db, err := gorm.Open(mysql.Open(constants.MySQLDefaultDSN))
	//if err != nil {
	//	panic(fmt.Errorf("连接mysql失败：%w", err))
	//}
	//dal.SetDefault(db)
	var user = model.FollowUser{UserID: "yanjiang",FollowID: "tingxia",Status: true}
	err := dal.FollowUser.Create(&user)
	if err != nil {
		panic(fmt.Errorf("创建user失败：%w", err))
	}
	acc, err := dal.FollowUser.Where(dal.FollowUser.UserID.Eq("yanjiang")).Take()
	fmt.Println(acc)
}