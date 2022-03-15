package im_message

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

func SendTextMessage(fromAccount string, toAccount string, msg string) error {
	// url = "https://console.tim.qq.com/v4/im_open_login_svc/account_import?sdkappid=88888888&identifier=admin&usersig=xxx&random=99999999&contenttype=json"
	url := "https://console.tim.qq.com/v4/openim/sendmsg"
	req_body := &SendTextMessageRequest{
		SyncOtherMachine: 1,
		FromAccount:      fromAccount,
		ToAccount:        toAccount,
		MsgRandom:        int64(rand.Uint32()),
		MsgBody:          []TextMessage{{MsgType: "TIMTextElem", MsgContent: TextMessageContent{Text: msg}}},
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
	obj := &SendTextMessageResponse{}
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
