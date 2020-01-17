package conf

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/go-redis/redis"
	"github.com/patrickmn/go-cache"
)

// Cache 全局缓存
var Cache = cache.New(5*time.Minute, 5*time.Minute)

// Conf struct
type Conf struct {
	Host      string    `json:"host"`
	Port      string    `json:"port"`
	IsSsl     bool      `json:"is_ssl,omitempty"`
	SslKey    string    `json:"ssl_key,omitempty"`
	SslPem    string    `json:"ssl_pem,omitempty"`
	Token     string    `json:"token,omitempty"`
	WeChat    WeChat    `json:"wechat,omitempty"`
	Mysql     Mysql     `json:"mysql,omitempty"`
	BaiduFace BaiduFace `json:"baidu_face,omitempty"`
	YP        YP        `json:"yp,omitempty"`
}

// WeChat 微信配置
type WeChat struct {
	AppID     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
}

// YP message
type YP struct {
	APIKEY string `json:"api_key"`
}

// BaiduFace 百度人脸
type BaiduFace struct {
	Name      string `json:"name"`
	AppID     string `json:"app_id"`
	Appkey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
}

// Mysql struct
type Mysql struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
}

// Redis 配置
func Redis() *redis.Options {
	return &redis.Options{
		Addr:     "118.89.153.31:6379",
		Password: "039213", // no password set
		DB:       0,        // use default DB
	}
}

// GetConf 获取配置信息
func GetConf() *Conf {
	var cf Conf
	var jst JSONStruct
	//下面使用的是相对路径，config.json文件和main.go文件处于同一目录下
	jst.Load("./conf.json", &cf)
	return &cf
}

// JSONStruct ...
type JSONStruct struct {
}

// Load ...
func (jst *JSONStruct) Load(filename string, v interface{}) {
	//ReadFile函数会读取文件的全部内容，并将结果以[]byte类型返回
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, v)
	if err != nil {
		panic(err)
	}
}
