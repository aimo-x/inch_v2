package inchv2

import (
	"inchv2/model"

	"github.com/gin-gonic/gin"
)

// New 运行
func New() {
	err := model.AutoMigrate()
	if err != nil {
		panic(err)
	}
	en := gin.New()
	en.Use(CORSMiddleware)
	en.Any("MP_verify_cHfFAGkBAWWnDXPk.txt", func(c *gin.Context) {
		c.Writer.WriteString("cHfFAGkBAWWnDXPk")
	})

	web(en.Group("v2"))
	api(en.Group("v2/api"))
	web3(en.Group("v3"))
	api3(en.Group("v3/api"))

	if GetConf().IsSsl {
		go RunTLS(en)
	}
	err = en.Run(":" + GetConf().Port)
	if err != nil {
		panic(err)
	}
}

// RunTLS GIN
func RunTLS(en *gin.Engine) {
	err := en.RunTLS(":443", GetConf().SslPem, GetConf().SslKey)
	if err != nil {
		panic(err)
	}
}

// CORSMiddleware ...
func CORSMiddleware(c *gin.Context) {
	c.Writer.Header().Set("taobao", "acad.taobao.com")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Token")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, authtoken, jwt_token, qm_jwt_token, wechat_jwt_token")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}
	c.Next()
}
