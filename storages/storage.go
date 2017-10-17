package storages

import (
	"github.com/NeuronEvolution/Stock/storages/fin-stock"
	"github.com/NeuronEvolution/log"
	"go.uber.org/zap"
)

type Options struct {
	ConnectionString string
}

type Storage struct {
	logger  *zap.Logger
	options *Options
	Stock   *fin_stock.DB
}

func NewStorage(options *Options) (s *Storage, err error) {
	s = &Storage{}
	s.logger = log.TypedLogger(s)
	s.options = options

	s.Stock, err = fin_stock.NewDB(s.options.ConnectionString)
	if err != nil {
		return nil, err
	}

	return s, nil
}
