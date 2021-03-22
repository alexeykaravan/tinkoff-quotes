# tinkoff-quotes
tinkoff investments quotes stream

package main

import (
	"log"
	"os"
	"time"

	"github.com/alexeykaravan/tinkoff-quotes"
)

func main() {
	quotesChanal := make(chan tinkoff.Quote, 1024)

	t := tinkoff.New("")

	curr, err := t.GetStocks()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		timer := time.NewTicker(5 * time.Second)

		err := t.SubscribeForQuotes(curr, quotesChanal)
		if err != nil {
			log.Println(err)
		}

		<-timer.C
	}()

	for quote := range quotesChanal {
		log.Println(quote)
	}
}
