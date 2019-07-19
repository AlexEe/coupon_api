package main

import (
	"coupon_api/pkg/http/handler"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/coupons", handler.GetAllCoupons)
	http.HandleFunc("/coupon", handler.GetSingleCoupon)
	http.HandleFunc("/add", handler.AddCoupon)

	http.ListenAndServe(":8080", nil)
}
