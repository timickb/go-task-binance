package config

type AppConfig struct {
	AppPort    int    `json:"app_port,omitempty"`
	BinanceURL string `json:"binance_url,omitempty"`
}

func NewDefault() *AppConfig {
	return &AppConfig{
		AppPort:    3001,
		BinanceURL: "https://api.binance.com/api/v3/",
	}
}
