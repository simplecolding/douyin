# 
```shell
hz new -I ./hzidl -idl hzidl/service/service.proto -module github.com/simplecolding/doushen
hz update -I ./hzidl -idl hzidl/user/user.proto
-I:指定import的文件位置

cd xx/doushen
hz new -I ./idl/hzidl -idl idl/hzidl/user/user.proto -module github.com/simplecolding/douyin
hz new -I ./idl/hzidl -idl idl/hzidl/video/video.proto -module github.com/simplecolding/douyin
hz new -I ./idl/hzidl -idl idl/hzidl/favorite/favorite.proto -module github.com/simplecolding/douyin
hz new -I ./idl/hzidl -idl idl/hzidl/comment/comment.proto -module github.com/simplecolding/douyin
```
