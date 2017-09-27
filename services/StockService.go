package services

import (
	"github.com/NeuronEvolution/Stock/models"
	"github.com/NeuronEvolution/Stock/storage/stock"
	"github.com/NeuronEvolution/pkg"
	"go.uber.org/zap"
)

type StockServiceOptions struct {
	StockStorageConnectionString string
}

type StockService struct {
	logger  *zap.Logger
	options *StockServiceOptions
	storage *stock.StockStorage
}

func NewStockService(options *StockServiceOptions) (s *StockService, err error) {
	s = &StockService{}
	s.logger = pkg.TypedLogger(s)
	s.options = options
	s.storage, err = stock.NewStockStorage(&stock.StockStorageOptions{ConnectionString: options.StockStorageConnectionString})
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
