package tinkoff

import "net/http"

const (
	api = "https://api-invest.tinkoff.ru/openapi"
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
		httpClient: &http.Client{}}
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
func (t *Tinkoff) SubscribeForQuotes(symbols []Symbol, ch <-chan Quote) error {
	return nil
}
