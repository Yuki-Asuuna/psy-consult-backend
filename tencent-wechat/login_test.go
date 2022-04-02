package tencent_wechat

import "testing"

func TestWeChatLogin(t *testing.T) {
	_, err := WeChatLogin("wx218c7ba2bb3da68e", "0137q3000IIFAN1mpj200IcPgN17q30t")
	if err != nil {
		t.Error(err)
	}
	t.Logf("login success")
}
