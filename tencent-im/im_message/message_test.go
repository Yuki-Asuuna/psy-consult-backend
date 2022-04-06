package im_message

import (
	"fmt"
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
	err := SendGroupMessage("@TGS#2HA3VPBIM", "fcdbd14ffa933c5622e48828e824c517", "你好，请问有什么可以帮您")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("success")
}

func TestAddGroupMember(t *testing.T) {
	err := AddGroupMember("@TGS#2HA3VPBIM", "8f1475499c41ae3a7f891979a5d993a2")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("success")
}
