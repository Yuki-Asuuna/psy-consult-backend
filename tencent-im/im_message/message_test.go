package im_message

import (
	"fmt"
	"psy-consult-backend/utils/helper"
	"testing"
)

func TestSearchAllHistoryMessage(t *testing.T) {
	res, err := SearchAllHistoryMessage("test1", "test2", 1648191651, 1648192108)
	if err != nil {
		t.Error(err)
		return
	}
	for _, m := range res {
		fmt.Println(m)
	}
	t.Log(res)
	t.Log("Success")
}

func TestSendTextMessage(t *testing.T) {
	err := SendTextMessage("37f525e2b6fc3cb4abd882f708ab80eb", "ace23c321i823", "是因为遇到什么不顺心的事情了吗")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("success")
}

func TestCreateNewGroup(t *testing.T) {
	groupID, err := CreateNewGroup("oLbsZ63ONd3jndP9hcfsoycBzqOk", "fcdbd14ffa933c5622e48828e824c517", "712836823612312")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("success")
	t.Logf(groupID)
}

func TestSendGroupMessage(t *testing.T) {
	for i := 0; i < 20; i++ {
		err := SendGroupMessage("@TGS#2P5MZQBII", "8f1475499c41ae3a7f891979a5d993a2", "知道了，我无能为力"+helper.I2S(i))
		if err != nil {
			t.Error(err)
			return
		}
		t.Logf("success")
	}
}

func TestAddGroupMember(t *testing.T) {
	err := AddGroupMember("@TGS#2P5MZQBII", "8f1475499c41ae3a7f891979a5d993a2")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("success")
}

func TestGetAllGroupMessage(t *testing.T) {
	resp, err := GetAllGroupMessage("@TGS#2P5MZQBII")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(resp)
}
