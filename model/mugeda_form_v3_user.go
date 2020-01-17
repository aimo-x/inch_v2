package model

import (
	"github.com/jinzhu/gorm"
)

// MugedaFormV3User 用户注册资料资料
type MugedaFormV3User struct {
	gorm.Model
	CampID   string `json:"camp_id"` // 所在阵营数组 (逗号分隔符)
	AppID    string `json:"app_id"`  // 公众号ID
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
func (uw *MugedaFormV3User) First(openid interface{}) (b bool, err error) {
	db, err := db()
	defer db.Close()
	if err != nil {
		return
	}
	rows := db.Where("open_id = ?", openid).First(&uw)
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
	rows := db.Where("open_id = ?", openid).First(&uw)
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
