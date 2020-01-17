package inchv2

import (
	"fmt"
	"net/http"

	"inchv2/log"

	"github.com/gin-gonic/gin"
)

func rwErr(msg, err interface{}, c *gin.Context) {
	var log = log.New()
	log.Debug(msg, err, c.Request.URL.String())
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": msg, "err": fmt.Sprint(err)})
}
func rwSus(msg, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 1, "msg": msg, "data": data})
}
