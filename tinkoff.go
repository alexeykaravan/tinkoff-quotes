package tinkoff

import (
	"errors"
	"net/http"

	"github.com/gorilla/websocket"
)

const (
	api = "https://api-invest.tinkoff.ru/openapi"

	wsapi = "https://api-invest.tinkoff.ru/openapi/md/v1/md-openapi/ws"
)

//Tinkoff ...
type Tinkoff struct {
	userKey    string
	httpClient *http.Client
}

//New ...
func New(key string) *Tinkoff {
	return &Tinkoff{
		userKey:    key,
		httpClient: &http.Client{},
	}
}

//GetAllSymbols ...
func (t *Tinkoff) GetAllSymbols() ([]Symbol, error) {
	return nil, nil
}

//GetStocks ...
func (t *Tinkoff) GetStocks() ([]Symbol, error) {
	return nil, nil
}

//GetBonds ...
func (t *Tinkoff) GetBonds() ([]Symbol, error) {
	return nil, nil
}

//GetEtfs ...
func (t *Tinkoff) GetEtfs() ([]Symbol, error) {
	return nil, nil
}

//GetCurrencies ...
func (t *Tinkoff) GetCurrencies() ([]Symbol, error) {
	return nil, nil
}

//SubscribeForQuotes ...
func (t *Tinkoff) SubscribeForQuotes(symbols []Symbol, ch chan<- Quote) error {
	client, _, err := websocket.DefaultDialer.Dial(wsapi, nil)
	if err != nil {
		return err
	}

	defer client.Close()

	socket := &socket{
		conn:    client,
		done:    make(chan bool),
		symbols: make(map[string]string),
	}

	go func() {
		socket.reader(ch)
	}()

	for _, s := range symbols {
		socket.symbols[s.Figi] = s.Name

		socket.subscribe(s.Figi)
	}

	select {
	case <-socket.done:
		return errors.New("client.DisconnectedChannel")
	}
}
