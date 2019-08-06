package coupon

import "testing"

func TestCreate(t *testing.T) {
	test := Coupon{
		ID:        "free food",
		Brand:     "Mac",
		Value:     20.2,
		CreatedAt: "2019-10-13",
		Expiry:    "2019-10-14",
	}

	coupon, err := Create(test.Name, test.Brand, test.Value, test.CreatedAt, test.Expiry)
	if coupon.Name != test.Name {
		t.Error("Expected", test.Name, "Got", coupon.Name)
	}
	if coupon.Brand != test.Brand {
		t.Error("Expected", test.Name, "Got", coupon.Name)
	}
	if coupon.Value != test.Value {
		t.Error("Expected", test.Value, "Got", coupon.Value)
	}
	if coupon.CreatedAt != test.CreatedAt {
		t.Error("Expected", test.CreatedAt, "Got", coupon.CreatedAt)
	}
	if coupon.Expiry != test.Expiry {
		t.Error("Expected", test.Expiry, "Got", coupon.Expiry)
	}
	if err != nil {
		t.Error("Expected err == nil Got", err)
	}
}
