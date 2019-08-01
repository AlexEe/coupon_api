package coupon

import (
	"coupon_api/pkg/storage/db"
	"fmt"
)

func ReadAll() []Coupon {
	var couponList []Coupon

	db, err := db.Open()
	checkErr(err, "Error opening database: ")

	result, err := db.Query("SELECT * FROM coupons")
	for result.Next() {
		var coupon Coupon
		err = result.Scan(&coupon.ID, &coupon.Name, &coupon.Brand, &coupon.Value, &coupon.CreatedAt, &coupon.Expiry)
		if err != nil {
			fmt.Println("Error getting coupons from database:", err)
		}
		couponList = append(couponList, coupon)
	}
	return couponList
}
