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
	err := DeleteIMSDKAccount("50809ebcb0d153cce1108159612eb7bf")
	if err != nil {
		t.Error(err)
	}
	t.Logf("succes")
}
