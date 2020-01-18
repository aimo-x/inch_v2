package model

import (
	"github.com/jinzhu/gorm"
)

// MugedaFormV3Bless 阵营
type MugedaFormV3Bless struct {
	gorm.Model
	Content string `json:"content"` // 祝福内容
	OpenID  string `json:"open_id"` // 创建祝福者
}

// Create 发起祝福
func (fb *MugedaFormV3Bless) Create() error {
	db, err := db()
	defer db.Close()
	if err != nil {
		return err
	}
	rows := db.Create(&fb)
	if rows.Error != nil {
		return rows.Error
	}
	return nil
}

// First 查询祝福
func (fb *MugedaFormV3Bless) First(id uint) (b bool, err error) {
	db, err := db()
	defer db.Close()
	if err != nil {
		return
	}
	rows := db.Where("id = ?", id).First(&fb)
	if b = rows.RecordNotFound(); b {
		return
	}
	if err = rows.Error; err != nil {
		return
	}
	return
}

// FirstID 查询祝福
func (fb *MugedaFormV3Bless) FirstID(id interface{}) (b bool, err error) {
	db, err := db()
	defer db.Close()
	if err != nil {
		return
	}
	rows := db.Where("id = ?", id).First(&fb)
	if b = rows.RecordNotFound(); b {
		return
	}
	if err = rows.Error; err != nil {
		return
	}
	return
}
