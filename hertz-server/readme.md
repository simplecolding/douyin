# http框架生成
```shell
1. 编写hzidl中的文件(注解还没写进去)
2. cd hertz-server
hz new -I ./hzidl -idl hzidl/user/user.proto -module github.com/simplecolding/douyin/hertz-server
hz new -I ./hzidl -idl hzidl/video/video.proto -module github.com/simplecolding/douyin/hertz-server
hz new -I ./hzidl -idl hzidl/favorite/favorite.proto -module github.com/simplecolding/douyin/hertz-server
hz new -I ./hzidl -idl hzidl/comment/comment.proto -module github.com/simplecolding/douyin/hertz-server
go mod tidy -go=1.16 && go mod tidy -go=1.17
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
# test
```shell
# 登陆
curl --location '127.0.0.1:8888/douyin/user/login' \
--form 'username="aaa"' \
--form 'password="aaa"'

# 获得喜欢列表
curl --location --request GET '127.0.0.1:8888/douyin/favorite/list' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGFpbSI6eyJJRCI6NSwiVXNlcm5hbWUiOiJhYWEifSwiZXhwIjoxNjc2NjUxNDMyLCJvcmlnX2lhdCI6MTY3NjY0NzgzMn0.J1yPDATM1ibD5CFOwMuNXsfoXRg0cwwo5bT9S7L-4bU' \
--form 'token="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGFpbSI6eyJJRCI6NSwiVXNlcm5hbWUiOiJhYWEifSwiZXhwIjoxNjc2NjI3NjQzLCJvcmlnX2lhdCI6MTY3NjYyNDA0M30.aLnwESNHl6ViC39Bl74Sx-hUBWObBej8BPIxQgnoe-A"' \
--form 'user_id="5"'
```

## 跨源资源共享
```shell
go get github.com/hertz-contrib/cors
```