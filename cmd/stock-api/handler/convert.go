package handler

import (
	api "github.com/NeuronEvolution/Stock/api/models"
	"github.com/NeuronEvolution/Stock/api/restapi/operations"
	"github.com/NeuronEvolution/Stock/models"
	"github.com/go-openapi/strfmt"
)

func fromStock(p *models.Stock) (r *api.Stock) {
	if p == nil {
		return nil
	}

	r = &api.Stock{}
	r.StockID = p.StockId
	r.ExchangeID = p.ExchangeId
	r.StockCode = p.StockCode
	r.StockNameCN = p.StockNameCN
	r.LaunchDate = strfmt.DateTime(p.LaunchDate)
	r.WebsiteURL=p.WebsiteUrl
	r.IndustryName=p.IndustryName
	r.CityNameCN=p.CityNameCN
	r.ProvinceNameCN=p.ProvinceNameCN

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
