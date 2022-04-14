package im_friend

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/rand"
	"net/http"
	tencent_im "psy-consult-backend/tencent-im"
	"psy-consult-backend/tencent-im/usersig"
	"psy-consult-backend/utils/helper"
)

const (
	expire_time = 3600
)

// https://console.tim.qq.com/v4/sns/friend_add?sdkappid=88888888&identifier=admin&usersig=xxx&random=99999999&contenttype=json

func AddFriend(accountA string, accountB string) error {
	// url = "https://console.tim.qq.com/v4/im_open_login_svc/account_import?sdkappid=88888888&identifier=admin&usersig=xxx&random=99999999&contenttype=json"
	url := "https://console.tim.qq.com/v4/sns/friend_add"
	req_body := &AddFriendRequest{
		FromAccount:   accountA,
		AddFriendItem: []AddFriendItem{{ToAccount: accountB, AddSource: "AddSource_Type_Admin", GroupName: "已绑定的督导"}},
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
	obj := &AddFriendResponse{}
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

func DeleteFriend(accountA string, accountB string) error {
	// url = "https://console.tim.qq.com/v4/im_open_login_svc/account_import?sdkappid=88888888&identifier=admin&usersig=xxx&random=99999999&contenttype=json"
	url := "https://console.tim.qq.com/v4/sns/friend_delete"
	req_body := &DeleteFriendRequest{
		FromAccount: accountA,
		ToAccount:   []string{accountB},
		DeleteType:  "Delete_Type_Both",
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
	obj := &DeleteFriendResponse{}
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
