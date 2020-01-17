package model

import (
	"inchv2/conf"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql
	//	_ "github.com/jinzhu/gorm/dialects/sqlite" // 正确的包
)

func db() (db *gorm.DB, err error) {
	//db, err = gorm.Open("sqlite3", "db.db")
	db, err = gorm.Open("mysql", conf.GetConf().Mysql.User+":"+conf.GetConf().Mysql.Password+"@tcp("+conf.GetConf().Mysql.Host+":"+conf.GetConf().Mysql.Port+")/"+conf.GetConf().Mysql.Name+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return
	}
	/*
		db.AutoMigrate(&MugedaForm{})
		db.AutoMigrate(&Choujiang{})
		db.AutoMigrate(&UserWechat{})
	*/

	return db, err
}

// Gorm ...
func Gorm() (db *gorm.DB, err error) {
	//db, err = gorm.Open("sqlite3", "db.db")
	db, err = gorm.Open("mysql", conf.GetConf().Mysql.User+":"+conf.GetConf().Mysql.Password+"@tcp("+conf.GetConf().Mysql.Host+":"+conf.GetConf().Mysql.Port+")/"+conf.GetConf().Mysql.Name+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return
	}
	/*
		db.AutoMigrate(&MugedaForm{})
		db.AutoMigrate(&Choujiang{})
		db.AutoMigrate(&UserWechat{})
	*/

	return db, err
}

// AutoMigrate ...
func AutoMigrate() error {
	db, err := gorm.Open("mysql", conf.GetConf().Mysql.User+":"+conf.GetConf().Mysql.Password+"@tcp("+conf.GetConf().Mysql.Host+":"+conf.GetConf().Mysql.Port+")/"+conf.GetConf().Mysql.Name+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&UserWechat{}, &MugedaFormV1{}, &MugedaFormV2User{}, &MugedaFormV2{}, &MugedaFormV3User{}, &MugedaFormV3{}).Error
	return err
}
