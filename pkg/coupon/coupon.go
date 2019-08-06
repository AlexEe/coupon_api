package coupon

import (
	"log"
)

type Coupon struct {
	Name      string  `json:"name"`
	Brand     string  `json:"brand"`
	Value     float64 `json:"value"`
	CreatedAt string  `json:"created"`
	Expiry    string  `json:"expiry"`
	ID        string  `json:"id"`
}

func checkErr(err error, message string) {
	if err != nil {
		log.Fatalln(message, err)
	}
}
