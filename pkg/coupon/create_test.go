package coupon

import "testing"

func TestCreate(t *testing.T) {
	type testCoupon struct {
		name    string
		brand   string
		value   float64
		created string
		expiry  string
	}

	test := testCoupon{"free food", "Mac", 20.2, "2019-10-13", "2019-10-14"}
	coupon := Create(test.name, test.brand, test.value, test.created, test.expiry)
	if coupon.Name != test.name {
		t.Error("Expected", test.name, "Got", coupon.Name)
	}
}
