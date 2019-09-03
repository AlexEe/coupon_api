package coupon

import (
	"database/sql"
	"fmt"
)

func Delete(db *sql.DB, column string, content string) error {
	query := fmt.Sprintf("delete from coupons where %v=?", column)

	stmt, err := db.Prepare(query)
	checkErr(err, "Error creating query statement: ")

	_, err = stmt.Exec(content)
	checkErr(err, "Error executing query statement: ")

	return err
}

func DeleteByValue(db *sql.DB, column string, content float64) error {
	query := fmt.Sprintf("delete from coupons where %v=?", column)

	stmt, err := db.Prepare(query)
	checkErr(err, "Error creating query statement: ")

	_, err = stmt.Exec(content)
	checkErr(err, "Error executing query statement: ")

	return err
}
