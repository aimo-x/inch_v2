package inchv2

import (
	"inchv2/conf"
)

// GetConf 获取配置信息
func GetConf() *conf.Conf {
	return conf.GetConf()
}
