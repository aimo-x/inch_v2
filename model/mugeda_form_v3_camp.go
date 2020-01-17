package model

import (
	"strconv"

	"github.com/jinzhu/gorm"
)

// MugedaFormV3Camp 阵营
type MugedaFormV3Camp struct {
	gorm.Model
	Name  string `json:"name"`  // 阵营名字
	Score int    `json:"score"` // 等分
}

// Create 创建阵营
func (uw *MugedaFormV3Camp) Create() error {
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

// First 查询阵营
func (uw *MugedaFormV3Camp) First(id string) (b bool, err error) {
	db, err := db()
	defer db.Close()
	if err != nil {
		return
	}
	ID, err := strconv.Atoi(id)
	if err != nil {
		return
	}
	rows := db.Where("id = ?", ID).First(&uw)
	if b = rows.RecordNotFound(); b {
		return
	}
	if err = rows.Error; err != nil {
		return
	}
	return
}

// Updates 更新阵营得分 +1
func (uw *MugedaFormV3Camp) Updates(id string) (b bool, err error) {
	db, err := db()
	defer db.Close()
	if err != nil {
		return
	}
	ID, err := strconv.Atoi(id)
	if err != nil {
		return
	}
	rows := db.Where("id = ?", ID).First(&uw)
	if b = rows.RecordNotFound(); b {
		return
	}
	uw.Score = uw.Score + 1
	rows = rows.Update("score", uw.Score)
	if b = rows.RecordNotFound(); b {
		return
	}
	if err = rows.Error; err != nil {
		return
	}
	return
}
