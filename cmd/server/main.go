package main

import (
	"github.com/sirupsen/logrus"
	"github.com/timickb/task-17apr/internal/binance"
	"github.com/timickb/task-17apr/internal/config"
	server "github.com/timickb/task-17apr/internal/delivery/http/v1"
	"os"
	"strconv"
)

func main() {
	logger := logrus.New()

	cfg := config.NewDefault()
	fillConfigFromEnv(cfg)

	svc := binance.New(logger, cfg.BinanceURL)
	srv := server.New(logger, cfg, svc)

	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}

func fillConfigFromEnv(cfg *config.AppConfig) {
	if os.Getenv("BINANCE_URL") != "" {
		cfg.BinanceURL = os.Getenv("BINANCE_URL")
	}
	if os.Getenv("APP_PORT") != "" {
		cfg.AppPort, _ = strconv.Atoi(os.Getenv("APP_PORT"))
	}
}
