package coupon

import (
	"coupon_api/pkg/storage/db"
	"fmt"
	"os"
)

// GetSingleCoupon retrieves a single coupon by name
func GetSingleCoupon(name string) Coupon {
	var cp Coupon

	db, err := db.Open()
	if db == nil || err != nil {
		fmt.Println("Error opening database:", err)
		os.Exit(1)
	}
	defer db.Close()

	query := fmt.Sprintf("SELECT * FROM coupons WHERE name = '%v'", name)

	err = db.QueryRow(query).Scan(&cp.ID, &cp.Name, &cp.Brand, &cp.Value, &cp.CreatedAt, &cp.Expiry)
	if err != nil {
		fmt.Println("Error getting coupon from database:", err)
		os.Exit(1)
	}
	return cp
}
