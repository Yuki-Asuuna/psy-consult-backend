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

func TestUpdateAccountAvatar(t *testing.T) {
	err := UpdateAccountAvatar("cce0d4561a778e03bacd7c4f1065577c", "http://8.130.13.233/images/2022/03/28/srchttp3A2F2Finews.gtimg.com2Fnewsapp_bt2F02F139647191402F641referhttp3A2F2Finews.gtimg.md.jpg")
	if err != nil {
		t.Error(err)
	}
	t.Logf("succes")
}
