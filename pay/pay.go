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
)

var log = logging.MustGetLogger("pay")

func Wx(context *gin.Context) {
	//TODO  处理微信支付
	log.Infof("wx-name %s",common.WeConfig.Name)
	log.Infof("wx-key %s",common.WeConfig.Key)
	log.Infof("wx-pay %s",time.Now())
	context.JSON(200,gin.H{"pay":"wx","time":time.Now().Format("2006-01-02 15:04:05")})
}

func Ali(context *gin.Context)  {
	//TODO 处理支付宝支付
	log.Infof("ali-pay %s",time.Now())
	context.JSON(200,gin.H{"pay":"ali","time":time.Now().Format("2006-01-02 15:04:05")})
}