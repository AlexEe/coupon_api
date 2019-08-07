package coupon

import (
	"database/sql"
	"fmt"
)

func Create(db *sql.DB, name string, brand string, value float64, created string, expiry string) (Coupon, error) {
	cp := Coupon{
		Name:      name,
		Brand:     brand,
		Value:     value,
		CreatedAt: created,
		Expiry:    expiry,
	}

	query := fmt.Sprintf("INSERT INTO coupons (name, brand, value, created, expiry) VALUES ( ?, ?, ?, ?, ? )")
	stmt, err := db.Prepare(query)
	if err != nil {
		return Coupon{}, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(cp.Name, cp.Brand, cp.Value, cp.CreatedAt, cp.Expiry)
	if err != nil {
		return Coupon{}, err
	}

	return cp, nil
}
