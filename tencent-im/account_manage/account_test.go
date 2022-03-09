package account_manage

import "testing"

func TestAddIMSDKAccount(t *testing.T) {
	err := AddIMSDKAccount("cce0d4561a778e03bacd7c4f1065577c", "ayato", "")
	if err != nil {
		t.Error(err)
	}
	t.Logf("success")
}

func TestDeleteIMSDKAccount(t *testing.T) {
	err := DeleteIMSDKAccount("cce0d4561a778e03bacd7c4f1065577c")
	if err != nil {
		t.Error(err)
	}
	t.Logf("succes")
}
