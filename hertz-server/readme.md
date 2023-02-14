# http框架生成
```shell
1. 编写hzidl中的文件(注解还没写进去)
2. cd hertz-server
hz new -I ./hzidl -idl hzidl/user/user.proto -module github.com/simplecolding/douyin/hertz-server
hz new -I ./hzidl -idl hzidl/video/video.proto -module github.com/simplecolding/douyin/hertz-server
hz new -I ./hzidl -idl hzidl/favorite/favorite.proto -module github.com/simplecolding/douyin/hertz-server
hz new -I ./hzidl -idl hzidl/comment/comment.proto -module github.com/simplecolding/douyin/hertz-server
3. 生成代码,业务代码在biz/handler
4. test(先启动数据库)
 curl --location --request POST '127.0.0.1:8888/douyin/user/register/?username=sgd&password=gf' --header 'User-Agent: Apifox/1.0.0 (https://www.apifox.cn)'
```
## 官方接口地址
https://www.apifox.cn/apidoc/shared-09d88f32-0b6c-4157-9d07-a36d32d7a75c/api-50707521

# ORM数据库
## 启动数据库
```shell
cd docker
docker-compose up
```
## 连接数据库建表，127.0.0.1：3306 账号：gorm，密码：gorm，数据库：gorm
执行./hertz-server/scripts/follow.sql(可以用navicat)
## 生成orm操作代码
```shell
cd ./hertz-server/scripts
go run generate.go
```