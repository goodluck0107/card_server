package wechatv3

import (
	"context"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/wechat/v3"
	"github.com/go-pay/xlog"
	"net/http"
	"time"
)

const (
	CallbackURL = "http://8.134.80.223:9032/callback/v1/payment"
	MchId       = "1720306540"
	AppId       = "wx7cf88f32b3c3ad62"
	SerialNo    = "4C873778AA0BE3C7CF36634353B111735EC45B42"
	APIv3Key    = "ggoldzyyl250702004533070120akaZb"
	PrivateKey  = `-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDTB+juDUFBxW3Z
MVN/TX4NZl0gl4QeU8CBPFNVhSFMP+iExeR3q2vQS8foaPD57FWw9MC2oaoRxTJe
YtEJIWrc63kQ53gqMOMMQz/R/6FsbkEv1sy9lel0DXjy3AGvbCvyde8XUFpVUq5h
bHaCaC7UauBgJT/RcKMwWwPup9OdX6ESyH9gzWCZpmnU+iXlhgKTEpEWV0d6TMMK
7lMYq+CR+8uhRVeuSg4iQULNrDdlckHKM2zy1gqbHZIn6c1ipBNKOI+gc/NDMIgx
XN4QMwUnDHxSd7K5veH9fpdF2zRdljw5wuE4Wq05vb25olijQyz8hpU+hmrrPmfi
ksrsuAz5AgMBAAECggEAL33/ONu4zPF+mDuWm/a3oJJh8PzIWf7lOvB1nLB6Xuuf
C4pxsVDd0aHMWoyLduNuIYXjfPaDDL7BaCZ6/sALia5gr6I4E96uqkrUKRGLpZhm
iTqhFyWOxXRmvNkwc+c3OLi71xZZTgxufDRps015BIsM9fZMF3lu8Gn7R5FzeV3s
hs8cBF/pc5pHY0nWJd3DttqCsRSAG3sHe06Iz1kKsBSLzem+vGHot93835Q7J1iT
uGKUGacrNedTRpBfVslxOcx3A0NOKGS9FloWvBMv1LVH64X4R2/nbEYUIECEE094
lTMp6ckDeas+PWY1zY0RKxin7h9MOPgEi7yazUXmEQKBgQDo+B9q5Doaaq4z9oZS
wH2gwACfHwvP3SKWRorxONzfL4ofu7CYcV4bdUWFkL/ehTGmUdATeTt3a5wHvDoG
2VHpFzARaFNRv/wrNK1W0R+XaYBh/+tHYGVyHlmT4rsCimX08IrcnRUEvDNv6fmM
n2vvVI3woBTks9xqJJwbwZQ5LQKBgQDn5JUAOPQGk/4crIeNX84w5sQ6WizG0JL6
Td+jzHDOyVU5pPN0GMpTUvCVRabsYlBZW5Hx5rBGERHYN5YxXdzQLmMH724s19/T
ZUs8yBOwz2NxCr57DNUHXIYMxYv6Aev65zC1bJ4SHSk+nNAFPttyNeu2J0hvHAGm
eYoN2onqfQKBgQDkesPFP4PEeK/UgoiGDAapauSxKe+ZstTC8Pg/T3c+5A7gxGCT
gUu8Pi0qqyWhhJuG9GHPV2x82Gq0I2P9Z5EvuvAHgnuEh3c2oHkH1hzXkD663hTP
cbjMTPupUAn8meMYb/igGOaOOE1yCtQVmBxxIkn6neUfz03yQ2lex2EpGQKBgQC/
W4h9e4Ib34olnVXqmvGqtvOc94bVtY5kEVkIcQ9yBQBYJj9kQYTMh7fSZnzdui91
3bOsu+Ign7trAkvlhwBNpsm/5Zu0U5v3dTJGAREGqcz0npobLraocXiJF4dwEp/q
F1fBjtVOO1Qqv/qFKZ6rO8W8NeR3E9RkzQzYa8u9fQKBgBHNJT4agqUKgqeTQ6k5
ZuhUQG0x+IUuXMLpsfRp8QT3gervyQtRwN0wMU/HmbPWZsxqvZWARUMWJ1CPuX+z
Ayp3+K2M2J9554pLknrrlEwBvL95tdjoDF7AGFISLb+2ZnjpCAqhYzWe3KdT6rzS
jjF/9zrUbgnlp4ZV+pOlOeLp
-----END PRIVATE KEY-----`
	PublicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA9kURbkO/nyiaqFri02zp
+erb2SK5vehBTA8u+7vFWX6eHwftrwb27Kb7jP/N6iHpqZSxHCgMt+wpm5G23bHh
mD9xe+3cn1Lcx8675lRQ93k3DtYxwmeQScXAiIHbqegkav1T4pO1Cn0nz0IXlwJ7
IEluAHy2WlabHUgwHbxsQRDci2TVV74djbAGf85gkcu4R5UozH7jNlaDfgWxlACX
FjbElWRZB3LvH/7HpPUtgUiMcIjfHW0/hbOYjPK/+AFmOAILfQloatz06F+/2/ai
3byBX+FMnGwixxcBzfoT1+o4P6EZhJUUbDpXq9D6MJ4cpZRgaz5wU/wtb7qhDSZs
hQIDAQAB
-----END PUBLIC KEY-----`
	PublicKeyID = "PUB_KEY_ID_0117203065402025070200181644000260"
)

var (
	client *wechat.ClientV3
)

func init() {
	var err error
	client, err = GetWechatClient()
	if err != nil {
		xlog.Error(err)
		panic(err)
	}
}

func GetWechatClient() (*wechat.ClientV3, error) {
	wxCli, err := wechat.NewClientV3(MchId, SerialNo, APIv3Key, PrivateKey)
	if err != nil {
		xlog.Error(err)
		return nil, err
	}
	err = wxCli.AutoVerifySignByPublicKey([]byte(PublicKey), PublicKeyID)
	if err != nil {
		xlog.Error(err)
		return nil, err
	}

	//// 自定义配置http请求接收返回结果body大小，默认 10MB
	//client.SetBodySize() // 没有特殊需求，可忽略此配置
	//
	//// 设置自定义RequestId生成方法，非必须
	//client.SetRequestIdFunc()

	// 打开Debug开关，输出日志，默认是关闭的
	wxCli.DebugSwitch = gopay.DebugOn

	return wxCli, nil
}

func CreateOrder(ctx context.Context, tradeNo string, desc string, amount int64) {
	expire := time.Now().Add(30 * time.Minute).Format(time.RFC3339)
	// 初始化 BodyMap
	bm := make(gopay.BodyMap)
	bm.Set("appid", AppId).
		Set("mchid", MchId).
		Set("description", desc).
		Set("out_trade_no", tradeNo).
		Set("time_expire", expire).
		Set("notify_url", CallbackURL).
		SetBodyMap("amount", func(bm gopay.BodyMap) {
			bm.Set("total", amount).
				Set("currency", "CNY")
		})
	//.
	//SetBodyMap("payer", func(bm gopay.BodyMap) {
	//	bm.Set("sp_openid", "asdas")
	//})

	wxRsp, err := client.V3TransactionApp(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	if wxRsp.Code == wechat.Success {
		xlog.Debugf("wxRsp: %#v", wxRsp.Response)
		return
	}
	pkMap := client.WxPublicKeyMap()
	// wxPublicKey：微信平台证书公钥内容，通过 client.WxPublicKeyMap() 获取，然后根据 signInfo.HeaderSerial 获取相应的公钥
	err = wechat.V3VerifySignByPK(wxRsp.SignInfo.HeaderTimestamp, wxRsp.SignInfo.HeaderNonce, wxRsp.SignInfo.SignBody, wxRsp.SignInfo.HeaderSignature, pkMap[wxRsp.SignInfo.HeaderSerial])
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Errorf("wxRsp:%s", wxRsp.Error)
}

func Callback(w http.ResponseWriter, r *http.Request) {
	notifyRsp, err := wechat.V3ParseNotify(r)
	if err != nil {
		xlog.Error(err)
		return
	}
}
