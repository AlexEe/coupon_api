package main

import (
	"coupon_api/pkg/http/handler"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/all", handler.GetAllCoupons)
	http.HandleFunc("/get", handler.GetCouponBy)
	http.HandleFunc("/add", handler.AddCoupon)

	http.ListenAndServe(":8080", nil)
}
