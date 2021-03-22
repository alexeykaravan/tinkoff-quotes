package tinkoff

import "time"

//Quote ...
type Quote struct {
	Symbol string
	Bid    float64
	Ask    float64
	Time   time.Time
}
