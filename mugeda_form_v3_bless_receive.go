package inchv2

import (
	"inchv2/model"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// MugedaFormV3BlessReceive 接收祝福语
type MugedaFormV3BlessReceive struct {
	MugedaFormV3User *MugedaFormV3User
}

// GET 查询祝福语 url bless_id
// 第一次接收祝福，使用 bless_id 和 openid 进行查询 是否存在 存在即返回，不存在即创建
// 查询结果，如果满足4个邀请人，即返回祝福语
func (fc *MugedaFormV3BlessReceive) GET(c *gin.Context) {
	var dfc model.MugedaFormV3BlessReceive
	blessID := c.Request.FormValue("bless_id")
	blessIDInt, err := strconv.Atoi(blessID)
	if err != nil {
		rwErr("error", err, c)
		return
	}
	b, err := dfc.First(fc.MugedaFormV3User.OpenID, uint(blessIDInt))
	if b {
		// 不可以接收自己的祝福语 后期考虑
		//dfc.Invite = ""
		dfc.BlessID = uint(blessIDInt)
		dfc.OpenID = fc.MugedaFormV3User.OpenID
		err = dfc.Create()
		if err != nil {
			rwErr("error", err, c)
			return
		}
		rwSus("success", dfc, c)
		return
	}
	if err != nil {
		rwErr("error", err, c)
		return
	}
	ins := strings.Split(dfc.Invite, ",")
	var data gin.H
	data = gin.H{"mugeda_form_v3_bless_receive": dfc, "mugeda_form_v3_bless": ""}
	if len(ins) > 3 {
		var f3b model.MugedaFormV3Bless
		b, err = f3b.First(dfc.BlessID)
		if err != nil || b {
			rwErr("error", err, c)
			return
		}
		data = gin.H{"mugeda_form_v3_bless_receive": dfc, "mugeda_form_v3_bless": f3b}
	}
	rwSus("success", data, c)
}

// AddInvite 接收祝福 bless_receive_id 助力者登录信息记录
func (fc *MugedaFormV3BlessReceive) AddInvite(c *gin.Context) {
	var dfc model.MugedaFormV3BlessReceive
	blessReceiveID := c.Request.FormValue("bless_receive_id")
	blessReceiveIDInt, err := strconv.Atoi(blessReceiveID)
	if err != nil {
		rwErr("error", err, c)
		return
	}
	b, err := dfc.AddInvite(uint(blessReceiveIDInt), fc.MugedaFormV3User.OpenID)
	if b || err != nil {
		rwErr("error", err, c)
	}
	rwSus("success", dfc, c)
}

// Create 接收祝福 bless_id 【暂时不启用】
// 第一次接收祝福，使用 bless_id 和 openid 进行查询 是否存在 存在即返回，不存在即创建
func (fc *MugedaFormV3BlessReceive) Create(c *gin.Context) {
	var dfc model.MugedaFormV3BlessReceive
	blessID := c.Request.FormValue("bless_id")
	blessIDInt, err := strconv.Atoi(blessID)
	if err != nil {
		rwErr("error", err, c)
		return
	}
	dfc.Invite = ""
	dfc.BlessID = uint(blessIDInt)
	dfc.OpenID = fc.MugedaFormV3User.OpenID
	err = dfc.Create()
	if err != nil {
		rwErr("error", err, c)
		return
	}
	rwSus("success", dfc, c)
}
