package gen

import (
	"bytes"
	"context"
	"fmt"
	"github.com/NeuronEvolution/log"
	"github.com/NeuronEvolution/sql/runtime"
	"go.uber.org/zap"
)

const EXCHANGE_TABLE_NAME = "exchange"

const EXCHANGE_FIELD_ID = "id"
const EXCHANGE_FIELD_EXCHANGE_ID = "exchange_id"
const EXCHANGE_FIELD_EXCHANGE_NAME_CN = "exchange_name_cn"
const EXCHANGE_FIELD_EXCHANGE_NAME_EN = "exchange_name_en"
const EXCHANGE_FIELD_CREATE_TIME = "create_time"
const EXCHANGE_FIELD_UPDATE_TIME = "update_time"

const EXCHANGE_ALL_FIELDS_STRING = "id,exchange_id,exchange_name_cn,exchange_name_en,create_time,update_time"

var EXCHANGE_ALL_FIELDS = []string{
	"id",
	"exchange_id",
	"exchange_name_cn",
	"exchange_name_en",
	"create_time",
	"update_time",
}

type Exchange struct {
	Id             int64  //size=20
	ExchangeId     string //size=32
	ExchangeNameCn string //size=128
	ExchangeNameEn string //size=128
	CreateTime     string
	UpdateTime     string
}

type ExchangeQuery struct {
	dao *ExchangeDao
	runtime.Query
}

func NewExchangeQuery(dao *ExchangeDao) *ExchangeQuery {
	q := &ExchangeQuery{}
	q.dao = dao
	q.WhereBuffer = bytes.NewBufferString("")
	q.LimitBuffer = bytes.NewBufferString("")
	q.OrderBuffer = bytes.NewBufferString("")

	return q
}

func (q *ExchangeQuery) Select(ctx context.Context, tx *runtime.Tx) (*Exchange, error) {
	return q.dao.Select(ctx, tx, q.BuildQueryString())
}

func (q *ExchangeQuery) SelectList(ctx context.Context, tx *runtime.Tx) (list []*Exchange, err error) {
	return q.dao.SelectList(ctx, tx, q.BuildQueryString())
}

func (q *ExchangeQuery) Column_Id(r runtime.Relation, v int64) *ExchangeQuery {
	q.WhereBuffer.WriteString("id" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

func (q *ExchangeQuery) Column_ExchangeId(r runtime.Relation, v string) *ExchangeQuery {
	q.WhereBuffer.WriteString("exchange_id" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

func (q *ExchangeQuery) Column_ExchangeNameCn(r runtime.Relation, v string) *ExchangeQuery {
	q.WhereBuffer.WriteString("exchange_name_cn" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

func (q *ExchangeQuery) Column_ExchangeNameEn(r runtime.Relation, v string) *ExchangeQuery {
	q.WhereBuffer.WriteString("exchange_name_en" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

func (q *ExchangeQuery) Column_CreateTime(r runtime.Relation, v string) *ExchangeQuery {
	q.WhereBuffer.WriteString("create_time" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

func (q *ExchangeQuery) Column_UpdateTime(r runtime.Relation, v string) *ExchangeQuery {
	q.WhereBuffer.WriteString("update_time" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

type ExchangeDao struct {
	logger                 *zap.Logger
	db                     *runtime.DB
	insertStmt             *runtime.Stmt
	updateStmt             *runtime.Stmt
	deleteStmt             *runtime.Stmt
	selectStmtAll          *runtime.Stmt
	selectStmtById         *runtime.Stmt
	selectStmtByUpdateTime *runtime.Stmt
	selectStmtByExchangeId *runtime.Stmt
}

func NewExchangeDao(db *runtime.DB) (t *ExchangeDao) {
	t = &ExchangeDao{}
	t.logger = log.TypedLogger(t)
	t.db = db

	return t
}

func (dao *ExchangeDao) Init() (err error) {
	err = dao.prepareInsertStmt()
	if err != nil {
		return err
	}

	err = dao.prepareUpdateStmt()
	if err != nil {
		return err
	}

	err = dao.prepareDeleteStmt()
	if err != nil {
		return err
	}

	err = dao.prepareSelectStmtAll()
	if err != nil {
		return err
	}

	err = dao.prepareSelectStmtById()
	if err != nil {
		return err
	}

	err = dao.prepareSelectStmtByUpdateTime()
	if err != nil {
		return err
	}

	err = dao.prepareSelectStmtByExchangeId()
	if err != nil {
		return err
	}

	return nil
}
func (dao *ExchangeDao) prepareInsertStmt() (err error) {
	dao.insertStmt, err = dao.db.Prepare(context.Background(), "INSERT INTO exchange (exchange_id,exchange_name_cn,exchange_name_en,create_time,update_time) VALUES (?,?,?,?,?)")
	return err
}

func (dao *ExchangeDao) prepareUpdateStmt() (err error) {
	dao.updateStmt, err = dao.db.Prepare(context.Background(), "UPDATE exchange SET exchange_id=?,exchange_name_cn=?,exchange_name_en=?,create_time=?,update_time=? WHERE id=?")
	return err
}

func (dao *ExchangeDao) prepareDeleteStmt() (err error) {
	dao.deleteStmt, err = dao.db.Prepare(context.Background(), "DELETE FROM exchange WHERE id=?")
	return err
}

func (dao *ExchangeDao) prepareSelectStmtAll() (err error) {
	dao.selectStmtAll, err = dao.db.Prepare(context.Background(), "SELECT "+EXCHANGE_ALL_FIELDS_STRING+" FROM exchange")
	return err
}

func (dao *ExchangeDao) prepareSelectStmtById() (err error) {
	dao.selectStmtById, err = dao.db.Prepare(context.Background(), "SELECT "+EXCHANGE_ALL_FIELDS_STRING+" FROM exchange WHERE id=?")
	return err
}

func (dao *ExchangeDao) prepareSelectStmtByUpdateTime() (err error) {
	dao.selectStmtByUpdateTime, err = dao.db.Prepare(context.Background(), "SELECT "+EXCHANGE_ALL_FIELDS_STRING+" FROM exchange WHERE update_time=?")
	return err
}

func (dao *ExchangeDao) prepareSelectStmtByExchangeId() (err error) {
	dao.selectStmtByExchangeId, err = dao.db.Prepare(context.Background(), "SELECT "+EXCHANGE_ALL_FIELDS_STRING+" FROM exchange WHERE exchange_id=?")
	return err
}

func (dao *ExchangeDao) Insert(ctx context.Context, tx *runtime.Tx, e *Exchange) (id int64, err error) {
	stmt := dao.insertStmt
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	result, err := stmt.Exec(ctx, e.ExchangeId, e.ExchangeNameCn, e.ExchangeNameEn, e.CreateTime, e.UpdateTime)
	if err != nil {
		return 0, err
	}

	id, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (dao *ExchangeDao) Update(ctx context.Context, tx *runtime.Tx, e *Exchange) (rowsAffected int64, err error) {
	stmt := dao.updateStmt
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	result, err := stmt.Exec(ctx, e.ExchangeId, e.ExchangeNameCn, e.ExchangeNameEn, e.CreateTime, e.UpdateTime, e.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (dao *ExchangeDao) Delete(ctx context.Context, tx *runtime.Tx, id int64) (rowsAffected int64, err error) {
	stmt := dao.deleteStmt
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	result, err := stmt.Exec(ctx, id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (dao *ExchangeDao) ScanRow(row *runtime.Row) (*Exchange, error) {
	e := &Exchange{}
	err := row.Scan(&e.Id, &e.ExchangeId, &e.ExchangeNameCn, &e.ExchangeNameEn, &e.CreateTime, &e.UpdateTime)
	if err != nil {
		if err == runtime.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return e, nil
}

func (dao *ExchangeDao) ScanRows(rows *runtime.Rows) (list []*Exchange, err error) {
	list = make([]*Exchange, 0)
	for rows.Next() {
		e := Exchange{}
		err = rows.Scan(&e.Id, &e.ExchangeId, &e.ExchangeNameCn, &e.ExchangeNameEn, &e.CreateTime, &e.UpdateTime)
		if err != nil {
			return nil, err
		}
		list = append(list, &e)
	}
	if rows.Err() != nil {
		err = rows.Err()
		return nil, err
	}

	return list, nil
}

func (dao *ExchangeDao) SelectAll(ctx context.Context, tx *runtime.Tx) (list []*Exchange, err error) {
	stmt := dao.selectStmtAll
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	rows, err := stmt.Query(ctx)
	if err != nil {
		dao.logger.Error("sqlDriver", zap.Error(err))
		return nil, err
	}

	return dao.ScanRows(rows)
}

func (dao *ExchangeDao) Select(ctx context.Context, tx *runtime.Tx, query string) (*Exchange, error) {
	row := dao.db.QueryRow(ctx, "SELECT "+EXCHANGE_ALL_FIELDS_STRING+" FROM exchange "+query)
	return dao.ScanRow(row)
}

func (dao *ExchangeDao) SelectList(ctx context.Context, tx *runtime.Tx, query string) (list []*Exchange, err error) {
	rows, err := dao.db.Query(ctx, "SELECT "+EXCHANGE_ALL_FIELDS_STRING+" FROM exchange "+query)
	if err != nil {
		dao.logger.Error("sqlDriver", zap.Error(err))
		return nil, err
	}

	return dao.ScanRows(rows)
}

func (dao *ExchangeDao) SelectById(ctx context.Context, tx *runtime.Tx, Id int64) (*Exchange, error) {
	stmt := dao.selectStmtById
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	return dao.ScanRow(stmt.QueryRow(ctx, Id))
}

func (dao *ExchangeDao) SelectByUpdateTime(ctx context.Context, tx *runtime.Tx, UpdateTime string) (*Exchange, error) {
	stmt := dao.selectStmtByUpdateTime
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	return dao.ScanRow(stmt.QueryRow(ctx, UpdateTime))
}

func (dao *ExchangeDao) SelectListByUpdateTime(ctx context.Context, tx *runtime.Tx, UpdateTime string) (list []*Exchange, err error) {
	stmt := dao.selectStmtByUpdateTime
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	rows, err := stmt.Query(ctx, UpdateTime)
	if err != nil {
		dao.logger.Error("sqlDriver", zap.Error(err))
		return nil, err
	}

	return dao.ScanRows(rows)
}

func (dao *ExchangeDao) SelectByExchangeId(ctx context.Context, tx *runtime.Tx, ExchangeId string) (*Exchange, error) {
	stmt := dao.selectStmtByExchangeId
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	return dao.ScanRow(stmt.QueryRow(ctx, ExchangeId))
}

func (dao *ExchangeDao) GetQuery() *ExchangeQuery {
	return NewExchangeQuery(dao)
}

const STOCK_TABLE_NAME = "stock"

const STOCK_FIELD_ID = "id"
const STOCK_FIELD_STOCK_ID = "stock_id"
const STOCK_FIELD_EXCHANGE_ID = "exchange_id"
const STOCK_FIELD_STOCK_CODE = "stock_code"
const STOCK_FIELD_STOCK_NAME_CN = "stock_name_cn"
const STOCK_FIELD_STOCK_NAME_EN = "stock_name_en"
const STOCK_FIELD_LAUNCH_DATE = "launch_date"
const STOCK_FIELD_COMPANY_NAME_CN = "company_name_cn"
const STOCK_FIELD_COMPANY_NAME_EN = "company_name_en"
const STOCK_FIELD_WEBSITE_URL = "website_url"
const STOCK_FIELD_INDUSTRY_NAME = "industry_name"
const STOCK_FIELD_CITY_NAME_CN = "city_name_cn"
const STOCK_FIELD_CITY_NAME_EN = "city_name_en"
const STOCK_FIELD_PROVINCE_NAME_CN = "province_name_cn"
const STOCK_FIELD_PROVINCE_NAME_EN = "province_name_en"
const STOCK_FIELD_CREATE_TIME = "create_time"
const STOCK_FIELD_UPDATE_TIME = "update_time"

const STOCK_ALL_FIELDS_STRING = "id,stock_id,exchange_id,stock_code,stock_name_cn,stock_name_en,launch_date,company_name_cn,company_name_en,website_url,industry_name,city_name_cn,city_name_en,province_name_cn,province_name_en,create_time,update_time"

var STOCK_ALL_FIELDS = []string{
	"id",
	"stock_id",
	"exchange_id",
	"stock_code",
	"stock_name_cn",
	"stock_name_en",
	"launch_date",
	"company_name_cn",
	"company_name_en",
	"website_url",
	"industry_name",
	"city_name_cn",
	"city_name_en",
	"province_name_cn",
	"province_name_en",
	"create_time",
	"update_time",
}

type Stock struct {
	Id             int64  //size=20
	StockId        string //size=32
	ExchangeId     string //size=32
	StockCode      string //size=32
	StockNameCn    string //size=32
	StockNameEn    string //size=32
	LaunchDate     string
	CompanyNameCn  string //size=128
	CompanyNameEn  string //size=128
	WebsiteUrl     string //size=128
	IndustryName   string //size=32
	CityNameCn     string //size=128
	CityNameEn     string //size=128
	ProvinceNameCn string //size=128
	ProvinceNameEn string //size=128
	CreateTime     string
	UpdateTime     string
}

type StockQuery struct {
	dao *StockDao
	runtime.Query
}

func NewStockQuery(dao *StockDao) *StockQuery {
	q := &StockQuery{}
	q.dao = dao
	q.WhereBuffer = bytes.NewBufferString("")
	q.LimitBuffer = bytes.NewBufferString("")
	q.OrderBuffer = bytes.NewBufferString("")

	return q
}

func (q *StockQuery) Select(ctx context.Context, tx *runtime.Tx) (*Stock, error) {
	return q.dao.Select(ctx, tx, q.BuildQueryString())
}

func (q *StockQuery) SelectList(ctx context.Context, tx *runtime.Tx) (list []*Stock, err error) {
	return q.dao.SelectList(ctx, tx, q.BuildQueryString())
}

func (q *StockQuery) Column_Id(r runtime.Relation, v int64) *StockQuery {
	q.WhereBuffer.WriteString("id" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) Column_StockId(r runtime.Relation, v string) *StockQuery {
	q.WhereBuffer.WriteString("stock_id" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) Column_ExchangeId(r runtime.Relation, v string) *StockQuery {
	q.WhereBuffer.WriteString("exchange_id" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) Column_StockCode(r runtime.Relation, v string) *StockQuery {
	q.WhereBuffer.WriteString("stock_code" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) Column_StockNameCn(r runtime.Relation, v string) *StockQuery {
	q.WhereBuffer.WriteString("stock_name_cn" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) Column_StockNameEn(r runtime.Relation, v string) *StockQuery {
	q.WhereBuffer.WriteString("stock_name_en" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) Column_LaunchDate(r runtime.Relation, v string) *StockQuery {
	q.WhereBuffer.WriteString("launch_date" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) Column_CompanyNameCn(r runtime.Relation, v string) *StockQuery {
	q.WhereBuffer.WriteString("company_name_cn" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) Column_CompanyNameEn(r runtime.Relation, v string) *StockQuery {
	q.WhereBuffer.WriteString("company_name_en" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) Column_WebsiteUrl(r runtime.Relation, v string) *StockQuery {
	q.WhereBuffer.WriteString("website_url" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) Column_IndustryName(r runtime.Relation, v string) *StockQuery {
	q.WhereBuffer.WriteString("industry_name" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) Column_CityNameCn(r runtime.Relation, v string) *StockQuery {
	q.WhereBuffer.WriteString("city_name_cn" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) Column_CityNameEn(r runtime.Relation, v string) *StockQuery {
	q.WhereBuffer.WriteString("city_name_en" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) Column_ProvinceNameCn(r runtime.Relation, v string) *StockQuery {
	q.WhereBuffer.WriteString("province_name_cn" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) Column_ProvinceNameEn(r runtime.Relation, v string) *StockQuery {
	q.WhereBuffer.WriteString("province_name_en" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) Column_CreateTime(r runtime.Relation, v string) *StockQuery {
	q.WhereBuffer.WriteString("create_time" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) Column_UpdateTime(r runtime.Relation, v string) *StockQuery {
	q.WhereBuffer.WriteString("update_time" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

type StockDao struct {
	logger                             *zap.Logger
	db                                 *runtime.DB
	insertStmt                         *runtime.Stmt
	updateStmt                         *runtime.Stmt
	deleteStmt                         *runtime.Stmt
	selectStmtAll                      *runtime.Stmt
	selectStmtById                     *runtime.Stmt
	selectStmtByUpdateTime             *runtime.Stmt
	selectStmtByIndustryName           *runtime.Stmt
	selectStmtByStockId                *runtime.Stmt
	selectStmtByExchangeId             *runtime.Stmt
	selectStmtByExchangeIdAndStockCode *runtime.Stmt
}

func NewStockDao(db *runtime.DB) (t *StockDao) {
	t = &StockDao{}
	t.logger = log.TypedLogger(t)
	t.db = db

	return t
}

func (dao *StockDao) Init() (err error) {
	err = dao.prepareInsertStmt()
	if err != nil {
		return err
	}

	err = dao.prepareUpdateStmt()
	if err != nil {
		return err
	}

	err = dao.prepareDeleteStmt()
	if err != nil {
		return err
	}

	err = dao.prepareSelectStmtAll()
	if err != nil {
		return err
	}

	err = dao.prepareSelectStmtById()
	if err != nil {
		return err
	}

	err = dao.prepareSelectStmtByUpdateTime()
	if err != nil {
		return err
	}

	err = dao.prepareSelectStmtByIndustryName()
	if err != nil {
		return err
	}

	err = dao.prepareSelectStmtByStockId()
	if err != nil {
		return err
	}

	err = dao.prepareSelectStmtByExchangeId()
	if err != nil {
		return err
	}

	err = dao.prepareSelectStmtByExchangeIdAndStockCode()
	if err != nil {
		return err
	}

	return nil
}
func (dao *StockDao) prepareInsertStmt() (err error) {
	dao.insertStmt, err = dao.db.Prepare(context.Background(), "INSERT INTO stock (stock_id,exchange_id,stock_code,stock_name_cn,stock_name_en,launch_date,company_name_cn,company_name_en,website_url,industry_name,city_name_cn,city_name_en,province_name_cn,province_name_en,create_time,update_time) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	return err
}

func (dao *StockDao) prepareUpdateStmt() (err error) {
	dao.updateStmt, err = dao.db.Prepare(context.Background(), "UPDATE stock SET stock_id=?,exchange_id=?,stock_code=?,stock_name_cn=?,stock_name_en=?,launch_date=?,company_name_cn=?,company_name_en=?,website_url=?,industry_name=?,city_name_cn=?,city_name_en=?,province_name_cn=?,province_name_en=?,create_time=?,update_time=? WHERE id=?")
	return err
}

func (dao *StockDao) prepareDeleteStmt() (err error) {
	dao.deleteStmt, err = dao.db.Prepare(context.Background(), "DELETE FROM stock WHERE id=?")
	return err
}

func (dao *StockDao) prepareSelectStmtAll() (err error) {
	dao.selectStmtAll, err = dao.db.Prepare(context.Background(), "SELECT "+STOCK_ALL_FIELDS_STRING+" FROM stock")
	return err
}

func (dao *StockDao) prepareSelectStmtById() (err error) {
	dao.selectStmtById, err = dao.db.Prepare(context.Background(), "SELECT "+STOCK_ALL_FIELDS_STRING+" FROM stock WHERE id=?")
	return err
}

func (dao *StockDao) prepareSelectStmtByUpdateTime() (err error) {
	dao.selectStmtByUpdateTime, err = dao.db.Prepare(context.Background(), "SELECT "+STOCK_ALL_FIELDS_STRING+" FROM stock WHERE update_time=?")
	return err
}

func (dao *StockDao) prepareSelectStmtByIndustryName() (err error) {
	dao.selectStmtByIndustryName, err = dao.db.Prepare(context.Background(), "SELECT "+STOCK_ALL_FIELDS_STRING+" FROM stock WHERE industry_name=?")
	return err
}

func (dao *StockDao) prepareSelectStmtByStockId() (err error) {
	dao.selectStmtByStockId, err = dao.db.Prepare(context.Background(), "SELECT "+STOCK_ALL_FIELDS_STRING+" FROM stock WHERE stock_id=?")
	return err
}

func (dao *StockDao) prepareSelectStmtByExchangeId() (err error) {
	dao.selectStmtByExchangeId, err = dao.db.Prepare(context.Background(), "SELECT "+STOCK_ALL_FIELDS_STRING+" FROM stock WHERE exchange_id=?")
	return err
}

func (dao *StockDao) prepareSelectStmtByExchangeIdAndStockCode() (err error) {
	dao.selectStmtByExchangeIdAndStockCode, err = dao.db.Prepare(context.Background(), "SELECT "+STOCK_ALL_FIELDS_STRING+" FROM stock WHERE exchange_id=? AND stock_code=?")
	return err
}

func (dao *StockDao) Insert(ctx context.Context, tx *runtime.Tx, e *Stock) (id int64, err error) {
	stmt := dao.insertStmt
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	result, err := stmt.Exec(ctx, e.StockId, e.ExchangeId, e.StockCode, e.StockNameCn, e.StockNameEn, e.LaunchDate, e.CompanyNameCn, e.CompanyNameEn, e.WebsiteUrl, e.IndustryName, e.CityNameCn, e.CityNameEn, e.ProvinceNameCn, e.ProvinceNameEn, e.CreateTime, e.UpdateTime)
	if err != nil {
		return 0, err
	}

	id, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (dao *StockDao) Update(ctx context.Context, tx *runtime.Tx, e *Stock) (rowsAffected int64, err error) {
	stmt := dao.updateStmt
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	result, err := stmt.Exec(ctx, e.StockId, e.ExchangeId, e.StockCode, e.StockNameCn, e.StockNameEn, e.LaunchDate, e.CompanyNameCn, e.CompanyNameEn, e.WebsiteUrl, e.IndustryName, e.CityNameCn, e.CityNameEn, e.ProvinceNameCn, e.ProvinceNameEn, e.CreateTime, e.UpdateTime, e.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (dao *StockDao) Delete(ctx context.Context, tx *runtime.Tx, id int64) (rowsAffected int64, err error) {
	stmt := dao.deleteStmt
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	result, err := stmt.Exec(ctx, id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (dao *StockDao) ScanRow(row *runtime.Row) (*Stock, error) {
	e := &Stock{}
	err := row.Scan(&e.Id, &e.StockId, &e.ExchangeId, &e.StockCode, &e.StockNameCn, &e.StockNameEn, &e.LaunchDate, &e.CompanyNameCn, &e.CompanyNameEn, &e.WebsiteUrl, &e.IndustryName, &e.CityNameCn, &e.CityNameEn, &e.ProvinceNameCn, &e.ProvinceNameEn, &e.CreateTime, &e.UpdateTime)
	if err != nil {
		if err == runtime.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return e, nil
}

func (dao *StockDao) ScanRows(rows *runtime.Rows) (list []*Stock, err error) {
	list = make([]*Stock, 0)
	for rows.Next() {
		e := Stock{}
		err = rows.Scan(&e.Id, &e.StockId, &e.ExchangeId, &e.StockCode, &e.StockNameCn, &e.StockNameEn, &e.LaunchDate, &e.CompanyNameCn, &e.CompanyNameEn, &e.WebsiteUrl, &e.IndustryName, &e.CityNameCn, &e.CityNameEn, &e.ProvinceNameCn, &e.ProvinceNameEn, &e.CreateTime, &e.UpdateTime)
		if err != nil {
			return nil, err
		}
		list = append(list, &e)
	}
	if rows.Err() != nil {
		err = rows.Err()
		return nil, err
	}

	return list, nil
}

func (dao *StockDao) SelectAll(ctx context.Context, tx *runtime.Tx) (list []*Stock, err error) {
	stmt := dao.selectStmtAll
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	rows, err := stmt.Query(ctx)
	if err != nil {
		dao.logger.Error("sqlDriver", zap.Error(err))
		return nil, err
	}

	return dao.ScanRows(rows)
}

func (dao *StockDao) Select(ctx context.Context, tx *runtime.Tx, query string) (*Stock, error) {
	row := dao.db.QueryRow(ctx, "SELECT "+STOCK_ALL_FIELDS_STRING+" FROM stock "+query)
	return dao.ScanRow(row)
}

func (dao *StockDao) SelectList(ctx context.Context, tx *runtime.Tx, query string) (list []*Stock, err error) {
	rows, err := dao.db.Query(ctx, "SELECT "+STOCK_ALL_FIELDS_STRING+" FROM stock "+query)
	if err != nil {
		dao.logger.Error("sqlDriver", zap.Error(err))
		return nil, err
	}

	return dao.ScanRows(rows)
}

func (dao *StockDao) SelectById(ctx context.Context, tx *runtime.Tx, Id int64) (*Stock, error) {
	stmt := dao.selectStmtById
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	return dao.ScanRow(stmt.QueryRow(ctx, Id))
}

func (dao *StockDao) SelectByUpdateTime(ctx context.Context, tx *runtime.Tx, UpdateTime string) (*Stock, error) {
	stmt := dao.selectStmtByUpdateTime
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	return dao.ScanRow(stmt.QueryRow(ctx, UpdateTime))
}

func (dao *StockDao) SelectListByUpdateTime(ctx context.Context, tx *runtime.Tx, UpdateTime string) (list []*Stock, err error) {
	stmt := dao.selectStmtByUpdateTime
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	rows, err := stmt.Query(ctx, UpdateTime)
	if err != nil {
		dao.logger.Error("sqlDriver", zap.Error(err))
		return nil, err
	}

	return dao.ScanRows(rows)
}

func (dao *StockDao) SelectByIndustryName(ctx context.Context, tx *runtime.Tx, IndustryName string) (*Stock, error) {
	stmt := dao.selectStmtByIndustryName
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	return dao.ScanRow(stmt.QueryRow(ctx, IndustryName))
}

func (dao *StockDao) SelectListByIndustryName(ctx context.Context, tx *runtime.Tx, IndustryName string) (list []*Stock, err error) {
	stmt := dao.selectStmtByIndustryName
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	rows, err := stmt.Query(ctx, IndustryName)
	if err != nil {
		dao.logger.Error("sqlDriver", zap.Error(err))
		return nil, err
	}

	return dao.ScanRows(rows)
}

func (dao *StockDao) SelectByStockId(ctx context.Context, tx *runtime.Tx, StockId string) (*Stock, error) {
	stmt := dao.selectStmtByStockId
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	return dao.ScanRow(stmt.QueryRow(ctx, StockId))
}

func (dao *StockDao) SelectByExchangeIdAndStockCode(ctx context.Context, tx *runtime.Tx, ExchangeId string, StockCode string) (*Stock, error) {
	stmt := dao.selectStmtByExchangeIdAndStockCode
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	return dao.ScanRow(stmt.QueryRow(ctx, ExchangeId, StockCode))
}

func (dao *StockDao) SelectListByExchangeId(ctx context.Context, tx *runtime.Tx, ExchangeId string) (list []*Stock, err error) {
	stmt := dao.selectStmtByExchangeId
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	rows, err := stmt.Query(ctx, ExchangeId)
	if err != nil {
		dao.logger.Error("sqlDriver", zap.Error(err))
		return nil, err
	}

	return dao.ScanRows(rows)
}

func (dao *StockDao) GetQuery() *StockQuery {
	return NewStockQuery(dao)
}
