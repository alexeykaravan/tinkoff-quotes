package tinkoff

//Message ...
type Message struct {
	TrackingID string  `json:"trackingId"`
	Payload    Payload `json:"payload"`
	Status     string  `json:"status"`
}

//Instruments ...
type Instrument struct {
	Figi              string  `json:"figi"`
	Ticker            string  `json:"ticker"`
	MinPriceIncrement float64 `json:"minPriceIncrement"`
	Lot               int     `json:"lot"`
	Currency          string  `json:"currency"`
	Name              string  `json:"name"`
	Type              string  `json:"type"`
}

//Payload ...
type Payload struct {
	Instruments []Instrument `json:"instruments"`
	Total       int          `json:"total"`
}

type OrderBook struct {
	FIGI  string          `json:"figi"`
	Depth int             `json:"depth"`
	Bids  []PriceQuantity `json:"bids"`
	Asks  []PriceQuantity `json:"asks"`
}

type PriceQuantity [2]float64
