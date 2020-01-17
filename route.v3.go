package inchv2

import (
	"github.com/gin-gonic/gin"
)

var f3 MugedaFormV3
var f3u MugedaFormV3User
var f3c MugedaFormV3Camp

func api3(rg *gin.RouterGroup) {
}
func web3(en *gin.RouterGroup) {
	{
		f3.MugedaFormV3User = &f3u
		f3c.MugedaFormV3User = &f3u
		f3.Route(en.Group("mugeda/form/v3"))
	}
	en.Any("/", func(c *gin.Context) {
		c.Writer.WriteString("success")
	})
	en.StaticFS("usr", gin.Dir("./usr", false))
}
