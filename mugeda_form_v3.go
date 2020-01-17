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
func (f2 *MugedaFormV3) Route(r *gin.RouterGroup) {
	r.Any("oauth/wechat", f2.MugedaFormV3User.OauthURL)
	r.Any("oauth/wechat/callback", f2.MugedaFormV3User.CallBack)
	r.GET("oauth/wechat/token", f2.MugedaFormV3User.UseCodeToToken) // 使用code 换取登录信息
	r.GET("oauth/wechat/userinfo", f2.MugedaFormV3User.Get)
	r.Use(f2.MugedaFormV3User.MiddleWare)
	r.POST("updates", f2.MugedaFormV3User.Updates)
	r.POST("", f2.Create)
	r.GET("finds", f2.Find)

}

// Create FormV1 mugeda
func (f2 *MugedaFormV3) Create(c *gin.Context) {
	var f2in model.MugedaFormV3
	f2in.AppID = f2.MugedaFormV3User.AppID
	f2in.UnionID = f2.MugedaFormV3User.UnionID
	f2in.OpenID = f2.MugedaFormV3User.OpenID
	f2in.Text = c.Request.FormValue("text")
	err = f2in.Create()
	if err != nil {
		rwErr("系统错误", err, c)
		return
	}
	rwSus("提交成功", f2in, c)
	return
}

// Find FormV3 mugeda
func (f2 *MugedaFormV3) Find(c *gin.Context) {
	var f2in model.MugedaFormV3
	f2ins, err := f2in.Find(f2.MugedaFormV3User.AppID, f2.MugedaFormV3User.OpenID)
	if err != nil {
		rwErr("系统错误", err, c)
		return
	}
	rwSus("查询成功", f2ins, c)
}
