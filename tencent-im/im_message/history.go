package im_message

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"math/rand"
	"net/http"
	"psy-consult-backend/constant"
	tencent_im "psy-consult-backend/tencent-im"
	"psy-consult-backend/tencent-im/usersig"
	"psy-consult-backend/utils/helper"
	"sort"
)

const (
	maxCount = 100
)

func SearchHistoryMessage(fromAccount string, toAccount string, minTime int64, maxTime int64, lastMsgKey string, lastMsgTime int64) (*SearchHistoryMessageResponse, error) {
	// url = "https://console.tim.qq.com/v4/im_open_login_svc/account_import?sdkappid=88888888&identifier=admin&usersig=xxx&random=99999999&contenttype=json"
	url := "https://console.tim.qq.com/v4/openim/admin_getroammsg"
	req_body := &SearchHistoryMessageRequest{
		FromAccount: fromAccount,
		ToAccount:   toAccount,
		MaxCnt:      maxCount,
		MinTime:     minTime,
		MaxTime:     maxTime,
	}
	if lastMsgTime != 0 {
		req_body.MaxTime = lastMsgTime
	}
	if lastMsgKey != "" {
		req_body.LastMsgKey = lastMsgKey
	}
	body, err := json.Marshal(req_body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
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
		return nil, err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	obj := &SearchHistoryMessageResponse{}
	err = json.Unmarshal(content, obj)
	if err != nil {
		return nil, err
	}
	if obj.ErrorCode != 0 {
		return nil, errors.New(obj.ErrorInfo)
	}
	// success
	return obj, nil
}

func SearchAllHistoryMessage(fromAccount string, toAccount string, minTime int64, maxTime int64) ([]MessageInfo, error) {
	res := make([]MessageInfo, 0)
	var lastMsgKey string
	var lastMsgTime int64
	for {
		resp, err := SearchHistoryMessage(fromAccount, toAccount, minTime, maxTime, lastMsgKey, lastMsgTime)
		if err != nil {
			logrus.Warn(constant.REST+"SearchAllHistoryMessage Failed, err= %v", err)
			return nil, err
		}
		res = append(res, resp.MsgList...)
		lastMsgKey = resp.LastMsgKey
		lastMsgTime = resp.LastMsgTime
		if resp.Complete == 1 {
			break
		}
	}
	// 根据timestamp排序
	sort.Sort(MessageInfoSlice(res))
	return res, nil
}

const (
	ReqMsgNumber = 1000 // 全量
)

func GetGroupMessage(groupID string, seq int64) (*GetGroupMessageHistoryResponse, error) {
	url := "https://console.tim.qq.com/v4/group_open_http_svc/group_msg_get_simple"
	req_body := &GetGroupMessageHistoryRequest{
		GroupId:      groupID,
		ReqMsgNumber: ReqMsgNumber,
	}
	if seq != -1 {
		req_body.ReqMsgSeq = seq
	}
	body, err := json.Marshal(req_body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
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
		return nil, err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	obj := &GetGroupMessageHistoryResponse{}
	err = json.Unmarshal(content, obj)
	if err != nil {
		return nil, err
	}
	if obj.ErrorCode != 0 {
		return nil, errors.New(obj.ErrorInfo)
	}
	// success
	return obj, nil
}

func min(a, b int64) int64 {
	if a > b {
		return b
	} else {
		return a
	}
}

func GetAllGroupMessage(groupID string) ([]GroupMessageHistory, error) {
	res := make([]GroupMessageHistory, 0)
	var seq int64 = -1
	for {
		resp, err := GetGroupMessage(groupID, seq)
		if err != nil {
			logrus.Warn(constant.REST+"GetAllGroupMessage Failed, err= %v", err)
			return nil, err
		}
		var mseq int64 = 1000000000
		for _, msg := range resp.RspMsgList {
			mseq = min(mseq, msg.MsgSeq)
		}
		res = append(res, resp.RspMsgList...)
		if resp.IsFinished == 1 {
			break
		}
		seq = mseq
	}
	// 根据timestamp排序
	sort.Sort(GroupMessageHistorySlice(res))
	return res, nil
}
