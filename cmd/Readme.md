# 关注数据库操作
- 使用gorm gen自动生成代码
```sql
fid,primary key ,
user_id,
follow_id,
status,
create_at
update_at
```
## 获取gorm-gen
```shell
go get -u gorm.io/gen
```
## 执行建表sql
## 编写generate.go
```golang
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
		OutPath: "../dal",
		ModelPkgPath: "../dal/model",
		Mode: gen.WithoutContext|gen.WithDefaultQuery|gen.WithQueryInterface, // generate mode
	})
	g.UseDB(gormdb) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	g.ApplyBasic(g.GenerateModel("follow_user")) // 生成对应数据库的操作代码

	// Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	//g.ApplyInterface(func(method FUserMethod) {}, g.GenerateModel("follow_user"))

	// Generate the code
	g.Execute()
}
```
## 运行generate.go
```shell
cd douyin/scripts/follow/cmd/generate.go
go run generate.go
```

## 生成代码使用
```go
package main

import (
	"fmt"
	"github.com/simplecolding/douyin/cmd/follow/dal"
	"github.com/simplecolding/douyin/cmd/follow/dal/model"
	"github.com/simplecolding/douyin/pkg/constants"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main()  {
	db, err := gorm.Open(mysql.Open(constants.MySQLDefaultDSN))
	if err != nil {
		panic(fmt.Errorf("连接mysql失败：%w", err))
	}
	dal.SetDefault(db)
	var user = model.FollowUser{UserID: "yanjiang",FollowID: "tingxia",Status: true}
	err = dal.FollowUser.Create(&user)
	println(err)
}
```

# 服务层
调用dal层的接口就行，为微服务提供服务接口