package services

import (
	"github.com/NeuronEvolution/Stock/models"
	"github.com/NeuronEvolution/Stock/storage/stock"
	"github.com/NeuronEvolution/log"
	"go.uber.org/zap"
)

type StockServiceOptions struct {
	StockStorageConnectionString string
}

type StockService struct {
	logger  *zap.Logger
	options *StockServiceOptions
	storage *stock.Storage
}

func NewStockService(options *StockServiceOptions) (s *StockService, err error) {
	s = &StockService{}
	s.logger = log.TypedLogger(s)
	s.options = options
	s.storage, err = stock.NewStorage(&stock.Options{ConnectionString: options.StockStorageConnectionString})
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (s *StockService) List(query *models.StockListQuery) (list []*models.Stock, err error) {
	return s.storage.SelectList(query)
}

func (s *StockService) Get(stockId string) (stock *models.Stock, err error) {
	return s.storage.SelectByStockId(stockId)
}
