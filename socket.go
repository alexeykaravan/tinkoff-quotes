package tinkoff

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

type socket struct {
	conn    *websocket.Conn
	done    chan bool
	symbols map[string]string
}

type event struct {
	Event     string `json:"event"`
	RequestID string `json:"request_id,omitempty"`
	Figi      string `json:"figi"`
	Depth     string `json:"depth"`
}

func (s *socket) reader(ch chan<- Quote) {
	defer close(s.done)

	for {
		_, msg, err := s.conn.ReadMessage()
		if err != nil {
			log.Printf("can't read message %s\n", err.Error())

			return
		}

		log.Println(msg)

	}
}

func (s *socket) subscribe(figi string) {
	byte, _ := json.Marshal(event{
		Event: "orderbook:subscribe",
		Figi:  figi,
		Depth: "1",
	})

	if err := s.conn.WriteMessage(websocket.TextMessage, byte); err != nil {
		log.Printf("can't subscribe to %s %s\n", figi, err.Error())
	}
}
