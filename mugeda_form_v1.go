package inchv2

import (
	"bytes"
	"encoding/base64"
	"inchv2/model"
	"io"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// MugedaFormV1 mugeda
type MugedaFormV1 struct {
	UserWechat *UserWechat
}

// Route mugeda
func (f1 *MugedaFormV1) Route(r *gin.RouterGroup) {
	r.Use(f1.UserWechat.MiddleWare)
	r.POST("", f1.Create)
	r.GET("", f1.First)
}

// Create FormV1 mugeda
func (f1 *MugedaFormV1) Create(c *gin.Context) {
	var f1in model.MugedaFormV1
	err = c.BindJSON(&f1in)
	if err != nil {
		rwErr("系统错误", err, c)
		return
	}
	path, err := f1.base64SaveImage(f1in.PicURL)
	if err != nil {
		rwErr("系统错误", err, c)
		return
	}
	f1in.PicURL = GetConf().Host + "/v2/" + path
	f1in.AppID = f1.UserWechat.AppID
	f1in.UnionID = f1.UserWechat.UnionID
	f1in.OpenID = f1.UserWechat.OpenID
	f1in.Love = 0
	var f1inB model.MugedaFormV1
	b, err := f1inB.First(f1.UserWechat.AppID, f1.UserWechat.UnionID)
	if b {
		// rwErr("没有信息", err, c)
		var usr model.UserWechat
		b, err = usr.First(f1.UserWechat.AppID, f1.UserWechat.UnionID)
		if b || err != nil {
			rwErr("系统错误", err, c)
			return
		}
		f1in.NickName = usr.NickName
		f1in.HeadImg = usr.HeadImg
		err = f1in.Create()
		if err != nil {
			rwErr("系统错误", err, c)
			return
		}
		rwSus("提交成功", f1in, c)
		return
	}
	if err != nil {
		rwErr("系统错误", err, c)
		return
	}
	// 存在 进行更新
	f1in.PicURL = GetConf().Host + "/v2/" + path
	b, err = f1in.Update(f1.UserWechat.AppID, f1.UserWechat.UnionID, map[string]interface{}{"pic_name": f1in.PicName, "text": f1in.Text, "pic_url": f1in.PicURL, "is_cook": f1in.IsCook})
	if b || err != nil {
		rwErr("更新失败", err, c)
		return
	}
	rwSus("完成更新", f1in, c)
}

// First FormV1 mugeda
func (f1 *MugedaFormV1) First(c *gin.Context) {
	var f1in model.MugedaFormV1
	b, err := f1in.First(f1.UserWechat.AppID, f1.UserWechat.UnionID)
	if b {
		rwErr("没有信息", err, c)
		return
	}
	if err != nil {
		rwErr("系统错误", err, c)
		return
	}
	rwSus("查询成功", f1in, c)
}

// base64SaveImage ...
func (f1 *MugedaFormV1) base64SaveImage(data string) (path string, err error) {
	ddd, err := base64.StdEncoding.DecodeString(data) //成图片文件并把文件写入到buffer
	if err != nil {
		return path, err
	}
	bbb := bytes.NewBuffer(ddd)
	um, err := strconv.ParseInt(strconv.Itoa(755), 8, 0)
	if err != nil {
		return path, err
	}
	err = os.MkdirAll("./usr/"+f1.UserWechat.AppID+"/"+f1.UserWechat.UnionID, os.FileMode(um))
	if err != nil {
		return path, err
	}
	path = "./usr/" + f1.UserWechat.AppID + "/" + f1.UserWechat.UnionID + "/" + uuid.New().String() + ".png"

	f, err := os.Create(path)
	defer f.Close()
	if err != nil {
		return path, err
	}
	_, err = io.Copy(f, bbb)
	if err != nil {
		return path, err
	}
	return path, err
}
