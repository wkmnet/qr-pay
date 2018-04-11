/**
 * Create with IntelliJ IDEA
 * Project name : qr-pay
 * Package name : 
 * Author : Wukunmeng
 * User : wukm
 * Date : 18-4-10
 * Time : 下午6:30
 * ---------------------------------
 * 
 */
package common

import (
	"github.com/op/go-logging"
	"os"
	"time"
	"io/ioutil"
	"encoding/json"
	"math/rand"
	"crypto/md5"
	"encoding/hex"
)

var log = logging.MustGetLogger("init")

// Example format string. Everything except the message has a custom color
// which is dependent on the log level. Many fields have a custom output
// formatting too, eg. the time returns the hour down to the milli second.
var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

func LogInit()  {
	// For demo purposes, create two backend for os.Stderr.
	backend1 := logging.NewLogBackend(os.Stderr, "", 0)
	backend2 := logging.NewLogBackend(os.Stderr, "", 0)

	// For messages written to backend2 we want to add some additional
	// information to the output, including the used log level and the name of
	// the function.
	backend2Formatter := logging.NewBackendFormatter(backend2, format)

	// Only errors and more severe messages should be sent to backend1
	backend1Leveled := logging.AddModuleLevel(backend1)
	backend1Leveled.SetLevel(logging.DEBUG, "qr.pay")
	// Set the backends to be used.
	logging.SetBackend(backend1Leveled, backend2Formatter)
	log.Infof("init-common-logging at %s", time.Now())
	conf()
}


func RandomString(count int) string{
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < count; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}



func Md5(resource string) string{
	h := md5.New()
	h.Write([]byte(resource)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	log.Infof("md5-resource:%s", resource)
	log.Infof("cipherStr:%s",string(cipherStr))
	log.Infof("md5-result:%s", hex.EncodeToString(cipherStr))
	return string(hex.EncodeToString(cipherStr))
}


type WxConfig struct {
	AppId string `json:"appId"`
	Secret string `json:"secret"`
	MerchantId string `json:"merchantId"`
	PaySecret string `json:"paySecret"`
}

var (
	WeConfig *WxConfig
)

func conf()  {
	f,_ := ioutil.ReadFile("/mnt/conf/wx.json")
	json.Unmarshal(f,&WeConfig)
	log.Infof("read-conf:%s", string(f))
}
