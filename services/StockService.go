package services

import (
	"context"
	"fmt"
	"github.com/NeuronEvolution/Stock/models"
	"github.com/NeuronEvolution/Stock/storages/fin-stock"
	"github.com/NeuronFramework/log"
	"go.uber.org/zap"
)

type StockServiceOptions struct {
	StockStorageConnectionString string
}

type StockService struct {
	logger  *zap.Logger
	options *StockServiceOptions
	db      *fin_stock.DB
}

func NewStockService(options *StockServiceOptions) (s *StockService, err error) {
	s = &StockService{}
	s.logger = log.TypedLogger(s)
	s.options = options
	s.db, err = fin_stock.NewDB(options.StockStorageConnectionString)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (s *StockService) List(query *models.StockListQuery) (list []*models.Stock, err error) {
	fmt.Println("List", query)

	q := s.db.Stock.GetQuery()
	if query.ExchangeId != nil {
		q.ExchangeId_Equal(*query.ExchangeId)
	}

	dbStockList, err := q.SelectList(context.Background())
	if err != nil {
		return nil, err
	}

	return fin_stock.FromStockList(dbStockList), nil
}

func (s *StockService) Get(stockId string) (stock *models.Stock, err error) {
	dbStock, err := s.db.Stock.GetQuery().StockId_Equal(stockId).Select(context.Background())
	if err != nil {
		return nil, err
	}

	return fin_stock.FromStock(dbStock), nil
}
