package binance

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type tickerPriceResponse struct {
	Symbol string `json:"symbol,omitempty"`
	Price  string `json:"price,omitempty"`
}

type errResponse struct {
	Code int    `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
}

type Service struct {
	logger  *logrus.Logger
	apiHost string
}

func New(logger *logrus.Logger, baseURL string) *Service {
	return &Service{
		logger:  logger,
		apiHost: baseURL,
	}
}

func (s *Service) GetPrices(symbols []string) (map[string]float64, error) {
	result := make(map[string]float64)

	for _, symbol := range symbols {
		price, err := s.getPrice(symbol)
		if err != nil {
			return nil, fmt.Errorf("err get price for %s: %w", symbol, err)
		}
		result[symbol] = price
	}

	return result, nil
}

func (s *Service) GetPrice(symbol string) (float64, error) {
	return s.getPrice(symbol)
}

func (s *Service) getPrice(symbol string) (float64, error) {
	s.logger.Info("Requested price for ", symbol)

	path, err := url.JoinPath(s.apiHost, "ticker/price")
	if err != nil {
		return 0, fmt.Errorf("err get price: %w", err)
	}

	resp, err := http.Get(path + "?symbol=" + strings.ReplaceAll(symbol, "-", ""))
	if err != nil {
		return 0, fmt.Errorf("err get price: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("err get price: received status %d", resp.StatusCode)
	}

	var response tickerPriceResponse

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("err get price: %w", err)
	}

	respErr, ok := s.parseError(b)
	if ok {
		return 0, fmt.Errorf("err get price: %s", respErr.Msg)
	}

	if err := json.Unmarshal(b, &response); err != nil {
		return 0, fmt.Errorf("err get price: %w", err)
	}

	price, err := strconv.ParseFloat(response.Price, 64)
	if err != nil {
		return 0, fmt.Errorf("err get price: %w", err)
	}

	return price, nil
}

// parseError returned flag indicates whether the error happened
func (s *Service) parseError(data []byte) (errResponse, bool) {
	var errResp errResponse

	_ = json.Unmarshal(data, &errResp)
	if errResp.Msg == "" {
		return errResp, false
	}

	return errResp, true
}
