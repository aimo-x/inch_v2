package inchv2

import (
	"inchv2/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

// MugedaFormV3Bless ...
type MugedaFormV3Bless struct {
	MugedaFormV3User *MugedaFormV3User
}

// GET 查询祝福语 bless_id
func (fc *MugedaFormV3Bless) GET(c *gin.Context) {
	var dfc model.MugedaFormV3Bless
	blessID := c.Request.FormValue("bless_id")
	blessIDInt, err := strconv.Atoi(blessID)
	if err != nil {
		rwErr("error", err, c)
		return
	}
	b, err := dfc.First(uint(blessIDInt))
	if err != nil || b {
		rwErr("error", err, c)
		return
	}
	rwSus("success", dfc, c)
}

// Create 创建祝福语 content
func (fc *MugedaFormV3Bless) Create(c *gin.Context) {
	var dfc model.MugedaFormV3Bless
	dfc.Content = c.Request.FormValue("content")
	dfc.OpenID = fc.MugedaFormV3User.OpenID
	err := dfc.Create()
	if err != nil {
		rwErr("error", err, c)
		return
	}
	rwSus("success", dfc, c)
}