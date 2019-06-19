package eurex

import (
	"errors"
	"fmt"
)

// ExchangeService handles exchanges request/responses.
type ExchangeService struct {
	client *Client
}

// NewExchangeService creates a new handler for this service.
func NewExchangeService(client *Client) *ExchangeService {
	return &ExchangeService{
		client,
	}
}

// ExchangeRate holds our exchanges rates for a given date
type ExchangeRate struct {
	Date  string
	Rates map[string]float64
}

// Get will fetch the exchange rate for a given information
func (s *ExchangeService) Get(quantity float64, currencyFrom string, currencyTo string, time string) (*float64, error) {
	if quantity <= 0 {
		return nil, errors.New("quantity must be more than 0")
	}
	if currencyFrom == "" {
		return nil, errors.New("currencyFrom code must be passed")
	}
	if currencyTo == "" {
		return nil, errors.New("currencyTo code must be passed")
	}
	if time == "" {
		return nil, errors.New("time code must be passed")
	}

	// If we have cached results, use them
	if results, ok := s.client.Cache.Get(time); ok {

		// Get rate value for the first currency code from cache
		rateCurrencyFrom, err := fetchValue(currencyFrom, results)
		if err != nil{
			return nil, err
		}

		// Get rate value for the second currency code from cache
		rateCurrencyTo, err := fetchValue(currencyTo, results)
		if err != nil{
			return nil, err
		}

		// Calculate the exchange
		result := quantity / rateCurrencyFrom * rateCurrencyTo

		return &result, nil
	}

	// No cached results. Fetch them.
	if err := s.fetch(time); err != nil {
		return nil, err
	}

	// Already cached the items. Let's call again the function.
	return s.Get(quantity, currencyFrom, currencyTo, time)
}

// fetch will update the rates for the given currency from the external service.
func (s *ExchangeService) fetch(time string) error {

	// Read the XML
	xml, err := s.client.ReadXML()
	if err != nil {
		return err
	}

	// Create a structure of data to parse the results
	var days map[string]interface{}
	days = make(map[string]interface{})

	for _, value := range xml.Body.Days {
		// Create a map for rates per day
		var rates map[string]float64
		rates = make(map[string]float64)

		for _, ratePerDay := range value.Rates {
			// Map the results using the currency as a key
			// and the rate as a value
			rates[ratePerDay.Currency] = ratePerDay.Rate
		}

		// Create the info for the given date
		var exchangeRate = &ExchangeRate{
			Date:  value.Time,
			Rates: rates,
		}

		// Store our results in cache
		s.client.Cache.Store(exchangeRate)

		// Store the results in memory to check if we found the element
		days[exchangeRate.Date] = exchangeRate.Rates
	}

	// Check if we get the information for the given date
	_, found := days[time]
	if !found {
		return errors.New("element not found")
	}

	return nil
}

// fetchValue will fetch rate value from the given currency code
func fetchValue(currencyCode string, data *map[string]float64) (float64, error) {
	var rate float64
	mapRates := *data

	if currencyCode == "EUR" {
		rate = float64(1)
	} else {
		if val, ok := mapRates[currencyCode]; ok {
			rate = val
		} else {
			return 0, fmt.Errorf("currency code (%q) not found", currencyCode)
		}
	}

	return rate, nil
}
