package coupon

import (
	"coupon_api/pkg/storage/db"
	"fmt"
	"os"
)

// GetSingleCoupon retrieves a single coupon
func GetSingleCoupon(column string, content string) Coupon {
	var cp Coupon

	db, err := db.Open()
	if db == nil || err != nil {
		fmt.Println("Error opening database:", err)
		os.Exit(1)
	}
	defer db.Close()

	query := fmt.Sprintf("SELECT * FROM coupons WHERE %v = '%v'", column, content)

	err = db.QueryRow(query).Scan(&cp.ID, &cp.Name, &cp.Brand, &cp.Value, &cp.CreatedAt, &cp.Expiry)
	if err != nil {
		fmt.Println("Error getting coupon from database:", err)
		os.Exit(1)
	}
	return cp
}

// GetSingleCoupon retrieves a single coupon by value which is a float, not string
func GetSingleCouponByValue(column string, content float64) Coupon {
	var cp Coupon

	db, err := db.Open()
	if db == nil || err != nil {
		fmt.Println("Error opening database:", err)
		os.Exit(1)
	}
	defer db.Close()

	query := fmt.Sprintf("SELECT * FROM coupons WHERE %v = '%v'", column, content)

	err = db.QueryRow(query).Scan(&cp.ID, &cp.Name, &cp.Brand, &cp.Value, &cp.CreatedAt, &cp.Expiry)
	if err != nil {
		fmt.Println("Error getting coupon from database:", err)
		os.Exit(1)
	}
	return cp
}
