package main

import (
	"coupon_api/pkg/http/handler"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/readall", handler.ReadAll)
	http.HandleFunc("/read", handler.Read)
	http.HandleFunc("/create", handler.Create)
	http.HandleFunc("/update", handler.Update)

	http.ListenAndServe(":8080", nil)
}
