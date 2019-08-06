package coupon

import (
	"coupon_api/pkg/storage/db"
	"fmt"
)

func Create(name string, brand string, value float64, created string, expiry string) (Coupon, error) {
	cp := Coupon{
		Name:      name,
		Brand:     brand,
		Value:     value,
		CreatedAt: created,
		Expiry:    expiry,
	}

	// Open up our database connection.
	db, err := db.Open()
	checkErr(err, "Error opening database: ")
	defer db.Close()

	query := fmt.Sprintf("INSERT INTO coupons (name, brand, value, created, expiry) VALUES ( '%v', '%v', '%v', '%v', '%v' )", cp.Name, cp.Brand, cp.Value, cp.CreatedAt, cp.Expiry)
	insert, err := db.Query(query)
	checkErr(err, "Error inserting values into database: ")

	defer insert.Close()
	return cp, err
}
