package main

import (
	"coupon_api/pkg/http/handler"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/listcoupons", handler.GetAllCoupons)
	http.HandleFunc("/listcoupon", handler.GetSingleCoupon)
	http.HandleFunc("/add", handler.AddCoupon)

	http.ListenAndServe(":8080", nil)
}
