package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Server struct {
	host   string
	secure bool
}

func New(host string, secure bool) *Server {
	return &Server{host: host, secure: secure}
}

func (s *Server) GetPairPrice(pair string) (float64, error) {
	u := url.URL{
		Scheme: "http",
		Host:   s.host,
		Path:   "api/v1/rates",
	}
	if s.secure {
		u.Scheme = "https"
	}

	resp, err := http.Get(u.String() + "?pairs=" + pair)
	if err != nil {
		return 0, err
	}

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("err received status %d", resp.StatusCode)
	}

	result := make(map[string]float64)
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	if err := json.Unmarshal(b, &result); err != nil {
		return 0, err
	}

	return result[pair], nil
}
