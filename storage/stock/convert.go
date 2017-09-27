package stock

import (
	"github.com/NeuronEvolution/Stock/models"
	"github.com/NeuronEvolution/Stock/storage/stock/gen"
	"time"
)

func fromStock(p *gen.Stock) (r *models.Stock) {
	if p == nil {
		return nil
	}

	r = &models.Stock{}
	r.StockId = p.StockId
	r.ExchangeId = p.ExchangeId
	r.StockCode = p.StockCode
	r.StockNameCN = p.StockNameCn
	launchDate, err := time.Parse("2006-01-02 15:04:05", p.LaunchDate)
	if err != nil {
		panic(err)
	}
	r.LaunchDate = launchDate
	r.WebsiteUrl = p.WebsiteUrl
	r.IndustryName = p.IndustryName
	r.CityNameCN = p.CityNameCn
	r.ProvinceNameCN = p.ProvinceNameCn

	return r
}

func fromStockList(p []*gen.Stock) (r []*models.Stock) {
	if p == nil {
		return nil
	}

	r = make([]*models.Stock, 0)
	for _, v := range p {
		r = append(r, fromStock(v))
	}

	return r
}
