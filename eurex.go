package eurex

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"time"

	cache "github.com/patrickmn/go-cache"
)

const (
	serviceURL = "https://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist-90d.xml"
)

// Client holds a connection to the service
type Client struct {
	// client is the HTTP client the package will use for requests.
	client *http.Client

	// serviceURL is the url to fetch the info from the service.
	serviceURL string

	// expiry is the time of duration of the cache
	expiry time.Duration

	// Services used for communicating with the external service.
	Exchange *ExchangeService
	Cache    *CacheService
}

// NewClient will create http client to create http request
func NewClient() *Client {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
		Transport: &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: true,
		},
	}

	expiry := 10 * time.Minute

	c := &Client{
		client:     httpClient,
		serviceURL: serviceURL,
		expiry:     expiry,
	}

	// Init a new store.
	store := cache.New(expiry, 10*time.Minute)

	// Init services
	c.Exchange = NewExchangeService(c)
	c.Cache = NewCacheService(c, store)

	return c
}

// NewRequest create a new request to the external service
func (c *Client) NewRequest() (*http.Request, error) {

	// Create the request.
	req, err := http.NewRequest("GET", serviceURL, nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// Do sends a service request and returns the response
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	response, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return response, err
}

// ReadXML is a function to read the XML of the external service
func (c *Client) ReadXML() (*XMLFile, error) {
	// Build request
	request, err := c.NewRequest()
	if err != nil {
		return nil, err
	}

	// Make request
	var response *http.Response
	if response, err = c.Do(request); err != nil {
		return nil, err
	}

	// Read the response
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// Parse XML
	res := &XMLFile{}
	err = xml.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// XMLFile represents the root of the document
type XMLFile struct {
	Body Body `xml:"Cube"`
}

// Body represents the body of the document
type Body struct {
	Days []Day `xml:"Cube"`
}

// Day represents one day and an array of rates
type Day struct {
	Time  string `xml:"time,attr"`
	Rates []Rate `xml:"Cube"`
}

// Rate represents the rates per day by currency
type Rate struct {
	Currency string  `xml:"currency,attr"`
	Rate     float64 `xml:"rate,attr"`
}
