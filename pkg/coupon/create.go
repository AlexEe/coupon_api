package coupon

import (
	"coupon_api/pkg/storage/db"
	"fmt"
)

func Create(name string, brand string, value float64, created string, expiry string) Coupon {
	cp := Coupon{
		Name:      name,
		Brand:     brand,
		Value:     value,
		CreatedAt: created,
		Expiry:    expiry,
	}

	// Open up our database connection.
	db, err := db.Open()

	// if there is an error opening the connection, handle it
	if err != nil {
		fmt.Println("Error opening database:", err)
	}

	// defer the close till after the main function has finished
	defer db.Close()

	query := fmt.Sprintf("INSERT INTO coupons (name, brand, value, created, expiry) VALUES ( '%v', '%v', '%v', '%v', '%v' )", cp.Name, cp.Brand, cp.Value, cp.CreatedAt, cp.Expiry)
	insert, err := db.Query(query)

	// if there is an error inserting, handle it
	if err != nil {
		fmt.Println("Error inserting values into database", err)
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()
	return cp
}
