package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
	"github.com/timickb/task-17apr/internal/config"
	"net/http"
	"strings"
)

type Binance interface {
	GetPrice(symbol string) (float64, error)
	GetPrices(symbols []string) (map[string]float64, error)
}

type Server struct {
	binance Binance
	logger  *logrus.Logger
	router  *gin.Engine
	cfg     *config.AppConfig
}

func New(log *logrus.Logger, cfg *config.AppConfig, binance Binance) *Server {
	srv := &Server{
		logger:  log,
		cfg:     cfg,
		binance: binance,
	}

	srv.router = gin.New()

	api := srv.router.Group("api/v1")
	api.GET("rates", srv.getRates)
	api.POST("rates", srv.postRates)

	return srv
}

func (s *Server) Run() error {
	if err := s.router.Run(fmt.Sprintf(":%d", s.cfg.AppPort)); err != nil {
		return err
	}
	return nil
}

func (s *Server) postRates(ctx *gin.Context) {
	req := &PairsRequest{}

	if err := ctx.ShouldBindBodyWith(req, binding.JSON); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, ErrResponse{
			Code: http.StatusBadRequest,
			Msg:  "invalid body",
		})
		return
	}

	prices, err := s.binance.GetPrices(req.Pairs)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, ErrResponse{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, prices)
}

func (s *Server) getRates(ctx *gin.Context) {
	pairs := ctx.Request.URL.Query().Get("pairs")
	if pairs == "" {
		ctx.IndentedJSON(http.StatusBadRequest, ErrResponse{
			Code: http.StatusBadRequest,
			Msg:  "no pairs specified",
		})
		return
	}
	symbols := strings.Split(pairs, ",")

	prices, err := s.binance.GetPrices(symbols)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, ErrResponse{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, prices)
}
