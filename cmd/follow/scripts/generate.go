package main

import (
	"github.com/simplecolding/douyin/pkg/constants"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)


func main() {
	gormdb, _ := gorm.Open(mysql.Open(constants.MySQLDefaultDSN))
	g := gen.NewGenerator(gen.Config{
		OutPath: "../orm/dal",
		ModelPkgPath: "../orm/model",
		Mode: gen.WithoutContext|gen.WithDefaultQuery|gen.WithQueryInterface, // generate mode
	})
	g.UseDB(gormdb) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	g.ApplyBasic(g.GenerateModel("follow_user"))

	//g.ApplyBasic(g.GenerateAllTable()...) // 生成所有表
	// Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	//g.ApplyInterface(func(method FUserMethod) {}, g.GenerateModel("follow_user"))

	// Generate the code
	g.Execute()
}
