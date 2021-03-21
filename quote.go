package tinkoff

import "time"

//Quote ...
type Quote struct {
	symbol string
	bid    float64
	ask    float64
	time   time.Time
}
