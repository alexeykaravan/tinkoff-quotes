package tinkoff

import (
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type socket struct {
	conn    *websocket.Conn
	done    chan bool
	symbols sync.Map
}

type wsEvent struct {
	Event string `json:"event"`
	Figi  string `json:"figi"`
	Depth int    `json:"depth"`
}

type wsMessage struct {
	Payload OrderBook `json:"payload"`
	Event   string    `json:"event"`
	Time    time.Time `json:"time"`
}

func (s *socket) reader(ch chan<- Quote) {
	defer close(s.done)

	for {
		_, msg, err := s.conn.ReadMessage()
		if err != nil {
			log.Printf("can't read message %s\n", err.Error())

			return
		}

		event := &wsMessage{}
		if err := json.Unmarshal(msg, event); err != nil {
			continue
		}

		if event.Event != "orderbook" {
			continue
		}

		if len(event.Payload.Bids) == 0 || len(event.Payload.Asks) == 0 {
			continue
		}

		v, ok := s.symbols.Load(event.Payload.FIGI)
		if ok {
			select {
			case ch <- Quote{
				Symbol: v.(string),
				Bid:    event.Payload.Bids[0][0],
				Ask:    event.Payload.Asks[0][0],
				Time:   event.Time}:
			default:
				log.Printf("chanal is full: %v\n", len(ch))
			}
		}
	}
}

func (s *socket) subscribe(figi string) {
	byte, _ := json.Marshal(wsEvent{
		Event: "orderbook:subscribe",
		Figi:  figi,
		Depth: 1,
	})

	if err := s.conn.WriteMessage(websocket.TextMessage, byte); err != nil {
		log.Printf("can't subscribe to %s %s\n", figi, err.Error())
	}
}
