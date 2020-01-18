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
	r.Any("oauth/wechat", f3.MugedaFormV3User.OauthURL)             // 返回授权地址
	r.Any("oauth/wechat/callback", f3.MugedaFormV3User.CallBack)    // 微信授权回调
	r.GET("oauth/wechat/token", f3.MugedaFormV3User.UseCodeToToken) // 使用code 换取登录信息
	r.GET("oauth/wechat/userinfo", f3.MugedaFormV3User.Get)         // 获取用户信息
	r.Use(f3.MugedaFormV3User.MiddleWare)                           // 用户登陆中间件
	r.PUT("userinfo", f3.MugedaFormV3User.PUTUserInfo)              // 更新用户信息
	r.GET("bless", f3b.GET)                                         // 查询祝福语
	r.PUT("bless", f3b.Create)                                      // 创建祝福语
	r.GET("bless/receive", f3br.GET)                                // 查询是否满足4人
	r.POST("bless/receive/invite", f3br.AddInvite)                  // 助力执行此操作
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
