# 主要是http框架，目前在hertz-server/handler中修改代码，其他不用改，现实现基本功能，后面慢慢调优

|date|name|note|
|---|---|---|
|2023/2/15|1|调试redis|
|2023/2/117|1|静态资源访问,修改后可以用于视频的访问,访问路径ip:port/public/filename,favorite_server的点赞操作|


# Todo
|iterm|note|
|---|---|
|1|引入jwt鉴权|
|2|完善登陆模块|
|3|feed模块|
|4|输入的参数需要校验，要验证给的参数个数|

# 问题
|iterm|note|
|---|---| 
|1|代码很烂|
|2|status_code为0的时候不显示|
|3|video数据库表中没有设置title列|

```golang
func main() {
    h := server.Default()
    // jwt鉴权
    mw.Init()
    // 静态资源访问
    h.Static("/public","./biz")
    register(h)
    h.Spin()
}

```

```shell
# 更新用户相关脚手架
cd xx/hertz-server
hz update -I ./hzidl -idl hzidl/user/user.proto -module github.com/simplecolding/douyin/hertz-server

```