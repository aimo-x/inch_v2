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
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

// UserWechat ...
type UserWechat struct {
	UnionID string
	OpenID  string
	AppID   string
}

// MiddleWare 中间件
func (uw *UserWechat) MiddleWare(c *gin.Context) {
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
func (uw *UserWechat) Get(c *gin.Context) {
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
	var u model.UserWechat
	b, err := u.First(appid, unionid)
	if b || err != nil {
		rwErr("没有查询到信息", err, c)
		c.Abort()
		return
	}
	rwSus("查询成功", u, c)
}

// CallBack ...
func (uw *UserWechat) CallBack(c *gin.Context) {
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
	var in model.UserWechat
	in.AppID = wx.Context.AppID
	in.UnionID = userInfo.Unionid
	in.OpenID = userInfo.OpenID
	b, err := in.First(in.AppID, in.UnionID)
	if b { // 写入数据库
		in.NickName = base64.StdEncoding.EncodeToString([]byte(userInfo.Nickname))
		in.HeadImg = userInfo.HeadImgURL
		in.OpenID = userInfo.OpenID
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
	// code := useTokenToCode(token)
	c.Redirect(302, c.Request.FormValue("state")+"?oauth=wechat&&wechat_jwt_token="+token)
}

// OauthURL ...
func (uw *UserWechat) OauthURL(c *gin.Context) {
	wx := uw.GetWeChat()
	oauth := wx.GetOauth()
	var redirectURI, scope, state = GetConf().Host + "/v2/api/oauth/wechat/callback", "snsapi_userinfo", c.Request.FormValue("state")
	uri, err := oauth.GetRedirectURL(redirectURI, scope, state)
	if err != nil {
		rwErr("获取授权地址错误", err, c)
		return
	}
	rwSus("获取授权地址成功", uri, c)
}

// useTokenToCode 使用token 存入兑换码
func useTokenToCode(token string) (code string, err error) {
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
func UseCodeToToken(c *gin.Context) {
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
func RandomCode(n int) string {
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
func (uw *UserWechat) GetWeChat() (wx *wechat.Wechat) {
	var opts cache.RedisOpts
	opts.Host = conf.Redis().Addr
	opts.Password = conf.Redis().Password
	opts.Database = conf.Redis().DB
	Redis := cache.NewRedis(&opts)
	var cfg wechat.Config
	cfg.AppID = GetConf().WeChat.AppID
	cfg.AppSecret = GetConf().WeChat.AppSecret
	cfg.Cache = Redis
	wx = wechat.NewWechat(&cfg)
	return wx
}
