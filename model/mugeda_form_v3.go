package model

import (
	"github.com/jinzhu/gorm"
)

// MugedaFormV3 用户的基本微信资料
type MugedaFormV3 struct {
	gorm.Model
	AppID    string `json:"app_id"` // 公众号
	UnionID  string `json:"union_id"`
	OpenID   string `json:"open_id"`
	NickName string `json:"nick_name"`
	HeadImg  string `json:"head_img"`
}

// Create 创建用户
func (uw *MugedaFormV3) Create() error {
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
func (uw *MugedaFormV3) First(openid interface{}) (b bool, err error) {
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

// Find 检查用户 appid and openid 是否存在
func (uw *MugedaFormV3) Find(openid interface{}) (mfv2 []MugedaFormV3, err error) {
	db, err := db()
	defer db.Close()
	if err != nil {
		return
	}
	rows := db.Model(&uw).Where("open_id = ?", openid).Offset(0).Limit(10).Order("id DESC").Find(&mfv2)
	if err = rows.Error; err != nil {
		return
	}
	return
}

// Update 更新 appid and openid 是否存在
func (uw *MugedaFormV3) Update(openid interface{}, msi map[string]interface{}) (b bool, err error) {
	db, err := db()
	defer db.Close()
	if err != nil {
		return
	}
	rows := db.Model(&uw).Where("open_id = ?", openid).Updates(msi)
	if b = rows.RecordNotFound(); b {
		return
	}
	if err = rows.Error; err != nil {
		return
	}
	return
}
