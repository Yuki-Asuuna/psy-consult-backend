package im_message

import (
	"fmt"
	"testing"
)

func TestSearchAllHistoryMessage(t *testing.T) {
	res, err := SearchAllHistoryMessage("test1", "test2", 1647315054, 1648315677)
	if err != nil {
		t.Error(err)
		return
	}
	for _, m := range res {
		fmt.Println(m)
	}
	t.Log("Success")
}

func TestSendTextMessage(t *testing.T) {
	err := SendTextMessage("test1", "test2", "什么问题？请说")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("success")
}
