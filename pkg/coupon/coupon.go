package coupon

type Coupon struct {
	Name      string  `json:"name"`
	Brand     string  `json:"brand"`
	Value     []uint8 `json:"value"`
	CreatedAt string  `json:"created"`
	Expiry    string  `json:"expiry"`
	ID        string  `json:"id"`
}
