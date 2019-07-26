package handler

import (
	"coupon_api/pkg/coupon"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func GetAllCoupons(res http.ResponseWriter, req *http.Request) {
	data := coupon.GetAllCoupons()
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(data)
}

func GetSingleCoupon(res http.ResponseWriter, req *http.Request) {
	name, ok := req.URL.Query()["name"]
	if !ok {
		fmt.Println("Error: Missing parameter 'name'.")
	}
	data := coupon.GetSingleCoupon(name[0])
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(data)
}

func AddCoupon(res http.ResponseWriter, req *http.Request) {
	name, ok := req.URL.Query()["name"]
	if !ok {
		fmt.Println("Error: Missing parameter 'name'.")
	}
	// io.WriteString(res, name[0])

	branch, ok := req.URL.Query()["branch"]
	if !ok {
		fmt.Println("Error: Missing parameter 'branch'.")
	}

	valueString, ok := req.URL.Query()["value"]
	if !ok {
		fmt.Println("Error: Missing parameter 'branch'.")
	}
	value, _ := strconv.ParseFloat(valueString[0], 32)

	created, ok := req.URL.Query()["created"]
	if !ok {
		fmt.Println("Error: Missing parameter 'branch'.")
	}

	expiry, ok := req.URL.Query()["expiry"]
	if !ok {
		fmt.Println("Error: Missing parameter 'branch'.")
	}

	data := coupon.Add(name[0], branch[0], value, created[0], expiry[0])
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(data)
}
