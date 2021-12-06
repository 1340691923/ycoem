package channel

import (
	"errors"
	"fmt"
	"net/url"

	"ycoem/lib/logs"

	"ycoem/lib/util"
)

// 微信小程序
type WeChat struct {
	AppID  string
	Secret string
}

const WeChatAPIDomain = "https://api.weixin.qq.com"

func (wc *WeChat) code2Session(jsCode string) (string, string, string, error) {
	var req struct {
		OpenID     string `json:"openid"`
		SessionKey string `json:"session_key"`
		UnionID    string `json:"unionid"`
		ErrCode    int    `json:"errcode"`
		ErrMsg     string `json:"errmsg"`
	}
	var params = url.Values{}
	params.Set("grant_type", "authorization_code")
	params.Set("appid", wc.AppID)
	params.Set("secret", wc.Secret)
	params.Set("js_code", jsCode)
	err := util.GetURLReceiveJSON(WeChatAPIDomain+"/sns/jscode2session", params, &req)
	if err != nil {
		return "", "", "", err
	}
	logs.Logger.Sugar().Infof("WeChat res:", req, params)
	if req.ErrCode != 0 {
		return "", "", "", fmt.Errorf("code2Session failed(%v)", req)
	}
	return req.OpenID, req.SessionKey, req.UnionID, nil
}

func (wc *WeChat) CheckLogin(chanUser, token string) (string, string, error) {
	if token == "" {
		return "", "", errors.New("参数异常")
	}
	openID, _, _, err := wc.code2Session(token)
	return openID, "", err
}
