package inchv2

import (
	"inchv2/model"

	"github.com/gin-gonic/gin"
)

// MugedaFormV3 mugeda
type MugedaFormV3 struct {
	MugedaFormV3User *MugedaFormV3User
}

// Route mugeda
func (f3 *MugedaFormV3) Route(r *gin.RouterGroup) {
	r.Any("oauth/wechat", f3.MugedaFormV3User.OauthURL)
	r.Any("oauth/wechat/callback", f3.MugedaFormV3User.CallBack)
	r.GET("oauth/wechat/token", f3.MugedaFormV3User.UseCodeToToken) // 使用code 换取登录信息
	r.GET("oauth/wechat/userinfo", f3.MugedaFormV3User.Get)
	r.Use(f3.MugedaFormV3User.MiddleWare)
	r.POST("update/info", f3.MugedaFormV3User.UpdateInfo) // 更新信息
	/*

		r.POST("", f3.Create)
		r.GET("finds", f3.Find)
	*/
}

// Create FormV3 mugeda
func (f3 *MugedaFormV3) Create(c *gin.Context) {
	var f3in model.MugedaFormV3
	f3in.AppID = f3.MugedaFormV3User.AppID
	f3in.UnionID = f3.MugedaFormV3User.UnionID
	f3in.OpenID = f3.MugedaFormV3User.OpenID
	err = f3in.Create()
	if err != nil {
		rwErr("系统错误", err, c)
		return
	}
	rwSus("提交成功", f3in, c)
	return
}

// Find FormV3 mugeda
func (f3 *MugedaFormV3) Find(c *gin.Context) {
	var f3in model.MugedaFormV3
	f3ins, err := f3in.Find(f3.MugedaFormV3User.OpenID)
	if err != nil {
		rwErr("系统错误", err, c)
		return
	}
	rwSus("查询成功", f3ins, c)
}
