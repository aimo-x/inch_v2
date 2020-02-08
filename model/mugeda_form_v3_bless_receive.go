package model

import (
	"strings"

	"github.com/jinzhu/gorm"
)

// MugedaFormV3BlessReceive 阵营
type MugedaFormV3BlessReceive struct {
	gorm.Model
	BlessID uint   `json:"bless_id"` // 祝福ID
	CampID  uint   `json:"camp_id"`  // 阵营
	OpenID  string `json:"open_id"`  // 接收者
	Invite  string `json:"invite"`   // 被邀请者 (逗号分隔符 限制 4个)
}

// Create 接收祝福
func (fbr *MugedaFormV3BlessReceive) Create() error {
	db, err := db()
	defer db.Close()
	if err != nil {
		return err
	}
	rows := db.Create(&fbr)
	if rows.Error != nil {
		return rows.Error
	}
	return nil
}

// First 查找我的祝福
func (fbr *MugedaFormV3BlessReceive) First(openid interface{}, blessID uint) (b bool, err error) {
	db, err := db()
	defer db.Close()
	if err != nil {
		return
	}
	rows := db.Where("open_id = ? AND bless_id = ?", openid, blessID).First(&fbr)
	if b = rows.RecordNotFound(); b {
		return
	}
	if err = rows.Error; err != nil {
		return
	}
	return
}

// AddInvite 增加被邀请者
func (fbr *MugedaFormV3BlessReceive) AddInvite(blessReceiveIDInt uint, invite string) (b bool, err error) {
	db, err := db()
	defer db.Close()
	if err != nil {
		return
	}
	rows := db.Where("id = ?", blessReceiveIDInt).First(&fbr)
	if b = rows.RecordNotFound(); b {
		return
	}
	if err = rows.Error; err != nil {
		return
	}
	ins := strings.Split(fbr.Invite, ",")
	for _, v := range ins {
		if v == invite {
			return
		}
	}
	// 最大4个助力
	if len(ins) > 3 {
		// err = errors.New("目标已完成助力")
		return
	}
	if len(fbr.Invite) > 0 {
		invite = fbr.Invite + "," + invite
	}
	rows = rows.Update("invite", invite)
	if b = rows.RecordNotFound(); b {
		return
	}
	if err = rows.Error; err != nil {
		return
	}
	return
}
