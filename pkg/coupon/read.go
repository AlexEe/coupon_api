package coupon

import (
	"coupon_api/pkg/storage/db"
	"fmt"
	"os"
)

// GetCouponBy retrieves a single coupon
func Read(column string, content string) []Coupon {
	var couponList []Coupon

	db, err := db.Open()
	if db == nil || err != nil {
		fmt.Println("Error opening database:", err)
		os.Exit(1)
	}
	defer db.Close()

	query := fmt.Sprintf("SELECT * FROM coupons WHERE %v = '%v'", column, content)

	result, err := db.Query(query)
	for result.Next() {
		var coupon Coupon
		err = result.Scan(&coupon.ID, &coupon.Name, &coupon.Brand, &coupon.Value, &coupon.CreatedAt, &coupon.Expiry)
		if err != nil {
			fmt.Println("Error getting coupons from database:", err)
			os.Exit(1)
		}
		couponList = append(couponList, coupon)
	}
	return couponList

}

// GetCouponByValue retrieves a single coupon by value which is a float, not string
func ReadByValue(column string, content float64) []Coupon {
	var couponList []Coupon

	db, err := db.Open()
	if db == nil || err != nil {
		fmt.Println("Error opening database:", err)
		os.Exit(1)
	}
	defer db.Close()

	query := fmt.Sprintf("SELECT * FROM coupons WHERE %v = '%v'", column, content)

	result, err := db.Query(query)
	for result.Next() {
		var coupon Coupon
		err = result.Scan(&coupon.ID, &coupon.Name, &coupon.Brand, &coupon.Value, &coupon.CreatedAt, &coupon.Expiry)
		if err != nil {
			fmt.Println("Error getting coupons from database:", err)
			os.Exit(1)
		}
		couponList = append(couponList, coupon)
	}
	return couponList
}
