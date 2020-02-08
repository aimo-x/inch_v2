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

// Create 创建祝福语 content=&camp_id= 并加入阵营
func (fc *MugedaFormV3Bless) Create(c *gin.Context) {
	var dfc model.MugedaFormV3Bless
	cid, err := strconv.Atoi(c.Request.FormValue("camp_id"))
	if err != nil {
		rwErr("error", err, c)
		return
	}
	dfc.CampID = uint(cid)
	dfc.Content = c.Request.FormValue("content")
	dfc.OpenID = fc.MugedaFormV3User.OpenID
	err = dfc.Create()
	if err != nil {
		rwErr("error", err, c)
		return
	}
	var in model.MugedaFormV3User
	b, err := in.AddCamp(fc.MugedaFormV3User.OpenID, strconv.Itoa(int(dfc.CampID)))
	if b || err != nil {
		rwErr("error", err, c)
		return
	}
	// 增加阵营得分
	var camp model.MugedaFormV3Camp
	b, err = camp.Updates(uint(cid))
	if err != nil || b {
		rwErr("error", err, c)
		return
	}
	rwSus("success", dfc, c)
}
