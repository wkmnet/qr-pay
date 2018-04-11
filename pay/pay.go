/**
 * Create with IntelliJ IDEA
 * Project name : qr-pay
 * Package name : 
 * Author : Wukunmeng
 * User : wukm
 * Date : 18-4-10
 * Time : 下午6:37
 * ---------------------------------
 * 
 */
package pay

import (
	"github.com/gin-gonic/gin"
	"github.com/op/go-logging"
	"time"
	"github.com/wkmnet/qr-pay/common"
	"fmt"
	"net/http"
	"io/ioutil"
)

var log = logging.MustGetLogger("pay")

func Wx(context *gin.Context) {
	//TODO  处理微信支付
	log.Infof("wx-name %s",common.WeConfig.AppId)
	log.Infof("wx-key %s",common.WeConfig.Secret)
	log.Infof("wx-pay %s",time.Now())
	var location =	fmt.Sprintf(
		"https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=snsapi_base&state=1",
		common.WeConfig.AppId,"https://dev-zhua-h5.vmovier.cc/go/callback")

	log.Infof("location:%s", location)
	context.Redirect(http.StatusMovedPermanently,location)
	//context.JSON(200,gin.H{"pay":"wx","time":time.Now().Format("2006-01-02 15:04:05")})
}

func Ali(context *gin.Context)  {
	//TODO 处理支付宝支付
	log.Infof("ali-pay %s",time.Now())
	context.JSON(200,gin.H{"pay":"ali","time":time.Now().Format("2006-01-02 15:04:05")})
}

func PayBack(context *gin.Context)  {
	//TODO
	log.Infof("success-html %s",time.Now())
	//context.HTML(200,"success.html",gin.H{"message":"支付成功"})
	data,_ := ioutil.ReadAll(context.Request.Body)
	context.Request.Body.Close()
	log.Infof("body:%s", string(data))
	context.JSON(http.StatusOK, gin.H{"success":true})
}

func PayHtml(context *gin.Context)  {
	//TODO
	log.Infof("pay-html %s",time.Now())
	context.HTML(200,"pay.html",gin.H{})
}