package main

import (
	"coupon_api/pkg/http/handler"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/readall", handler.ReadAll)
	http.HandleFunc("/read", handler.Read)
	http.HandleFunc("/add", handler.Add)

	http.ListenAndServe(":8080", nil)
}
