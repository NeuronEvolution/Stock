package stock

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"github.com/NeuronEvolution/Stock/storage/stock/gen"
	"github.com/NeuronEvolution/sql/runtime"
	"context"
)

type Database struct {
	db          *runtime.DB
	StockDao    *gen.StockDao
	ExchangeDao *gen.ExchangeDao
}

func NewDatabase(connectionString string) (d *Database, err error) {
	if connectionString == "" {
		return nil, fmt.Errorf("connectionString nil")
	}

	d = &Database{}

	d.db, err = runtime.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	err = d.db.Ping(context.Background())
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

func (d *Database) BeginTransaction(ctx context.Context,level sql.IsolationLevel,readOnly bool) (*runtime.Tx, error) {
	return d.db.BeginTx(ctx,&sql.TxOptions{Isolation:level,ReadOnly:readOnly})
}
