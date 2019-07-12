package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Datastore interface {
	GetAllCoupons() []Coupon
	GetSingleCoupon() Coupon
	// AddCoupon(Coupon) int
	// UpdateCoupon(int, Coupon) (bool, error)
}

type ActualDataStore struct {
	db *sql.DB
}

type FakeDatastore struct {
	Coupons []Coupon
}

type Coupon struct {
	Name      string  `json:"name"`
	Brand     string  `json:"brand"`
	Value     []uint8 `json:"value"`
	CreatedAt string  `json:"created"`
	Expiry    string  `json:"expiry"`
	ID        string  `json:"id"`
}

type CouponRetriever struct {
	db Datastore
}

func (ds *ActualDataStore) GetAllCoupons() []Coupon {
	var couponList []Coupon

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/coupons")

	result, err := db.Query("SELECT * FROM coupons")
	for result.Next() {
		var coupon Coupon
		err = result.Scan(&coupon.ID, &coupon.Name, &coupon.Brand, &coupon.Value, &coupon.CreatedAt, &coupon.Expiry)
		if err != nil {
			fmt.Println("Error getting coupons from database:", err)
		}
		couponList = append(couponList, coupon)
	}
	return couponList
}

func (ds *ActualDataStore) GetSingleCoupon() Coupon {
	var coupon Coupon

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/coupons")

	err = db.QueryRow("SELECT * FROM coupons WHERE id = 1").Scan(&coupon.ID, &coupon.Name, &coupon.Brand, &coupon.Value, &coupon.CreatedAt, &coupon.Expiry)
	if err != nil {
		fmt.Println("Error getting coupon from database:", err)
	}
	return coupon
}

func (c *CouponRetriever) GetAllCoupons() []Coupon {
	return c.db.GetAllCoupons()
}

func (c *CouponRetriever) GetSingleCoupon() Coupon {
	return c.db.GetSingleCoupon()
}

func (dbs *ActualDataStore) AddCoupon(coupon Coupon) (int, error) {
	// Open up our database connection.
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/coupons")

	// if there is an error opening the connection, handle it
	if err != nil {
		fmt.Println("Error opening database:", err)
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	insert, err := db.Query("INSERT INTO coupons (name, brand, value, created, expiry) VALUES ( 'free trousers', 'Nike', '35.50', '2019-07-12', '2019-12-31' );")

	// if there is an error inserting, handle it
	if err != nil {
		fmt.Println("Error inserting values into database", err)
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()
	return 0, nil
}

type Handler struct {
	Coupons *CouponRetriever
}

func (h *Handler) GetAllCoupons(res http.ResponseWriter, req *http.Request) {
	data := h.Coupons.GetAllCoupons()
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(data)
}

func (h *Handler) GetSingleCoupon(res http.ResponseWriter, req *http.Request) {
	data := h.Coupons.GetSingleCoupon()
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(data)
}

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/coupons")
	if err != nil {
		fmt.Println("Error opening database", err)
	}
	ds := &ActualDataStore{db}
	Coupons := &CouponRetriever{ds}

	h := &Handler{Coupons}
	http.HandleFunc("/coupons", h.GetAllCoupons)
	http.HandleFunc("/coupon", h.GetSingleCoupon)

	http.ListenAndServe(":8080", nil)
}

// func (db *ActualDataStore) UpdateCoupon(couponID int, coupon Coupon) (bool, error) {
// 	// does a bunch of MySQL shit
// 	return true, nil
// }

// func (db *ActualDataStore) GetCoupon(CouponID int) (Coupon, error) {
// 	// does a bunch of MySQL shit
// 	return Coupon{}, nil
// }
