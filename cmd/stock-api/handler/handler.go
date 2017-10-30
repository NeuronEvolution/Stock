package handler

import (
	"github.com/NeuronEvolution/Stock/api/gen/restapi/operations"
	"github.com/NeuronEvolution/Stock/services"
	"github.com/NeuronFramework/errors"
	"github.com/NeuronFramework/log"
	"github.com/NeuronFramework/restful"
	"github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"
)

type StockHandlerOptions struct {
	StorageConnectionString string
}

type StockHandler struct {
	logger  *zap.Logger
	options *StockHandlerOptions
	service *services.StockService
}

func NewStockHandler(options *StockHandlerOptions) (h *StockHandler, err error) {
	h = &StockHandler{}
	h.logger = log.TypedLogger(h)
	h.options = options
	h.service, err = services.NewStockService(
		&services.StockServiceOptions{
			StockStorageConnectionString: options.StorageConnectionString,
		})
	if err != nil {
		return nil, err
	}

	return h, nil
}

func (h *StockHandler) List(params operations.StocksListParams) middleware.Responder {
	query := toStockListQuery(&params)
	stockList, err := h.service.List(query)
	if err != nil {
		return restful.Responder(err)
	}

	return operations.NewStocksListOK().WithPayload(fromStockList(stockList))
}

func (h *StockHandler) Get(params operations.StocksGetParams) middleware.Responder {
	stock, err := h.service.Get(params.StockID)
	if err != nil {
		return restful.Responder(err)
	}

	if stock == nil {
		return restful.Responder(errors.NotFound("stock not exist"))
	}

	return operations.NewStocksGetOK().WithPayload(fromStock(stock))
}
