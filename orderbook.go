package tinkoff

type OrderBook struct {
	FIGI  string          `json:"figi"`
	Depth int             `json:"depth"`
	Bids  []PriceQuantity `json:"bids"`
	Asks  []PriceQuantity `json:"asks"`
}

type PriceQuantity [2]float64
