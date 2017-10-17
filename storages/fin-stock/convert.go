package fin_stock

import (
	"github.com/NeuronEvolution/Stock/models"
)

func FromStock(p *Stock) (r *models.Stock) {
	if p == nil {
		return nil
	}

	r = &models.Stock{}
	r.StockId = p.StockId
	r.ExchangeId = p.ExchangeId
	r.StockCode = p.StockCode
	r.StockNameCN = p.StockNameCn
	r.LaunchDate = p.LaunchDate
	r.WebsiteUrl = p.WebsiteUrl
	r.IndustryName = p.IndustryName
	r.CityNameCN = p.CityNameCn
	r.ProvinceNameCN = p.ProvinceNameCn

	return r
}

func FromStockList(p []*Stock) (r []*models.Stock) {
	if p == nil {
		return nil
	}

	r = make([]*models.Stock, 0)
	for _, v := range p {
		r = append(r, FromStock(v))
	}

	return r
}
