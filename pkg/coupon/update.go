package coupon

import (
	"coupon_api/pkg/storage/db"
	"fmt"
)

// Update all other columns for coupon identified by id
func Update(id string, column string, content string) []Coupon {
	db, err := db.Open()
	checkErr(err, "Error opening database: ")

	query := fmt.Sprintf("update coupons set %v=? where id=?", column)

	stmt, err := db.Prepare(query)
	checkErr(err, "Error creating query statement: ")

	_, err = stmt.Exec(content, id)
	checkErr(err, "Error executing query statement: ")

	coupons := Read("id", id)

	return coupons
}

// Update value column for coupon identified by id
func UpdateValue(id string, value float64) []Coupon {
	db, err := db.Open()
	checkErr(err, "Error opening database: ")

	stmt, err := db.Prepare("update coupons set value=? where id=?")
	checkErr(err, "Error creating query statement: ")

	_, err = stmt.Exec(value, id)
	checkErr(err, "Error executing query statement: ")

	coupons := Read("id", id)

	return coupons
}
