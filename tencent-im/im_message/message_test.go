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
