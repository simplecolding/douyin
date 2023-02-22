// Code generated by hertz generator.

package favorite

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/simplecolding/douyin/hertz-server/biz/orm/dal"
	"github.com/simplecolding/douyin/hertz-server/biz/orm/model"
	"github.com/simplecolding/douyin/hertz-server/biz/utils"
	"strconv"

	favorite "github.com/simplecolding/douyin/hertz-server/biz/model/hertz/favorite"
)

// FavoriteAction .
// @router /douyin/favorite/action [POST]
func FavoriteAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req favorite.DouyinFavoriteActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest,"参数错误!!!")
	}
	// 鉴权
	flag, _ ,uid := utils.Auth(ctx,req.Token)
	if !flag{
		c.JSON(consts.StatusBadRequest, "token错误")
		return
	}
	//println("uid:",uid,req.VideoId)
	data, err := dal.Favorite.FilterWithVidAndUid(req.VideoId, uid)
	//data, err := dal.Favorite.WithContext(ctx).FilterWithVidAndUid(req.VideoId, uid)
	exitRow := len(data) != 0
	action := req.ActionType
	if action == 1 {
		if exitRow  {
			// 取消点赞
			dal.Favorite.WithContext(ctx).Where(dal.Favorite.Lid.Eq(data[0].Lid)).Update(dal.Favorite.Status, false)
			resp := new(favorite.DouyinFavoriteActionResponse)
			resp.StatusCode = 0
			resp.StatusMsg = "重复favorite"
			c.JSON(consts.StatusOK, resp)
			return
		} else {
			dal.Favorite.Create(&model.Favorite{Vid: req.VideoId, UID: uid, Status: false})
		}
	} else {
		dal.Favorite.WithContext(ctx).Where(dal.Favorite.Lid.Eq(data[0].Lid)).Delete()
	}
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(favorite.DouyinFavoriteActionResponse)
	resp.StatusCode = 0
	resp.StatusMsg = "favorite successfully"
	c.JSON(consts.StatusOK, resp)
}

// GetFavoriteList .
// @router /douyin/favorite/list [GET]
func GetFavoriteList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req favorite.DouyinFavoriteListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	resp := new(favorite.DouyinFavoriteListResponse)
	uid := req.UserId
	// todo 这里其实应该分页limit8;
	data, _ := dal.Favorite.Where(dal.Favorite.UID.Eq(req.UserId)).Order(dal.Favorite.UpdatedAt).Find()
	userInfo := FavoriteQueryUser(uid)
	var video []*favorite.Video
	for _, d := range data {
		videoInfo, err := dal.Video.Where(dal.Video.Vid.Eq(d.Vid)).First()
		if err != nil {
			println(d.Vid)
			c.JSON(consts.StatusBadRequest, "视频查询失败")
		}
		var v favorite.Video
		v.Id = d.Lid // 视频唯一标识
		v.IsFavorite = true
		v.Author = &userInfo
		v.PlayUrl = videoInfo.PlayURL
		v.CoverUrl = videoInfo.CoverURL
		v.FavoriteCount, _ = dal.Favorite.Where(dal.Favorite.Vid.Eq(d.Vid)).Count()
		v.CommentCount, _ = dal.Comment.Where(dal.Comment.Vid.Eq(d.Vid)).Count()
		data, err := dal.Favorite.FilterWithVidAndUid(d.Vid, uid)
		if len(data) > 0{
			v.IsFavorite = true
		}else {
			v.IsFavorite = false
		}
		v.Title = videoInfo.Title
		video = append(video, &v)
	}

	resp.StatusCode = 0
	resp.StatusMsg = "success"
	resp.VideoList = video

	c.JSON(consts.StatusOK, resp)
}

// FavoriteQueryUser query userinfo
func FavoriteQueryUser(uid int64) favorite.User {
	dal.UserAuth.Where(dal.UserAuth.UID.Eq(uid))
	totalFavorited := int64(0)
	v, err := dal.Video.Where(dal.Video.UID.Eq(uid)).Find()
	for _, t := range v {
		tmpcount, _ := dal.Favorite.Where(dal.Favorite.Vid.Eq(t.Vid)).Count()
		totalFavorited += tmpcount
	}
	// 低性能代码
	workCount, _ := dal.Video.Where(dal.Video.UID.Eq(uid)).Count()
	favoriteCount, _ := dal.Favorite.Where(dal.Favorite.UID.Eq(uid)).Count()
	userInfoDB, err := dal.UserAuth.Where(dal.UserAuth.UID.Eq(uid)).First()
	if err != nil {
		println("database err")
	}
	userInfoDB.WorkCount = workCount
	userInfoDB.FavoriteCount = favoriteCount
	userInfoDB.TotalFavorite = strconv.FormatInt(totalFavorited, 10)
	dal.UserAuth.Save(userInfoDB)

	return favorite.User{
		Id:              userInfoDB.UID,
		Name:            userInfoDB.UserName,
		FollowCount:     userInfoDB.FollowCount,
		IsFollow:        userInfoDB.IsFollow,
		Avatar:          userInfoDB.Avatar,
		BackgroundImage: userInfoDB.BackgroundImage,
		Signature:       userInfoDB.Signature,
		TotalFavorited:  userInfoDB.TotalFavorite,
		WorkCount:       userInfoDB.WorkCount,
		FavoriteCount:   userInfoDB.FavoriteCount,
	}
}
