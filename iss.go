package iss

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/shopspring/decimal"
)

// response represents response data received
// upon successful call to the ISS API.
type response struct {
	Timestamp   int64
	Message     string
	ISSPosition struct {
		Lat  string `json:"latitude"`
		Long string `json:"longitude"`
	} `json:"iss_position"`
}

// Position represents geographical coordinates
// of the International Space Station for the given time.
type Position struct {
	Lat  decimal.Decimal
	Long decimal.Decimal
}

// Client is the International Space Station (ISS) client.
//
// More information about the ISS API can be obtained on
// the website: http://open-notify.org/Open-Notify-API/
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

type option func(*Client) error

// WithHTTPClient is a func option that configures
// ISS client to use custom HTTP client.
func WithHTTPClient(hc *http.Client) option {
	return func(c *Client) error {
		if hc == nil {
			return errors.New("nil http client")
		}
		c.HTTPClient = hc
		return nil
	}
}

// WithBaseURL is a func option that configures
// ISS client to use custom ISS URL.
func WithBaseURL(s string) option {
	return func(c *Client) error {
		if s == "" {
			return errors.New("empty URL string")
		}
		c.BaseURL = s
		return nil
	}
}

// New returns a default ISS Client API or error
// if one provided func option returns err.
func New(opts ...option) (*Client, error) {
	c := Client{
		BaseURL: "http://api.open-notify.org/iss-now.json",
		HTTPClient: &http.Client{
			Timeout: time.Second * 10,
		},
	}
	for _, opt := range opts {
		if err := opt(&c); err != nil {
			return nil, err
		}
	}
	return &c, nil
}

// GetPosition returns International Space Station coordinates
// (latitude/longitude) at the time of the request.
func (c Client) GetPosition() (Position, error) {
	req, err := http.NewRequest(http.MethodGet, c.BaseURL, nil)
	if err != nil {
		return Position{}, err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return Position{}, fmt.Errorf("calling ISS API: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return Position{}, fmt.Errorf("unexpected response status code: %v", res.StatusCode)
	}
	return parseResponse(res.Body)
}

func parseResponse(r io.Reader) (Position, error) {
	var res response
	if err := json.NewDecoder(r).Decode(&res); err != nil {
		return Position{}, err
	}
	lat, err := toDecimal(res.ISSPosition.Lat)
	if err != nil {
		return Position{}, err
	}
	long, err := toDecimal(res.ISSPosition.Long)
	if err != nil {
		return Position{}, err
	}
	return Position{Lat: lat, Long: long}, nil
}

func toDecimal(coordinate string) (decimal.Decimal, error) {
	coordinate64, err := strconv.ParseFloat(coordinate, 64)
	if err != nil {
		return decimal.Decimal{}, fmt.Errorf("converting coordinate %s to decimal %w", coordinate, err)
	}
	return decimal.NewFromFloatWithExponent(coordinate64, -4), nil
}

// GetPosition returns current position of the International Space Station
// It returns latitude and longitude coordinates or error if the position
// cannot be checked at the time.
func GetPosition() (decimal.Decimal, decimal.Decimal, error) {
	issClient, err := New()
	if err != nil {
		return decimal.Decimal{}, decimal.Decimal{}, fmt.Errorf("creating ISS client %w", err)
	}
	p, err := issClient.GetPosition()
	if err != nil {
		return decimal.Decimal{}, decimal.Decimal{}, fmt.Errorf("retrieving ISS position, %w", err)
	}
	return p.Lat, p.Long, nil
}

// GetPositionAsStrings returns current position of the ISS.
// The lat and long are represented as strings. It returns
// an error is the position of the ISS cannot be obtained.
func GetPositionAsStrings() (string, string, error) {
	lat, long, err := GetPosition()
	if err != nil {
		return "", "", err
	}
	return lat.String(), long.String(), nil
}
