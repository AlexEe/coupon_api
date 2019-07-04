package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type Datastore interface {
	GetCoupons() []Coupon
	// GetCoupon(int) Coupon
	// AddCoupon(Coupon) int
	// UpdateCoupon(int, Coupon) (bool, error)
}

// no filepath but pointer to databasedriver
type ActualDataStore struct {
	db *sql.DB
}

// func (db *ActualDataStore) UpdateCoupon(couponID int, coupon Coupon) (bool, error) {
// 	// does a bunch of MySQL shit
// 	return true, nil
// }

// func (db *ActualDataStore) AddCoupon(coupon Coupon) (int, error) {
// 	// does a bunch of MySQL shit
// 	return 0, nil
// }

// func (db *ActualDataStore) GetCoupon(CouponID int) (Coupon, error) {
// 	// does a bunch of MySQL shit
// 	return Coupon{}, nil
// }

func (db *ActualDataStore) GetCoupons() []Coupon {
	// does a bunch of MySQL shit
	return nil
}

type Coupon struct {
	Name      string `json:"name"`
	Brand     string `json:"brand"`
	Value     int    `json:"value"`
	CreatedAt string `json:"createdAt"`
	Expiry    string `json:"expiry"`
	ID        string `json:"id"`
}

type CouponRetriever struct {
	db Datastore
}

func (c *CouponRetriever) ListCoupons() []Coupon {
	return c.db.GetCoupons()
}

type FakeDatastore struct {
	Coupons []Coupon
}

func (f *FakeDatastore) GetCoupons() []Coupon {
	return f.Coupons
}

type Handler struct {
	Coupons *CouponRetriever
}

func (h *Handler) ListCoupons(res http.ResponseWriter, req *http.Request) {
	data := h.Coupons.ListCoupons()
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(data)
}

func main() {
	// create a user in mysql to access database
	// db, _ := sql.Open("mysql", "root:password1@tcp(127.0.0.1:3306)/coupons")
	// actualDS := &ActualDataStore{db}

	couponsList := []Coupon{
		Coupon{
			Name:      "example Coupon",
			Brand:     "example Brand",
			Value:     20,
			CreatedAt: "Today",
			Expiry:    "Tomorrow",
		},
	}

	fakeDb := &FakeDatastore{couponsList}
	// db := &ActualDataStore{}
	Coupons := &CouponRetriever{fakeDb}
	h := &Handler{Coupons}
	http.HandleFunc("/coupon", h.ListCoupons)
	http.ListenAndServe(":8080", nil)
}
