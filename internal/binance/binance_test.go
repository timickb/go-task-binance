package binance

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestGetPrice(t *testing.T) {
	baseUrl := "https://api.binance.com/api/v3/"
	s := New(logrus.New(), baseUrl)

	price, err := s.GetPrice("ETH-USDT")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(price)
}

func TestGetPriceInvalidSymbol(t *testing.T) {
	baseUrl := "https://api.binance.com/api/v3/"
	s := New(logrus.New(), baseUrl)

	_, err := s.GetPrice("invalid_symbol")
	if err == nil {
		t.Fatal("expected err")
	}
}
