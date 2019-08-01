package handler

import (
	"coupon_api/pkg/coupon"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func Update(res http.ResponseWriter, req *http.Request) {
	id, ok := req.URL.Query()["id"]
	switch ok {
	case true:
		valueString, ok := req.URL.Query()["value"]
		if ok {
			value, _ := strconv.ParseFloat(valueString[0], 32)
			data := coupon.UpdateValue(id[0], value)
			// io.WriteString(res, "Coupon updated")
			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(http.StatusCreated)
			json.NewEncoder(res).Encode(data)
		}
	default:
		fmt.Println("Error: Missing parameter 'id'.")
		os.Exit(1)
	}
}

func ReadAll(res http.ResponseWriter, req *http.Request) {
	data := coupon.ReadAll()
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(data)
}

func Read(res http.ResponseWriter, req *http.Request) {
	valueString, ok := req.URL.Query()["value"]
	if ok {
		value, _ := strconv.ParseFloat(valueString[0], 32)
		data := coupon.ReadByValue("Value", value)
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusCreated)
		json.NewEncoder(res).Encode(data)
	}

	name, ok := req.URL.Query()["name"]
	if ok {
		data := coupon.Read("name", name[0])
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusCreated)
		json.NewEncoder(res).Encode(data)
	}

	brand, ok := req.URL.Query()["brand"]
	if ok {
		data := coupon.Read("brand", brand[0])
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusCreated)
		json.NewEncoder(res).Encode(data)
	}

	created, ok := req.URL.Query()["created"]
	if ok {
		data := coupon.Read("created", created[0])
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusCreated)
		json.NewEncoder(res).Encode(data)
	}

	expiry, ok := req.URL.Query()["expiry"]
	if ok {
		data := coupon.Read("expiry", expiry[0])
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusCreated)
		json.NewEncoder(res).Encode(data)
	}

}

func Create(res http.ResponseWriter, req *http.Request) {
	name, ok := req.URL.Query()["name"]
	if !ok {
		fmt.Println("Error: Missing parameter 'name'.")
		os.Exit(1)
	}
	// io.WriteString(res, name[0])

	brand, ok := req.URL.Query()["brand"]
	if !ok {
		fmt.Println("Error: Missing parameter 'brand'.")
		os.Exit(1)
	}

	valueString, ok := req.URL.Query()["value"]
	if !ok {
		fmt.Println("Error: Missing parameter 'value'.")
		os.Exit(1)
	}
	value, _ := strconv.ParseFloat(valueString[0], 32)

	created, ok := req.URL.Query()["created"]
	if !ok {
		fmt.Println("Error: Missing parameter 'created'.")
		os.Exit(1)
	}

	expiry, ok := req.URL.Query()["expiry"]
	if !ok {
		fmt.Println("Error: Missing parameter 'expiry'.")
		os.Exit(1)
	}

	data := coupon.Create(name[0], brand[0], value, created[0], expiry[0])
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(data)
}
