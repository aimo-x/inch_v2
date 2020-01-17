package inchv2

import (
	"inchv2/model"

	"github.com/gin-gonic/gin"
)

// MugedaFormV2 mugeda
type MugedaFormV2 struct {
	MugedaFormV2User *MugedaFormV2User
}

// Route mugeda
func (f2 *MugedaFormV2) Route(r *gin.RouterGroup) {
	r.Any("oauth/wechat", f2.MugedaFormV2User.OauthURL)
	r.Any("oauth/wechat/callback", f2.MugedaFormV2User.CallBack)
	r.GET("oauth/wechat/token", f2.MugedaFormV2User.UseCodeToToken) // 使用code 换取登录信息
	r.GET("oauth/wechat/userinfo", f2.MugedaFormV2User.Get)
	r.Use(f2.MugedaFormV2User.MiddleWare)
	r.POST("updates", f2.MugedaFormV2User.Updates)
	r.POST("", f2.Create)
	r.GET("finds", f2.Find)

}

// Create FormV1 mugeda
func (f2 *MugedaFormV2) Create(c *gin.Context) {
	var f2in model.MugedaFormV2
	f2in.AppID = f2.MugedaFormV2User.AppID
	f2in.UnionID = f2.MugedaFormV2User.UnionID
	f2in.OpenID = f2.MugedaFormV2User.OpenID
	f2in.Text = c.Request.FormValue("text")
	err = f2in.Create()
	if err != nil {
		rwErr("系统错误", err, c)
		return
	}
	rwSus("提交成功", f2in, c)
	return
}

// Find FormV2 mugeda
func (f2 *MugedaFormV2) Find(c *gin.Context) {
	var f2in model.MugedaFormV2
	f2ins, err := f2in.Find(f2.MugedaFormV2User.AppID, f2.MugedaFormV2User.OpenID)
	if err != nil {
		rwErr("系统错误", err, c)
		return
	}
	rwSus("查询成功", f2ins, c)
}
