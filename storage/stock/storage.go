package stock

import (
	"context"
	"github.com/NeuronEvolution/Stock/models"
	"github.com/NeuronEvolution/log"
	"github.com/NeuronEvolution/sql/runtime"
	"go.uber.org/zap"
)

type Options struct {
	ConnectionString string
}

type Storage struct {
	logger  *zap.Logger
	options *Options
	db      *Database
}

func NewStorage(options *Options) (s *Storage, err error) {
	s = &Storage{}
	s.logger = log.TypedLogger(s)
	s.options = options

	s.db, err = NewDatabase(s.options.ConnectionString)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Storage) SelectByStockId(stockId string) (stock *models.Stock, err error) {
	dbStock, err := s.db.StockDao.SelectByStockId(context.Background(), nil, stockId)
	if err != nil {
		return nil, err
	}

	return fromStock(dbStock), nil
}

func (s *Storage) SelectList(query *models.StockListQuery) (stockList []*models.Stock, err error) {
	q := s.db.StockDao.GetQuery()
	if query.ExchangeId != nil {
		q.Column_ExchangeId(runtime.RELATION_EQUAL, *query.ExchangeId)
	}

	dbStockList, err := q.SelectList(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return fromStockList(dbStockList), nil
}
