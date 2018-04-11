/**
 * Create with IntelliJ IDEA
 * Project name : qr-pay
 * Package name : 
 * Author : Wukunmeng
 * User : wukm
 * Date : 18-4-10
 * Time : 下午5:23
 * ---------------------------------
 * 
 */
package main

import (
	"github.com/op/go-logging"
	"github.com/gin-gonic/gin"
	"github.com/wkmnet/qr-pay/common"
	"github.com/wkmnet/qr-pay/pay"
	"time"
	"strings"
	"github.com/wkmnet/qr-pay/auth"
)

var log = logging.MustGetLogger("main")

func main() {
	common.LogInit()
	log.Debugf("project start... %s", time.Now())
	engine := gin.Default()
	engine.LoadHTMLGlob("template/*")
	engine.GET("/go/qr",func(context *gin.Context) {
		h := context.Request.Header
		for k,v := range h{
			log.Infof("key:%s value:%s", k, v)
		}
		ua := h.Get("user-agent")
		if strings.Contains(ua,"MicroMessenger") {
			pay.Wx(context)
			return
		}
		if strings.Contains(ua,"AliApp") {
			pay.Ali(context)
			return
		}
		context.JSON(200,gin.H{"ping":"pong","time":time.Now().Format("2006-01-02 15:04:05")})
	})

	engine.GET("/go/callback", auth.AuthCallback)
	engine.GET("/go/static", pay.PayHtml)
	engine.GET("/go/payback", pay.PayBack)
	log.Debugf("project start port: %s", "8080")
	engine.Run()
}
