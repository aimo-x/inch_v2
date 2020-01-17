package model

import (
	"github.com/jinzhu/gorm"
)

// UserWechat 用户的基本微信资料
type UserWechat struct {
	gorm.Model
	AppID    string `json:"app_id"` // 公众号
	UnionID  string `json:"union_id"`
	OpenID   string `json:"open_id"`
	NickName string `json:"nick_name"`
	HeadImg  string `json:"head_img"`
}

// Create 创建用户
func (uw *UserWechat) Create() error {
	db, err := db()
	defer db.Close()
	if err != nil {
		return err
	}
	rows := db.Create(&uw)
	if rows.Error != nil {
		return rows.Error
	}
	return nil
}

// First 检查用户 appid and unionid 是否存在
func (uw *UserWechat) First(appid, unionid interface{}) (b bool, err error) {
	db, err := db()
	defer db.Close()
	if err != nil {
		return
	}
	uw.ID = 0
	rows := db.Where("app_id = ? AND union_id = ?", appid, unionid).First(&uw)
	if b = rows.RecordNotFound(); b {
		return
	}
	if err = rows.Error; err != nil {
		return
	}
	return
}
