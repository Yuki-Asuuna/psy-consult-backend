package account_manage

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/rand"
	"net/http"
	"psy-consult-backend/tencent-im"
	"psy-consult-backend/tencent-im/usersig"
	"psy-consult-backend/utils/helper"
)

const (
	expire_time = 3600
)

// 在腾讯IMSDK后端中添加一个用户
func AddIMSDKAccount(userID string, userName string, avatarUrl string) error {
	// url = "https://console.tim.qq.com/v4/im_open_login_svc/account_import?sdkappid=88888888&identifier=admin&usersig=xxx&random=99999999&contenttype=json"
	url := "https://console.tim.qq.com/v4/im_open_login_svc/account_import"
	req_body := &AccountImportRequest{
		UserID:  userID,
		Nick:    userName,
		FaceUrl: avatarUrl,
	}
	body, err := json.Marshal(req_body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	query_params := req.URL.Query()
	query_params.Add("sdkappid", helper.I2S(tencent_im.SDKAppID))
	query_params.Add("identifier", tencent_im.AdminAccount)
	user_sig, _ := usersig.GenUserSig(tencent_im.SDKAppID, tencent_im.SDKSecretKey, tencent_im.AdminAccount, expire_time)
	query_params.Add("usersig", user_sig)
	query_params.Add("random", helper.I642S(int64(rand.Uint32())))
	query_params.Add("contenttype", "json")
	req.URL.RawQuery = query_params.Encode()
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	obj := &AccountImportResponse{}
	err = json.Unmarshal(content, obj)
	if err != nil {
		return err
	}
	if obj.ErrorCode != 0 {
		return errors.New(obj.ErrorInfo)
	}
	// success
	return nil
}

func DeleteIMSDKAccount(userID string) error {
	url := "https://console.tim.qq.com/v4/im_open_login_svc/account_delete"
	req_body := &AccountDeleteRequest{
		DeleteItem: []DeleteItem{{UserID: userID}},
	}
	body, err := json.Marshal(req_body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	query_params := req.URL.Query()
	query_params.Add("sdkappid", helper.I2S(tencent_im.SDKAppID))
	query_params.Add("identifier", tencent_im.AdminAccount)
	user_sig, _ := usersig.GenUserSig(tencent_im.SDKAppID, tencent_im.SDKSecretKey, tencent_im.AdminAccount, expire_time)
	query_params.Add("usersig", user_sig)
	query_params.Add("random", helper.I642S(int64(rand.Uint32())))
	query_params.Add("contenttype", "json")
	req.URL.RawQuery = query_params.Encode()
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	obj := &AccountDeleteResponse{}
	err = json.Unmarshal(content, obj)
	if err != nil {
		return err
	}
	if obj.ErrorCode != 0 {
		return errors.New(obj.ErrorInfo)
	}
	if len(obj.ResultItem) == 0 {
		return errors.New("No response")
	}
	if obj.ResultItem[0].ResultCode != 0 {
		return errors.New(obj.ResultItem[0].ResultInfo)
	}
	// success
	return nil
}
