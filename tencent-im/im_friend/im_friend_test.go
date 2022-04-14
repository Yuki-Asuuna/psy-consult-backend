package im_friend

import "testing"

func TestAddFriend(t *testing.T) {
	err := AddFriend("bcdff683f9eee098e376d53a19e2fdd3", "048e2cbc4dbbe5e3493b4495c3facbed")
	if err != nil {
		t.Error(err)
	}
	t.Logf("success")
}

func TestDeleteFriend(t *testing.T) {
	err := DeleteFriend("bcdff683f9eee098e376d53a19e2fdd3", "048e2cbc4dbbe5e3493b4495c3facbed")
	if err != nil {
		t.Error(err)
	}
	t.Logf("success")
}
