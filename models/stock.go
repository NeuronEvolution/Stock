package models

import "time"

type Stock struct {
	StockId        string
	ExchangeId     string
	StockCode      string
	StockNameCN    string
	LaunchDate     time.Time
	WebsiteUrl     string
	IndustryName   string
	CityNameCN     string
	ProvinceNameCN string
}

type StockListQuery struct {
	ExchangeId *string
}
