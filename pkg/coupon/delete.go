package coupon

import (
	"coupon_api/pkg/storage/db"
	"fmt"
)

func Delete(column string, content string) {
	db, err := db.Open()
	checkErr(err, "Error opening database: ")
	defer db.Close()

	query := fmt.Sprintf("delete from coupons where %v=?", column)

	stmt, err := db.Prepare(query)
	checkErr(err, "Error creating query statement: ")

	_, err = stmt.Exec(content)
	checkErr(err, "Error executing query statement: ")
}

func DeleteByValue(column string, content float64) {
	db, err := db.Open()
	checkErr(err, "Error opening database: ")
	defer db.Close()

	query := fmt.Sprintf("delete from coupons where %v=?", column)

	stmt, err := db.Prepare(query)
	checkErr(err, "Error creating query statement: ")

	_, err = stmt.Exec(content)
	checkErr(err, "Error executing query statement: ")
}
