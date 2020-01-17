package inchv2

import (
	"github.com/gin-gonic/gin"
)

var uw UserWechat
var f1 MugedaFormV1
var f2 MugedaFormV2
var f2u MugedaFormV2User

func api(rg *gin.RouterGroup) {
	rg.Any("oauth/wechat", uw.OauthURL)
	rg.Any("oauth/wechat/callback", uw.CallBack)
	rg.GET("oauth/wechat/userinfo", uw.Get)

}
func web(en *gin.RouterGroup) {
	en.Any("qr/odmls", func(c *gin.Context) {
		c.Redirect(302, "https://6.u.mgd5.com/c/5z0l/au_t/index.html")
	})
	{
		f1.UserWechat = &uw
		f1.Route(en.Group("mugeda/form/v1"))
	}
	{
		f2.MugedaFormV2User = &f2u
		f2.Route(en.Group("mugeda/form/v2"))
	}
	en.Any("/", func(c *gin.Context) {
		c.Writer.WriteString("success")
	})
	en.StaticFS("usr", gin.Dir("./usr", false))
}
