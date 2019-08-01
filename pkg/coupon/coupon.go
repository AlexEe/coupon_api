package coupon

type Coupon struct {
	Name      string  `json:"name"`
	Brand     string  `json:"brand"`
	Value     float64 `json:"value"`
	CreatedAt string  `json:"created"`
	Expiry    string  `json:"expiry"`
	ID        string  `json:"id"`
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
