package stock

import (
	"fmt"
	"github.com/NeuronEvolution/Stock/models"
	"github.com/NeuronEvolution/pkg"
	"go.uber.org/zap"
)

type StockStorageOptions struct {
	ConnectionString string
}

type StockStorage struct {
	logger  *zap.Logger
	options *StockStorageOptions
	db      *Database
}

func NewStockStorage(options *StockStorageOptions) (s *StockStorage, err error) {
	s = &StockStorage{}
	s.logger = pkg.TypedLogger(s)
	s.options = options

	s.db, err = NewDatabase(s.options.ConnectionString)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (s *StockStorage) SelectByStockId(stockId string) (stock *models.Stock, err error) {
	return nil, nil
}

func (s *StockStorage) SelectList(query *models.StockListQuery) (stockList []*models.Stock, err error) {
	sql := ""
	if query.ExchangeId != nil {
		sql = fmt.Sprintf("WHERE exchange_id='%s'", *query.ExchangeId)
	}

	dbStockList, err := s.db.StockDao.SelectList(nil, sql)
	if err != nil {
		return nil, err
	}

	return fromStockList(dbStockList), nil
}
