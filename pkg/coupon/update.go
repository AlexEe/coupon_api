package coupon

import (
	"coupon_api/pkg/storage/db"
)

// Update value column for coupon identified by id
func UpdateValue(id string, value float64) []Coupon {
	db, err := db.Open()
	checkErr(err)

	stmt, err := db.Prepare("update coupons set value=? where id=?")
	checkErr(err)

	_, err = stmt.Exec(value, id)
	checkErr(err)

	coupons := Read("id", id)

	return coupons
}
