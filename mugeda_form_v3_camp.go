package inchv2

import (
	"inchv2/model"

	"github.com/gin-gonic/gin"
)

// MugedaFormV3Camp ...
type MugedaFormV3Camp struct {
	MugedaFormV3User *MugedaFormV3User
}

// GET 查询福气值
func (fc *MugedaFormV3Camp) GET(c *gin.Context) {
	var dfc model.MugedaFormV3Camp
	b, err := dfc.First(c.Request.FormValue("id"))
	if err != nil || b {
		rwErr("error", err, c)
		return
	}
	rwSus("success", dfc, c)
}
