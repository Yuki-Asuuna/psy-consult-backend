package tencent_wechat

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func WeChatLogin(appid string, jscode string) (*WeChatLoginResponse, error) {
	// GET https://api.weixin.qq.com/sns/jscode2session?appid=APPID&secret=SECRET&js_code=JSCODE&grant_type=authorization_code
	url := "https://api.weixin.qq.com/sns/jscode2session"
	body := new(bytes.Buffer)
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	query_params := req.URL.Query()
	query_params.Add("appid", appid)
	query_params.Add("secret", AppSecret)
	query_params.Add("js_code", jscode)
	query_params.Add("grant_type", "authorization_code")
	req.URL.RawQuery = query_params.Encode()
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	obj := &WeChatLoginResponse{}
	err = json.Unmarshal(content, obj)
	if err != nil {
		return nil, err
	}
	if obj.ErrCode != 0 {
		return nil, errors.New(obj.ErrMsg)
	}
	return obj, nil
}
