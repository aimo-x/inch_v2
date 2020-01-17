package model

import (
	"github.com/jinzhu/gorm"
)

// MugedaFormV2 用户的基本微信资料
type MugedaFormV2 struct {
	gorm.Model
	AppID    string `json:"app_id"` // 公众号
	UnionID  string `json:"union_id"`
	OpenID   string `json:"open_id"`
	NickName string `json:"nick_name"`
	HeadImg  string `json:"head_img"`
	Text     string `json:"text"`
}

// Create 创建用户
func (uw *MugedaFormV2) Create() error {
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

// First  检查用户 appid and openid 是否存在
func (uw *MugedaFormV2) First(appid, openid interface{}) (b bool, err error) {
	db, err := db()
	defer db.Close()
	if err != nil {
		return
	}
	rows := db.Where("app_id = ? AND open_id = ?", appid, openid).First(&uw)
	if b = rows.RecordNotFound(); b {
		return
	}
	if err = rows.Error; err != nil {
		return
	}
	return
}

// Find 检查用户 appid and openid 是否存在
func (uw *MugedaFormV2) Find(appid, openid interface{}) (mfv2 []MugedaFormV2, err error) {
	db, err := db()
	defer db.Close()
	if err != nil {
		return
	}
	rows := db.Model(&uw).Where("app_id = ? AND open_id = ?", appid, openid).Offset(0).Limit(10).Order("id DESC").Find(&mfv2)
	if err = rows.Error; err != nil {
		return
	}
	return
}

// Update 更新 appid and openid 是否存在
func (uw *MugedaFormV2) Update(appid, openid interface{}, msi map[string]interface{}) (b bool, err error) {
	db, err := db()
	defer db.Close()
	if err != nil {
		return
	}
	rows := db.Model(&uw).Where("app_id = ? AND open_id = ?", appid, openid).Updates(msi)
	if b = rows.RecordNotFound(); b {
		return
	}
	if err = rows.Error; err != nil {
		return
	}
	return
}
