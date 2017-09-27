package stock

import (
	"database/sql"
	"fmt"
	"github.com/NeuronEvolution/Stock/storage/stock/gen"
	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	db          *sql.DB
	StockDao    *gen.StockDao
	ExchangeDao *gen.ExchangeDao
}

func NewDatabase(connectionString string) (d *Database, err error) {
	if connectionString == "" {
		return nil, fmt.Errorf("connectionString nil")
	}

	d = &Database{}

	d.db, err = sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	err = d.db.Ping()
	if err != nil {
		return nil, err
	}

	d.ExchangeDao = gen.NewExchangeDao(d.db)
	err = d.ExchangeDao.Init()
	if err != nil {
		return nil, err
	}

	d.StockDao = gen.NewStockDao(d.db)
	err = d.StockDao.Init()
	if err != nil {
		return nil, err
	}

	return d, nil
}

func (d *Database) BeginTransaction() (*sql.Tx, error) {
	return d.db.Begin()
}
