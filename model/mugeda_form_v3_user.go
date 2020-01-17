package model

import (
	"github.com/jinzhu/gorm"
)

// MugedaFormV3User 用户的基本微信资料
type MugedaFormV3User struct {
	gorm.Model
	AppID    string `json:"app_id"` // 公众号
	UnionID  string `json:"union_id"`
	OpenID   string `json:"open_id"`
	NickName string `json:"nick_name"`
	HeadImg  string `json:"head_img"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}

// Create 创建用户
func (uw *MugedaFormV3User) Create() error {
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

// First 检查用户 appid and openid 是否存在
func (uw *MugedaFormV3User) First(appid, openid interface{}) (b bool, err error) {
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

// Updates 检查用户 appid and openid 更新用户信息 {"name":"","phone":"", "address":"}
func (uw *MugedaFormV3User) Updates(appid, openid interface{}, msi map[string]interface{}) (b bool, err error) {
	db, err := db()
	defer db.Close()
	if err != nil {
		return
	}
	rows := db.Where("app_id = ? AND open_id = ?", appid, openid).First(&uw)
	if b = rows.RecordNotFound(); b {
		return
	}
	rows = rows.Updates(msi)
	if b = rows.RecordNotFound(); b {
		return
	}
	if err = rows.Error; err != nil {
		return
	}
	return
}
