package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type Datastore interface {
	GetCoupons() ([]Coupon, error)
	GetCoupon(int) (Coupon, error)
	AddCoupon(Coupon) (int, error)
	UpdateCoupon(int, Coupon) (bool, error)
}

type ActualDataStore struct {
	db *sql.DB
}

func (fb *FakeDatastore) UpdateCoupon(couponID int, coupon Coupon) (bool, error) {
	return true, nil
}

func (fb *FakeDatastore) AddCoupon(coupon Coupon) (int, error) {
	return 0, nil
}

func (fb *FakeDatastore) GetCoupon(CouponID int) (Coupon, error) {
	return Coupon{}, nil
}

func (fb *FakeDatastore) GetCoupons() ([]Coupon, error) {
	return nil, nil
}

func (db *ActualDataStore) UpdateCoupon(couponID int, coupon Coupon) (bool, error) {
	return true, nil
}

func (db *ActualDataStore) AddCoupon(coupon Coupon) (int, error) {
	return 0, nil
}

func (db *ActualDataStore) GetCoupon(CouponID int) (Coupon, error) {
	return Coupon{}, nil
}

func (db *ActualDataStore) GetCoupons() ([]Coupon, error) {
	return nil, nil
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
	coupons, _ := c.db.GetCoupons()
	return coupons
}

type FakeDatastore struct {
	Coupons []Coupon
}

// func (f *FakeDatastore) GetCoupons() []Coupon {
// 	return f.Coupons
// }

type Handler struct {
	Coupons *CouponRetriever
}

func (h *Handler) HandleCoupons(res http.ResponseWriter, req *http.Request) {
	data := h.Coupons.ListCoupons()
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(data)
}

func main() {
	// create a user in mysql to access database
	// db, _ := sql.Open("mysql", "root:password1@tcp(127.0.0.1:3306)/coupons")
	// actualDS := &ActualDataStore{db}
	// db := &ActualDataStore{}

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

	couponRetriever := &CouponRetriever{fakeDb}
	h := &Handler{couponRetriever}
	http.HandleFunc("/coupon", h.HandleCoupons)
	http.ListenAndServe(":8080", nil)
}
