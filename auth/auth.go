/**
 * Create with IntelliJ IDEA
 * Project name : qr-pay
 * Package name : 
 * Author : Wukunmeng
 * User : wukm
 * Date : 18-4-11
 * Time : 上午11:21
 * ---------------------------------
 * 
 */
package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/op/go-logging"
	"net/http"
	"fmt"
	"github.com/wkmnet/qr-pay/common"
	"io/ioutil"
	"encoding/json"
	"encoding/xml"
	"time"
	"strings"
	"io"
	"os"
	"strconv"
)

var log = logging.MustGetLogger("auth")

func AuthCallback(context *gin.Context) {
	//TODO  处理微信支付
	h := context.Request.Header
	for k,v := range h{
		log.Infof("key:%s value:%s", k, v)
	}
	log.Infof("Params len:%d", len(context.Params))
	ps := context.Request.URL.Query()
	for key,value := range ps {
		log.Infof("key:%s  value:%s", key, value[0])
	}
	var accessTokenUrl = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code";
	var url = fmt.Sprintf(accessTokenUrl,common.WeConfig.AppId,common.WeConfig.Secret,context.Query("code"))
	log.Infof("accessToken url:%s", url)
	resp,_ := http.Get(url)
	data,_ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	log.Infof("request-data:%s",string(data))
	var accessToken AccessToken
	json.Unmarshal(data,&accessToken)
	formatData,_ := json.MarshalIndent(accessToken,"","    ")
	log.Infof("accessToken:%s", string(formatData))
	if accessToken.ErrCode > 0 {
		context.HTML(http.StatusOK,"error.html",gin.H{"error":accessToken.ErrCode,"message":accessToken.ErrMsg})
		return
	}
	nonceStr := common.RandomString(20)
	var tradeNo = time.Now().Format("20160102150405")
	order := UnifiedOrder{xml.Name{"xml","xml"},common.WeConfig.AppId,common.WeConfig.MerchantId,"WEB",
		nonceStr,"","MD5","测试订单",tradeNo,1,"192.168.8.36",
							"https://dev-zhua-h5.vmovier.cc/go/payback","JSAPI",accessToken.Openid}
	var signStr = SignUnifiedOrder(common.WeConfig.PaySecret,&order)
	order.Sign = signStr;
	req,_ := xml.MarshalIndent(order,"","")
	log.Infof("signStr:%s",signStr)
	log.Infof("unified-order:%s",string(req))
	var orderUrl = "https://api.mch.weixin.qq.com/pay/unifiedorder"
	orderResp,_ := http.Post(orderUrl,"",strings.NewReader(string(req)))
	orderRespBody,_ := ioutil.ReadAll(orderResp.Body)
	orderResp.Body.Close()
	log.Infof("unified-order-response:%s",string(orderRespBody))
	context.JSON(http.StatusOK,gin.H{"ok":true,"openid":accessToken.Openid,"accessToken":accessToken.AccessToken})
}

func startElement(doc string) (element xml.StartElement) {
	dec := xml.NewDecoder(strings.NewReader(doc))
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			return
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			return tok
		case xml.EndElement:

		case xml.CharData:

		}
	}
	return
// containsAll reports
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn uint32 `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Openid string `json:"openid"`
	Scope string `json:"scope"`
	ErrCode uint16 `json:"errcode"`
	ErrMsg string `json:"errmsg"`
}

type UnifiedOrder struct {
	XMLName xml.Name `xml:"xml"`

	AppId string `xml:"appid"`
	MerchantId string `xml:"mch_id"`
	DeviceInfo string `xml:"device_info"`
	NonceStr string `xml:"nonce_str"`
	Sign string `xml:"sign"`
	SignType string `xml:"sign_type"`
	Body string `xml:"body"`
	OutTradeNo string `xml:"out_trade_no"`
	TotalFee uint32 `xml:"total_fee"`
	Address string `xml:"spbill_create_ip"`
	NotifyUrl string `xml:"notify_url"`
	TradeType string `xml:"trade_type"`
	Openid string `xml:"openid"`
}


func SignUnifiedOrder(key string,order *UnifiedOrder) string {
	var res = fmt.Sprintf("appid=%s&body=%s&device_info=%s&mch_id=%s&nonce_str=%s&notify_url=%s&openid=%s&out_trade_no=%s&sign_type=%s&spbill_create_ip=%s&total_fee=%s&trade_type=%s&key=%s",
		order.AppId,order.Body,order.DeviceInfo,order.MerchantId,order.NonceStr,
		order.NotifyUrl,order.Openid,order.OutTradeNo,order.SignType,order.Address,strconv.FormatInt(int64(order.TotalFee),10),order.TradeType,key)
	return strings.ToLower(common.Md5(res))
}