package coupon

import (
	"coupon_api/pkg/storage/db"
	"fmt"
	"os"
)

func GetSingleCoupon() Coupon {
	var cp Coupon

	db, err := db.Open()
	if db == nil || err != nil {
		fmt.Println("Error opening database:", err)
		os.Exit(1)
	}
	defer db.Close()

	err = db.QueryRow("SELECT * FROM coupons WHERE id = 1").Scan(&cp.ID, &cp.Name, &cp.Brand, &cp.Value, &cp.CreatedAt, &cp.Expiry)
	if err != nil {
		fmt.Println("Error getting coupon from database:", err)
	}
	return cp
}
