package tencent_wechat

import "testing"

func TestWeChatLogin(t *testing.T) {
	_, err := WeChatLogin("wxd28f9fee161a9210", "053G7FFa18evUC0jqaHa1orxAU2G7FFk")
	if err != nil {
		t.Error(err)
	}
	t.Logf("login success")
}
