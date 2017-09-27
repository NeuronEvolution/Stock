package handler

import (
	api "github.com/NeuronEvolution/Stock/api/stock/http/server/models"
	"github.com/NeuronEvolution/Stock/api/stock/http/server/restapi/operations"
	"github.com/NeuronEvolution/Stock/models"
)

func fromStock(p *models.Stock) (r *api.Stock) {
	if p == nil {
		return nil
	}

	r = &api.Stock{
		StockID:    &(p.StockId),
		ExchangeID: &(p.ExchangeId),
	}

	return r
}

func fromStockList(p []*models.Stock) (r []*api.Stock) {
	if p == nil {
		return nil
	}

	for _, v := range p {
		r = append(r, fromStock(v))
	}

	return r
}

func toStockListQuery(p *operations.StocksListParams) (r *models.StockListQuery) {
	if p == nil {
		return nil
	}

	r = &models.StockListQuery{}
	r.ExchangeId = p.ExchangeID

	return r
}
