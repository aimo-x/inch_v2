package model

import (
	"github.com/jinzhu/gorm"
)

// MugedaFormV1 用户的基本微信资料
type MugedaFormV1 struct {
	gorm.Model
	AppID    string `json:"app_id"` // 公众号
	UnionID  string `json:"union_id"`
	OpenID   string `json:"open_id"`
	NickName string `json:"nick_name"`
	HeadImg  string `json:"head_img"`
	PicName  string `json:"pic_name"`
	PicURL   string `json:"pic_url"`
	Text     string `json:"text"`
	Love     int    `json:"love"`
	IsCook   bool   `json:"is_cook"` // 是否喜欢做饭
}

// Create 创建用户
func (uw *MugedaFormV1) Create() error {
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

// First  检查用户 appid and unionid 是否存在
func (uw *MugedaFormV1) First(appid, unionid interface{}) (b bool, err error) {
	db, err := db()
	defer db.Close()
	if err != nil {
		return
	}
	rows := db.Where("app_id = ? AND union_id = ?", appid, unionid).First(&uw)
	if b = rows.RecordNotFound(); b {
		return
	}
	if err = rows.Error; err != nil {
		return
	}
	return
}

// Update 更新 appid and unionid 是否存在
func (uw *MugedaFormV1) Update(appid, unionid interface{}, msi map[string]interface{}) (b bool, err error) {
	db, err := db()
	defer db.Close()
	if err != nil {
		return
	}
	rows := db.Model(&uw).Where("app_id = ? AND union_id = ?", appid, unionid).Updates(msi)
	if b = rows.RecordNotFound(); b {
		return
	}
	if err = rows.Error; err != nil {
		return
	}
	return
}
