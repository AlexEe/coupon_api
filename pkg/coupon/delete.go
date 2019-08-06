package coupon

import (
	"coupon_api/pkg/storage/db"
	"fmt"
)

func Delete(column string, content string) error {
	db, err := db.Open()
	checkErr(err, "Error opening database: ")
	defer db.Close()

	query := fmt.Sprintf("delete from coupons where %v=?", column)

	stmt, err := db.Prepare(query)
	checkErr(err, "Error creating query statement: ")

	_, err = stmt.Exec(content)
	checkErr(err, "Error executing query statement: ")

	return err
}

func DeleteByValue(column string, content float64) error {
	db, err := db.Open()
	checkErr(err, "Error opening database: ")
	defer db.Close()

	query := fmt.Sprintf("delete from coupons where %v=?", column)

	stmt, err := db.Prepare(query)
	checkErr(err, "Error creating query statement: ")

	_, err = stmt.Exec(content)
	checkErr(err, "Error executing query statement: ")

	return err
}
