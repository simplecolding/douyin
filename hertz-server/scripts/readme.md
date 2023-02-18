# 利用gorm-gen生成对数据库操作的代码，有防注入等功能优势
```shell
cd douyin/hertz-server/scripts
go run generate.go
```
tips: 生成的代码在/biz/orm/dal和/biz/orm/model