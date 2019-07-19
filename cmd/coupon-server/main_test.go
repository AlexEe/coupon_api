package main

// func TestNoCoupons(t *testing.T) {
// 	fakeDb := &FakeDatastore{[]Coupon{}}
// 	couponRetriever := &CouponRetriever{fakeDb}

// 	coupons := couponRetriever.ListCoupons()
// 	if len(coupons) != 0 {
// 		t.Error("Expected no coupons")
// 	}
// }

// func TestOneCoupon(t *testing.T) {
// 	couponsList := []Coupon{
// 		Coupon{
// 			Name:      "example Coupon",
// 			Brand:     "example Brand",
// 			Value:     20,
// 			CreatedAt: "Today",
// 			Expiry:    "Tomorrow",
// 		},
// 	}

// 	fakeDb := &FakeDatastore{couponsList}
// 	couponRetriever := &CouponRetriever{fakeDb}

// 	coupons := couponRetriever.ListCoupons()
// 	if len(coupons) != 1 {
// 		t.Error("Expected 1 coupons")
// 	}
// }
