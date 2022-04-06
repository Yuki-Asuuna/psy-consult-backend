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

func CreateNewGroup(visitorAccount, counsellorAccount, groupName string) (string, error) {
	// https://console.tim.qq.com/v4/group_open_http_svc/create_group?sdkappid=88888888&identifier=admin&usersig=xxx&random=99999999&contenttype=json
	url := "https://console.tim.qq.com/v4/group_open_http_svc/create_group"

	req_body := &CreateGroupRequest{
		Name:       groupName,
		Type:       "Public",
		MemberList: []GroupMember{{MemberAccount: visitorAccount}, {MemberAccount: counsellorAccount}},
	}

	body, err := json.Marshal(req_body)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return "", err
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
		return "", err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	obj := &CreateGroupResponse{}
	err = json.Unmarshal(content, obj)
	if err != nil {
		return "", err
	}
	if obj.ErrorCode != 0 {
		return "", errors.New(obj.ErrorInfo)
	}
	// success
	return obj.GroupId, nil
}

func SendGroupMessage(groupID, account, msg string) error {
	//https://console.tim.qq.com/v4/group_open_http_svc/send_group_msg?sdkappid=88888888&identifier=admin&usersig=xxx&random=99999999&contenttype=json
	url := "https://console.tim.qq.com/v4/group_open_http_svc/send_group_msg"
	req_body := &SendGroupMessageRequest{
		GroupId:     groupID,
		FromAccount: account,
		Random:      int64(rand.Uint32()),
		MsgBody: []GroupMessage{
			{MsgType: "TIMTextElem", MsgContent: TextMessageContent{Text: msg}},
		},
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
	if err != nil {
		return err
	}
	obj := &SendGroupMessageResponse{}
	err = json.Unmarshal(content, obj)
	if err != nil {
		return err
	}
	if obj.ErrorCode != 0 {
		return errors.New(obj.ErrorInfo)
	}
	return nil
}

func AddGroupMember(groupID string, supervisorID string) error {
	// https://console.tim.qq.com/v4/group_open_http_svc/add_group_member?sdkappid=88888888&identifier=admin&usersig=xxx&random=99999999&contenttype=json
	url := "https://console.tim.qq.com/v4/group_open_http_svc/add_group_member"
	req_body := &AddGroupMemberRequest{
		GroupId: groupID,
		MemberList: []MemberAccount{
			{MemberAccount: supervisorID},
		},
		Silence: 1,
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
	if err != nil {
		return err
	}
	obj := &AddGroupMemberResponse{}
	err = json.Unmarshal(content, obj)
	if err != nil {
		return err
	}
	if obj.ErrorCode != 0 {
		return errors.New(obj.ErrorInfo)
	}
	return nil
}
