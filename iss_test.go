package iss_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/qba73/iss"
	"github.com/shopspring/decimal"
)

func TestISSClient_ReturnsISSPositionOnInputWithCorrectData(t *testing.T) {
	t.Parallel()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"timestamp": 1638559834, "message": "success", "iss_position": {"latitude": "29.9314", "longitude": "11.3786"}}`)
	}))
	defer ts.Close()

	issClient, err := iss.New(iss.WithBaseURL(ts.URL))
	if err != nil {
		t.Fatal(err)
	}
	got, err := issClient.GetPosition()
	if err != nil {
		t.Fatal(err)
	}

	want := iss.Position{
		Lat:  decimal.NewFromFloatWithExponent(29.9314, -4),
		Long: decimal.NewFromFloatWithExponent(11.3786, -4),
	}

	if !cmp.Equal(got, want) {
		t.Error(cmp.Diff(got, want))
	}
}

func TestISSClient_ErrorsOnInputWithInvalidEmptyData(t *testing.T) {
	t.Parallel()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{}`)
	}))
	defer ts.Close()

	issClient, err := iss.New(iss.WithBaseURL(ts.URL))
	if err != nil {
		t.Fatal(err)
	}
	_, err = issClient.GetPosition()
	if err == nil {
		t.Fatal("should error on invalid data")
	}
}

func TestISSClient_ErrorsOnInputWithInvalidLatitudeCoordinates(t *testing.T) {
	t.Parallel()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"timestamp": 1638559834, "message": "success", "iss_position": {"latitude": "", "longitude": "11.3786"}}`)
	}))
	defer ts.Close()

	issClient, err := iss.New(iss.WithBaseURL(ts.URL))
	if err != nil {
		t.Fatal()
	}
	_, err = issClient.GetPosition()
	if err == nil {
		t.Fatal("should error on input with invalid data")
	}
}

func TestISSClient_ErrorsOnInvalidLongitudeCoordinates(t *testing.T) {
	t.Parallel()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"timestamp": 1638559834, "message": "success", "iss_position": {"latitude": "29.9314", "longitude": ""}}`)
	}))
	defer ts.Close()

	issClient, err := iss.New(iss.WithBaseURL(ts.URL))
	if err != nil {
		t.Fatal(err)
	}
	_, err = issClient.GetPosition()
	if err == nil {
		t.Fatal("should error on input with invalid data")
	}
}
