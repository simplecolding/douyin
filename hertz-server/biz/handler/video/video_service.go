// Code generated by hertz generator.

package video

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	video "github.com/simplecolding/douyin/hertz-server/biz/model/hertz/video"
	"github.com/simplecolding/douyin/hertz-server/biz/mw"
	"github.com/simplecolding/douyin/hertz-server/biz/orm/dal"
	"github.com/simplecolding/douyin/hertz-server/biz/orm/model"
)

// VideoPublish .
// @router /douyin/publish/action [POST]
func VideoPublish(ctx context.Context, c *app.RequestContext) {
	//req.Data type : bytes
	var err error
	var r video.PublishActionRequest
	var resp video.DouyinPublishActionResponse
	var byteContainer []byte
	err = c.BindAndValidate(&r)
	if err != nil {
		fmt.Println("err in bind", err)
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	println(r.Data)
	// 将*multipart.FileHeader转为byte[]
	file := r.Data
	fileContent, _ := file.Open()
	// defer file.Close()
	byteContainer, err = ioutil.ReadAll(fileContent)
	if err != nil {
		fmt.Println("err in read file, err: ", err)
		c.String(consts.StatusBadRequest, err.Error())
	}
	fmt.Println("len(data): ", len(byteContainer))
	// jwt授权，从token中获取uid
	fmt.Println("get uid:", mw.JwtMiddleware.IdentityKey)

	// !!!!
	// todo 这里mw会报错
	// u, _ := c.Get(mw.JwtMiddleware.IdentityKey)
	// println("video:", u.(*mw.Claim).ID, u.(*mw.Claim).Username)

	//tk := req.Token
	//jwtToken, err := mw.JwtMiddleware.ParseTokenString(tk)
	//tmp := jwt.ExtractClaimsFromToken(jwtToken)
	// 检查文件类型
	// fileType := http.DetectContentType(req.Data)

	// userName := u.(*mw.Claim).Username
	// if len(userName) == 0 {
	// 	c.String(consts.StatusBadRequest, err.Error())
	// 	return
	// }

	// TEST
	fileName := strconv.FormatInt(time.Now().Unix(), 10) + "userName"
	//fileEndings, err := mime.ExtensionsByType(fileType)
	if err != nil {
		println("get filetype failed")
		return
	}

	// if fileType != "video/mp4" {
	// 	resp.StatusCode = int32(1)
	// 	resp.StatusMsg = "video type incorrect"
	// 	c.JSON(consts.StatusOK, resp)
	// 	println("video incorrect, type is ", fileType)
	// 	return
	// }

	filePath := filepath.Join("./biz/public", fileName+".mp4")

	newFile, err := os.Create(filePath)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		println("create file failed")
		return
	}
	defer newFile.Close()
	if _, err := newFile.Write(byteContainer); err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		println("write file failed")
		return
	}

	playUrl := "http://43.139.145.135:7777/public/" + fileName + ".mp4"
	println("playUrl: ", playUrl)
	// test
	err = dal.Video.WithContext(ctx).Create(&model.Video{UID: 1, PlayURL: playUrl, CoverURL: "coverUrl"})
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		println("write to database failed")
		return
	}
	println(filePath)
	resp.StatusCode = 0
	resp.StatusMsg = "success"
	// println(len(req.Data))
	c.JSON(consts.StatusOK, resp)
}

// GetPublishList .
// @router /douyin/publish/list [GET]
func GetPublishList(ctx context.Context, c *app.RequestContext) {

	// todo
	// resp是正常的，应该是playUrl有问题 所以播放不了
	var err error
	var req video.DouyinPublishListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	// jwt授权，从token中获取uid
	// u, _ := c.Get(mw.JwtMiddleware.IdentityKey)
	// println("video list:", u.(*mw.Claim).ID, u.(*mw.Claim).Username)

	uid := req.UserId

	resp := new(video.DouyinFavoriteListResponse)

	data, err := dal.Video.Where(dal.Video.UID.Eq(uid)).Find()
	if err != nil {
		println("query database failed")
		return
	}
	var v []*video.Video
	for _, d := range data {
		var tmp video.Video
		tmp.Id = d.Vid
		tmp.CoverUrl = d.PlayURL
		tmp.PlayUrl = d.PlayURL
		// 还有很多没有实现的

		fmt.Println(tmp)
		fmt.Println("vid: ", d.Vid)
		v = append(v, &tmp)
	}
	resp.VideoList = v
	resp.StatusCode = 0
	resp.StatusMsg = "success"
	c.JSON(consts.StatusOK, resp)
}

// GetFeed .
// @router /douyin/feed [GET]
func GetFeed(ctx context.Context, c *app.RequestContext) {
	var err error
	var req video.DouyinFeedRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(video.DouyinFeedResponse)
	// todo
	// 30 videos for a single time
	limit := 30
	// data, err := dal.Video.Order("created_at desc").Limit(limit).Find()
	// todo
	// 这里不知道Order不知道咋放函数
	data, err := dal.Video.Limit(limit).Find()
	if err != nil {
		println("query database failed")
		return
	}
	resp.StatusCode = 0
	resp.StatusMsg = "success"
	var v []*video.Video
	for _, d := range data {
		var tmp video.Video
		tmp.Id = d.Vid
		tmp.CoverUrl = d.CoverURL
		tmp.PlayUrl = d.PlayURL
		v = append(v, &tmp)
	}
	resp.VideoList = v
	c.JSON(consts.StatusOK, resp)
}
