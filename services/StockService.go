package services

import (
	"context"
	"github.com/NeuronEvolution/Stock/models"
	"github.com/NeuronEvolution/Stock/storages"
	"github.com/NeuronEvolution/Stock/storages/fin-stock"
	"github.com/NeuronEvolution/log"
	"github.com/NeuronEvolution/sql/runtime"
	"go.uber.org/zap"
)

type StockServiceOptions struct {
	StockStorageConnectionString string
}

type StockService struct {
	logger  *zap.Logger
	options *StockServiceOptions
	storage *storages.Storage
}

func NewStockService(options *StockServiceOptions) (s *StockService, err error) {
	s = &StockService{}
	s.logger = log.TypedLogger(s)
	s.options = options
	s.storage, err = storages.NewStorage(&storages.Options{ConnectionString: options.StockStorageConnectionString})
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (s *StockService) List(query *models.StockListQuery) (list []*models.Stock, err error) {
	q := s.storage.Stock.Stock.GetQuery()
	if query.ExchangeId != nil {
		q.ExchangeId_Column(runtime.RELATION_EQUAL, *query.ExchangeId)
	}

	dbStockList, err := q.SelectList(context.Background())
	if err != nil {
		return nil, err
	}

	return fin_stock.FromStockList(dbStockList), nil
}

func (s *StockService) Get(stockId string) (stock *models.Stock, err error) {
	dbStock, err := s.storage.Stock.Stock.SelectByStockId(context.Background(), nil, stockId)
	if err != nil {
		return nil, err
	}

	return fin_stock.FromStock(dbStock), nil
}
