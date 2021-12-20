package core

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/codycoding/goDuck/global"
	"github.com/go-resty/resty/v2"
	"time"
)

type DingTalk struct {
}

//
// TokenRes
//  @Description: Token回复结构
//
type TokenRes struct {
	Errcode     int    `json:"errcode"`      // 错误代码
	AccessToken string `json:"access_token"` // 权限Token
	Errmsg      string `json:"errmsg"`       // 错误信息
	ExpiresIn   int    `json:"expires_in"`   // 过期时间
}

//
// CommonRes
//  @Description: 钉钉API通用返回结构
//
type CommonRes struct {
	Errcode   int    `json:"errcode"`    // 错误码，0为正常
	Errmsg    string `json:"errmsg"`     // 错误信息
	RequestId string `json:"request_id"` // 请求ID
}

//
// GetToken
//  @Description: 获取接口访问Token， 未过期token保存到redis
//  @receiver TokenService
//  @return token
//  @return err
//
func (d *DingTalk) GetToken() (token string, err error) {
	tokenUrl := "https://oapi.dingtalk.com/gettoken"
	// 判断是否有token缓存
	if token, err = global.Redis.Get(context.Background(), global.Config.DingTalk.CacheKey).Result(); err != nil {
		// 没有缓存数据或已过期
		// 从新获取
		client := resty.New()
		var resp *resty.Response
		if resp, err = client.R().SetQueryParams(map[string]string{
			"appkey":    global.Config.DingTalk.AppKey,
			"appsecret": global.Config.DingTalk.AppSecret,
		}).Get(tokenUrl); err != nil {
			// 网络错误
			return token, err
		} else {
			// 解析反馈数据包
			var tokenRes TokenRes
			if err = json.Unmarshal(resp.Body(), &tokenRes); err != nil {
				return token, errors.New("钉钉Token解析失败:" + err.Error())
			}
			// 判断ErrCode
			if tokenRes.Errcode != 0 {
				return token, errors.New("钉钉Token获取失败:" + tokenRes.Errmsg)
			}
			// 写入cache
			if err = global.Redis.Set(context.Background(), global.Config.DingTalk.CacheKey, tokenRes.AccessToken, time.Duration(tokenRes.ExpiresIn)*time.Second).Err(); err != nil {
				return token, errors.New("钉钉Token写入缓存失败:" + err.Error())
			}
			return tokenRes.AccessToken, nil
		}
	} else {
		return token, nil
	}
}

//
// CallPostApi
//  @Description: 调取[Post]API接口
//  @receiver t
//  @param apiUrl api地址
//  @param bodyStruct 请求主体结构
//  @param resBody 返回数据
//  @return err
//
func (d *DingTalk) CallPostApi(apiUrl string, bodyStruct interface{}, resBody *interface{}) (err error) {
	client := resty.New().R()
	var resp *resty.Response
	var token string
	if token, err = d.GetToken(); err != nil {
		return err
	}
	url := apiUrl + "&access_token=" + token
	if resp, err = client.SetBody(bodyStruct).Post(url); err != nil {
		return err
	}
	if err = json.Unmarshal(resp.Body(), resBody); err != nil {
		return err
	}
	return nil
}
