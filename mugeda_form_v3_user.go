package inchv2

import (
	"encoding/base64"
	"errors"
	"fmt"
	"inchv2/conf"
	"inchv2/jwt"
	"inchv2/model"
	"inchv2/wechat"
	"inchv2/wechat/cache"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

// MugedaFormV3User ...
type MugedaFormV3User struct {
	UnionID string
	OpenID  string
	AppID   string
}

// MiddleWare 中间件
func (uw *MugedaFormV3User) MiddleWare(c *gin.Context) {
	jwtToken := c.GetHeader("wechat_jwt_token")
	if len(jwtToken) < 20 {
		rwErr("没有授权信息", errors.New("not find wechat_jwt_token"), c)
		c.Abort()
		return
	}
	var j jwt.WeChat
	unionid, openid, appid, err := j.Verify(jwtToken)
	if err != nil {
		rwErr("验证授权失败", err, c)
		c.Abort()
		return
	}
	uw.UnionID = unionid
	uw.OpenID = openid
	uw.AppID = appid
	c.Next()
}

// Get 中间件
func (uw *MugedaFormV3User) Get(c *gin.Context) {
	jwtToken := c.GetHeader("wechat_jwt_token")
	if len(jwtToken) < 20 {
		rwErr("没有授权信息", errors.New("not find wechat_jwt_token"), c)
		c.Abort()
		return
	}
	var j jwt.WeChat
	unionid, openid, appid, err := j.Verify(jwtToken)
	if err != nil {
		rwErr("验证授权失败", err, c)
		c.Abort()
		return
	}
	uw.UnionID = unionid
	uw.OpenID = openid
	uw.AppID = appid
	var u model.MugedaFormV3User
	b, err := u.First(openid)
	if b || err != nil {
		rwErr("没有查询到信息", err, c)
		c.Abort()
		return
	}
	rwSus("查询成功", u, c)
}

// CallBack ...
func (uw *MugedaFormV3User) CallBack(c *gin.Context) {
	wx := uw.GetWeChat()
	oauth := wx.GetOauth()
	rat, err := oauth.GetUserAccessToken(c.Request.FormValue("code"))
	if err != nil {
		rwErr("授权错误", err, c)
		return
	}
	userInfo, err := oauth.GetUserInfo(rat.AccessToken, rat.OpenID)
	if err != nil {
		c.Writer.Write([]byte("<title>授权登陆失败</title><h1>" + userInfo.ErrMsg + "</h1>"))
		return
	}
	var in model.MugedaFormV3User
	in.AppID = wx.Context.AppID
	in.UnionID = userInfo.Unionid
	in.OpenID = userInfo.OpenID
	b, err := in.First(in.OpenID)
	if b { // 写入数据库
		in.NickName = base64.StdEncoding.EncodeToString([]byte(userInfo.Nickname))
		in.HeadImg = userInfo.HeadImgURL
		err = in.Create()
		if err != nil {
			c.Writer.Write([]byte("<title>授权登陆失败</title><h1>" + fmt.Sprint(err) + "</h1>"))
			return
		}
	}
	if err != nil {
		c.Writer.Write([]byte("<title>授权登陆失败</title><h1>" + fmt.Sprint(err) + "</h1>"))
		return
	}
	var j jwt.WeChat
	token, err := j.Token(in.UnionID, in.OpenID, wx.Context.AppID)
	if err != nil {
		c.Writer.Write([]byte("<title>授权登陆失败</title><h1>" + fmt.Sprint(err) + "</h1>"))
		return
	}
	code, err := uw.useTokenToCode(token)
	if err != nil {
		c.Writer.Write([]byte("<title>授权登陆失败</title><h1>" + fmt.Sprint(err) + "</h1>"))
		return
	}
	state := c.Request.FormValue("state")
	if strings.Index(state, "?") == -1 {
		c.Redirect(302, state+"?oauth=wechat&&code="+code)
	} else {
		c.Redirect(302, state+"&oauth=wechat&&code="+code)
	}

}

// PUTUserInfo name + phone + address
func (uw *MugedaFormV3User) PUTUserInfo(c *gin.Context) {
	var in model.MugedaFormV3User
	in.AppID = uw.AppID
	in.UnionID = uw.UnionID
	in.OpenID = uw.OpenID
	in.Name = c.Request.FormValue("name")
	in.Phone = c.Request.FormValue("phone")
	in.Address = c.Request.FormValue("address")
	msi := map[string]interface{}{"name": in.Name, "phone": in.Phone, "address": in.Address}
	b, err := in.Updates(in.OpenID, msi)
	if b || err != nil {
		rwErr("系统错误", err, c)
		return
	}
	rwSus("更新成功", in, c)
}

// OauthURL ...
func (uw *MugedaFormV3User) OauthURL(c *gin.Context) {
	wx := uw.GetWeChat()
	oauth := wx.GetOauth()
	var redirectURI, scope, state = "https://www.inch.online/v3/mugeda/form/v3/oauth/wechat/callback", "snsapi_userinfo", c.Request.FormValue("state")
	// var redirectURI, scope, state = "https://t.iuu.pub/v3/api/oauth/wechat/callback", "snsapi_userinfo", c.Request.FormValue("state")
	uri, err := oauth.GetRedirectURL(redirectURI, scope, state)
	if err != nil {
		rwErr("获取授权地址错误", err, c)
		return
	}
	rwSus("获取授权地址成功", uri, c)
}

// useTokenToCode 使用token 存入兑换码
func (uw *MugedaFormV3User) useTokenToCode(token string) (code string, err error) {
	client := redis.NewClient(conf.Redis())
	_, err = client.Ping().Result()
	if err != nil {
		return code, err
	}
	code = RandomCode(16) + strconv.FormatInt(time.Now().Unix(), 10)
	_, err = client.Set("useCodeToToken"+code, token, time.Minute*5).Result()
	if err != nil {
		return code, err
	}
	return code, err
}

// UseCodeToToken 使用code 换取token
func (uw *MugedaFormV3User) UseCodeToToken(c *gin.Context) {
	code := c.Request.FormValue("code")
	client := redis.NewClient(conf.Redis())
	_, err := client.Ping().Result()
	if err != nil {
		rwErr("系统错误", err, c)
		return
	}
	token, err := client.Get("useCodeToToken" + code).Result()
	if err != nil {
		rwErr("系统错误", err, c)
		return
	}
	_, err = client.Del("useCodeToToken" + code).Result()
	if err != nil {
		rwErr("系统错误", err, c)
		return
	}
	rwSus("获取成功", token, c)
}

// RandomCode 随机码
func (uw *MugedaFormV3User) RandomCode(n int) string {
	str := "0123456789asdfghjklqwertyuiopzxcvbnm"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// GetWeChat 示例
func (uw *MugedaFormV3User) GetWeChat() (wx *wechat.Wechat) {
	var opts cache.RedisOpts
	opts.Host = conf.Redis().Addr
	opts.Password = conf.Redis().Password
	opts.Database = 1
	Redis := cache.NewRedis(&opts)
	var cfg wechat.Config
	cfg.AppID = "wxbdb9cd64895da3d3"                   // "wxa67a64f664dfba26"                   // conf.GetConf().WeChat.AppID         // "wxa67a64f664dfba26"
	cfg.AppSecret = "25295943ffeaa1e9e8b7de4c8588eaf0" // "2ddcf1edc8a54ca2de8b2ebd8f15fcca" // conf.GetConf().WeChat.AppSecret //  "2ddcf1edc8a54ca2de8b2ebd8f15fcca"
	cfg.Cache = Redis
	wx = wechat.NewWechat(&cfg)
	return wx
}

// GetWeChatHeadIMG 获取头像集合 open_id_arr ,分隔符
func (uw *MugedaFormV3User) GetWeChatHeadIMG(c *gin.Context) {
	var u model.MugedaFormV3User
	openidArr := c.Request.FormValue("open_id_arr")
	us, err := u.FindHeadIMG(openidArr)
	if err != nil {
		rwErr("error", err, c)
		return
	}
	rwSus("success", us, c)
}
