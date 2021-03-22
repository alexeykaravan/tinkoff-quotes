package tinkoff

import (
	"errors"
	"net/http"

	"github.com/gorilla/websocket"
)

const (
	api = "https://api-invest.tinkoff.ru/openapi"

	wsapi = "wss://api-invest.tinkoff.ru/openapi/md/v1/md-openapi/ws"
)

//Tinkoff ...
type Tinkoff struct {
	userKey string
	client  *Client
}

//New ...
func New(key string) *Tinkoff {
	return &Tinkoff{
		userKey: key,
		client:  &Client{&http.Client{}},
	}
}

//GetStocks ...
func (t *Tinkoff) GetStocks() ([]Instrument, error) {
	event := &Message{}

	err := t.client.do(api+"/market/stocks", t.userKey, event)
	if err != nil {
		return nil, err
	}

	return event.Payload.Instruments, nil
}

//GetBonds ...
func (t *Tinkoff) GetBonds() ([]Instrument, error) {
	event := &Message{}

	err := t.client.do(api+"/market/bonds", t.userKey, event)
	if err != nil {
		return nil, err
	}

	return event.Payload.Instruments, nil
}

//GetEtfs ...
func (t *Tinkoff) GetEtfs() ([]Instrument, error) {
	event := &Message{}

	err := t.client.do(api+"/market/etfs", t.userKey, event)
	if err != nil {
		return nil, err
	}

	return event.Payload.Instruments, nil
}

//GetCurrencies ...
func (t *Tinkoff) GetCurrencies() ([]Instrument, error) {
	event := &Message{}

	err := t.client.do(api+"/market/currencies", t.userKey, event)
	if err != nil {
		return nil, err
	}

	return event.Payload.Instruments, nil
}

//SubscribeForQuotes ...
func (t *Tinkoff) SubscribeForQuotes(instruments []Instrument, ch chan<- Quote) error {
	client, _, err := websocket.DefaultDialer.Dial(wsapi, http.Header{"Authorization": {"Bearer " + t.userKey}})
	if err != nil {
		return err
	}

	defer client.Close()

	socket := &socket{
		conn: client,
		done: make(chan bool),
	}

	go func() {
		socket.reader(ch)
	}()

	for _, s := range instruments {
		socket.symbols.Store(s.Figi, s.Ticker)

		socket.subscribe(s.Figi)
	}

	select {
	case <-socket.done:
		return errors.New("client.DisconnectedChannel")
	}
}
