package tencent_wechat

import "testing"

func TestWeChatLogin(t *testing.T) {
	_, err := WeChatLogin("wx218c7ba2bb3da68e", "")
	if err != nil {
		t.Error(err)
	}
	t.Logf("login success")
}
