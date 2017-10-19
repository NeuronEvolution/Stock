package fin_stock

import (
	"bytes"
	"context"
	"fmt"
	"github.com/NeuronEvolution/log"
	"github.com/NeuronEvolution/sql/wrap"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"time"
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
	CreateTime     time.Time
	UpdateTime     time.Time
}

type ExchangeQuery struct {
	dao         *ExchangeDao
	forUpdate   bool
	forShare    bool
	whereBuffer *bytes.Buffer
	limitBuffer *bytes.Buffer
	orderBuffer *bytes.Buffer
}

func NewExchangeQuery(dao *ExchangeDao) *ExchangeQuery {
	q := &ExchangeQuery{}
	q.dao = dao
	q.whereBuffer = bytes.NewBufferString("")
	q.limitBuffer = bytes.NewBufferString("")
	q.orderBuffer = bytes.NewBufferString("")

	return q
}

func (q *ExchangeQuery) buildQueryString() string {
	buf := bytes.NewBufferString("")

	if q.forShare {
		buf.WriteString(" FOR UPDATE ")
	}

	if q.forUpdate {
		buf.WriteString(" LOCK IN SHARE MODE ")
	}

	whereSql := q.whereBuffer.String()
	if whereSql != "" {
		buf.WriteString(" WHERE ")
		buf.WriteString(whereSql)
	}

	limitSql := q.limitBuffer.String()
	if limitSql != "" {
		buf.WriteString(limitSql)
	}

	orderSql := q.orderBuffer.String()
	if orderSql != "" {
		buf.WriteString(orderSql)
	}

	return buf.String()
}

func (q *ExchangeQuery) Select(ctx context.Context) (*Exchange, error) {
	return q.dao.doSelect(ctx, nil, q.buildQueryString())
}

func (q *ExchangeQuery) SelectForUpdate(ctx context.Context, tx *wrap.Tx) (*Exchange, error) {
	q.forUpdate = true
	return q.dao.doSelect(ctx, tx, q.buildQueryString())
}

func (q *ExchangeQuery) SelectForShare(ctx context.Context, tx *wrap.Tx) (*Exchange, error) {
	q.forShare = true
	return q.dao.doSelect(ctx, tx, q.buildQueryString())
}

func (q *ExchangeQuery) SelectList(ctx context.Context) (list []*Exchange, err error) {
	return q.dao.doSelectList(ctx, nil, q.buildQueryString())
}

func (q *ExchangeQuery) SelectListForUpdate(ctx context.Context, tx *wrap.Tx) (list []*Exchange, err error) {
	q.forUpdate = true
	return q.dao.doSelectList(ctx, tx, q.buildQueryString())
}

func (q *ExchangeQuery) SelectListForShare(ctx context.Context, tx *wrap.Tx) (list []*Exchange, err error) {
	q.forShare = true
	return q.dao.doSelectList(ctx, tx, q.buildQueryString())
}

func (q *ExchangeQuery) Left() *ExchangeQuery {
	q.whereBuffer.WriteString(" ( ")
	return q
}

func (q *ExchangeQuery) Right() *ExchangeQuery {
	q.whereBuffer.WriteString(" ) ")
	return q
}

func (q *ExchangeQuery) And() *ExchangeQuery {
	q.whereBuffer.WriteString(" AND ")
	return q
}

func (q *ExchangeQuery) Or() *ExchangeQuery {
	q.whereBuffer.WriteString(" OR ")
	return q
}

func (q *ExchangeQuery) Not() *ExchangeQuery {
	q.whereBuffer.WriteString(" NOT ")
	return q
}

func (q *ExchangeQuery) Limit(startIncluded int64, count int64) *ExchangeQuery {
	q.limitBuffer.WriteString(fmt.Sprintf(" limit %d,%d", startIncluded, count))
	return q
}

func (q *ExchangeQuery) Sort(fieldName string, asc bool) *ExchangeQuery {
	if asc {
		q.orderBuffer.WriteString(fmt.Sprintf(" order by %s asc", fieldName))
	} else {
		q.orderBuffer.WriteString(fmt.Sprintf(" order by %s desc", fieldName))
	}

	return q
}
func (q *ExchangeQuery) Id_Equal(v int64) *ExchangeQuery {
	q.whereBuffer.WriteString("id='" + fmt.Sprint(v) + "'")
	return q
}

func (q *ExchangeQuery) Id_NotEqual(v int64) *ExchangeQuery {
	q.whereBuffer.WriteString("id<>'" + fmt.Sprint(v) + "'")
	return q
}

func (q *ExchangeQuery) Id_Less(v int64) *ExchangeQuery {
	q.whereBuffer.WriteString("id<'" + fmt.Sprint(v) + "'")
	return q
}

func (q *ExchangeQuery) Id_LessEqual(v int64) *ExchangeQuery {
	q.whereBuffer.WriteString("id<='" + fmt.Sprint(v) + "'")
	return q
}

func (q *ExchangeQuery) Id_Greater(v int64) *ExchangeQuery {
	q.whereBuffer.WriteString("id>='" + fmt.Sprint(v) + "'")
	return q
}

func (q *ExchangeQuery) Id_GreaterEqual(v int64) *ExchangeQuery {
	q.whereBuffer.WriteString("id='" + fmt.Sprint(v) + "'")
	return q
}

func (q *ExchangeQuery) ExchangeId_Equal(v string) *ExchangeQuery {
	q.whereBuffer.WriteString("exchange_id='" + fmt.Sprint(v) + "'")
	return q
}

func (q *ExchangeQuery) ExchangeId_NotEqual(v string) *ExchangeQuery {
	q.whereBuffer.WriteString("exchange_id<>'" + fmt.Sprint(v) + "'")
	return q
}

func (q *ExchangeQuery) ExchangeId_Less(v string) *ExchangeQuery {
	q.whereBuffer.WriteString("exchange_id<'" + fmt.Sprint(v) + "'")
	return q
}

func (q *ExchangeQuery) ExchangeId_LessEqual(v string) *ExchangeQuery {
	q.whereBuffer.WriteString("exchange_id<='" + fmt.Sprint(v) + "'")
	return q
}

func (q *ExchangeQuery) ExchangeId_Greater(v string) *ExchangeQuery {
	q.whereBuffer.WriteString("exchange_id>='" + fmt.Sprint(v) + "'")
	return q
}

func (q *ExchangeQuery) ExchangeId_GreaterEqual(v string) *ExchangeQuery {
	q.whereBuffer.WriteString("exchange_id='" + fmt.Sprint(v) + "'")
	return q
}

func (q *ExchangeQuery) ExchangeNameCn_Equal(v string) *ExchangeQuery {
	q.whereBuffer.WriteString("exchange_name_cn='" + fmt.Sprint(v) + "'")
	return q
}

func (q *ExchangeQuery) ExchangeNameCn_NotEqual(v string) *ExchangeQuery {
	q.whereBuffer.WriteString("exchange_name_cn<>'" + fmt.Sprint(v) + "'")
	return q
}

func (q *ExchangeQuery) ExchangeNameCn_Less(v string) *ExchangeQuery {
	q.whereBuffer.WriteString("exchange_name_cn<'" + fmt.Sprint(v) + "'")
	return q
}

func (q *ExchangeQuery) ExchangeNameCn_LessEqual(v string) *ExchangeQuery {
	q.whereBuffer.WriteString("exchange_name_cn<='" + fmt.Sprint(v) + "'")
	return q
}

func (q *ExchangeQuery) ExchangeNameCn_Greater(v string) *ExchangeQuery {
	q.whereBuffer.WriteString("exchange_name_cn>='" + fmt.Sprint(v) + "'")
	return q
}

func (q *ExchangeQuery) ExchangeNameCn_GreaterEqual(v string) *ExchangeQuery {
	q.whereBuffer.WriteString("exchange_name_cn='" + fmt.Sprint(v) + "'")
	return q
}

func (q *ExchangeQuery) ExchangeNameEn_Equal(v string) *ExchangeQuery {
	q.whereBuffer.WriteString("exchange_name_en='" + fmt.Sprint(v) + "'")
	return q
}

func (q *ExchangeQuery) ExchangeNameEn_NotEqual(v string) *ExchangeQuery {
	q.whereBuffer.WriteString("exchange_name_en<>'" + fmt.Sprint(v) + "'")
	return q
}

func (q *ExchangeQuery) ExchangeNameEn_Less(v string) *ExchangeQuery {
	q.whereBuffer.WriteString("exchange_name_en<'" + fmt.Sprint(v) + "'")
	return q
}

func (q *ExchangeQuery) ExchangeNameEn_LessEqual(v string) *ExchangeQuery {
	q.whereBuffer.WriteString("exchange_name_en<='" + fmt.Sprint(v) + "'")
	return q
}

func (q *ExchangeQuery) ExchangeNameEn_Greater(v string) *ExchangeQuery {
	q.whereBuffer.WriteString("exchange_name_en>='" + fmt.Sprint(v) + "'")
	return q
}

func (q *ExchangeQuery) ExchangeNameEn_GreaterEqual(v string) *ExchangeQuery {
	q.whereBuffer.WriteString("exchange_name_en='" + fmt.Sprint(v) + "'")
	return q
}

func (q *ExchangeQuery) CreateTime_Equal(v time.Time) *ExchangeQuery {
	q.whereBuffer.WriteString("create_time='" + fmt.Sprint(v) + "'")
	return q
}

func (q *ExchangeQuery) CreateTime_NotEqual(v time.Time) *ExchangeQuery {
	q.whereBuffer.WriteString("create_time<>'" + fmt.Sprint(v) + "'")
	return q
}

func (q *ExchangeQuery) CreateTime_Less(v time.Time) *ExchangeQuery {
	q.whereBuffer.WriteString("create_time<'" + fmt.Sprint(v) + "'")
	return q
}

func (q *ExchangeQuery) CreateTime_LessEqual(v time.Time) *ExchangeQuery {
	q.whereBuffer.WriteString("create_time<='" + fmt.Sprint(v) + "'")
	return q
}

func (q *ExchangeQuery) CreateTime_Greater(v time.Time) *ExchangeQuery {
	q.whereBuffer.WriteString("create_time>='" + fmt.Sprint(v) + "'")
	return q
}

func (q *ExchangeQuery) CreateTime_GreaterEqual(v time.Time) *ExchangeQuery {
	q.whereBuffer.WriteString("create_time='" + fmt.Sprint(v) + "'")
	return q
}

func (q *ExchangeQuery) UpdateTime_Equal(v time.Time) *ExchangeQuery {
	q.whereBuffer.WriteString("update_time='" + fmt.Sprint(v) + "'")
	return q
}

func (q *ExchangeQuery) UpdateTime_NotEqual(v time.Time) *ExchangeQuery {
	q.whereBuffer.WriteString("update_time<>'" + fmt.Sprint(v) + "'")
	return q
}

func (q *ExchangeQuery) UpdateTime_Less(v time.Time) *ExchangeQuery {
	q.whereBuffer.WriteString("update_time<'" + fmt.Sprint(v) + "'")
	return q
}

func (q *ExchangeQuery) UpdateTime_LessEqual(v time.Time) *ExchangeQuery {
	q.whereBuffer.WriteString("update_time<='" + fmt.Sprint(v) + "'")
	return q
}

func (q *ExchangeQuery) UpdateTime_Greater(v time.Time) *ExchangeQuery {
	q.whereBuffer.WriteString("update_time>='" + fmt.Sprint(v) + "'")
	return q
}

func (q *ExchangeQuery) UpdateTime_GreaterEqual(v time.Time) *ExchangeQuery {
	q.whereBuffer.WriteString("update_time='" + fmt.Sprint(v) + "'")
	return q
}

type ExchangeDao struct {
	logger     *zap.Logger
	db         *DB
	insertStmt *wrap.Stmt
	updateStmt *wrap.Stmt
	deleteStmt *wrap.Stmt
}

func NewExchangeDao(db *DB) (t *ExchangeDao, err error) {
	t = &ExchangeDao{}
	t.logger = log.TypedLogger(t)
	t.db = db
	err = t.init()
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (dao *ExchangeDao) init() (err error) {
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

func (dao *ExchangeDao) Insert(ctx context.Context, tx *wrap.Tx, e *Exchange) (id int64, err error) {
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

func (dao *ExchangeDao) Update(ctx context.Context, tx *wrap.Tx, e *Exchange) (err error) {
	stmt := dao.updateStmt
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	result, err := stmt.Exec(ctx, e.ExchangeId, e.ExchangeNameCn, e.ExchangeNameEn, e.CreateTime, e.UpdateTime, e.Id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return fmt.Errorf("rowsAffected:%s", rowsAffected)
	}

	return nil
}

func (dao *ExchangeDao) Delete(ctx context.Context, tx *wrap.Tx, id int64) (err error) {
	stmt := dao.deleteStmt
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	result, err := stmt.Exec(ctx, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return fmt.Errorf("rowsAffected:%s", rowsAffected)
	}

	return nil
}

func (dao *ExchangeDao) scanRow(row *wrap.Row) (*Exchange, error) {
	e := &Exchange{}
	err := row.Scan(&e.Id, &e.ExchangeId, &e.ExchangeNameCn, &e.ExchangeNameEn, &e.CreateTime, &e.UpdateTime)
	if err != nil {
		if err == wrap.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return e, nil
}

func (dao *ExchangeDao) scanRows(rows *wrap.Rows) (list []*Exchange, err error) {
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

func (dao *ExchangeDao) doSelect(ctx context.Context, tx *wrap.Tx, query string) (*Exchange, error) {
	row := dao.db.QueryRow(ctx, "SELECT "+EXCHANGE_ALL_FIELDS_STRING+" FROM exchange "+query)
	return dao.scanRow(row)
}

func (dao *ExchangeDao) doSelectList(ctx context.Context, tx *wrap.Tx, query string) (list []*Exchange, err error) {
	rows, err := dao.db.Query(ctx, "SELECT "+EXCHANGE_ALL_FIELDS_STRING+" FROM exchange "+query)
	if err != nil {
		dao.logger.Error("sqlDriver", zap.Error(err))
		return nil, err
	}

	return dao.scanRows(rows)
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
	LaunchDate     time.Time
	CompanyNameCn  string //size=128
	CompanyNameEn  string //size=128
	WebsiteUrl     string //size=128
	IndustryName   string //size=32
	CityNameCn     string //size=128
	CityNameEn     string //size=128
	ProvinceNameCn string //size=128
	ProvinceNameEn string //size=128
	CreateTime     time.Time
	UpdateTime     time.Time
}

type StockQuery struct {
	dao         *StockDao
	forUpdate   bool
	forShare    bool
	whereBuffer *bytes.Buffer
	limitBuffer *bytes.Buffer
	orderBuffer *bytes.Buffer
}

func NewStockQuery(dao *StockDao) *StockQuery {
	q := &StockQuery{}
	q.dao = dao
	q.whereBuffer = bytes.NewBufferString("")
	q.limitBuffer = bytes.NewBufferString("")
	q.orderBuffer = bytes.NewBufferString("")

	return q
}

func (q *StockQuery) buildQueryString() string {
	buf := bytes.NewBufferString("")

	if q.forShare {
		buf.WriteString(" FOR UPDATE ")
	}

	if q.forUpdate {
		buf.WriteString(" LOCK IN SHARE MODE ")
	}

	whereSql := q.whereBuffer.String()
	if whereSql != "" {
		buf.WriteString(" WHERE ")
		buf.WriteString(whereSql)
	}

	limitSql := q.limitBuffer.String()
	if limitSql != "" {
		buf.WriteString(limitSql)
	}

	orderSql := q.orderBuffer.String()
	if orderSql != "" {
		buf.WriteString(orderSql)
	}

	return buf.String()
}

func (q *StockQuery) Select(ctx context.Context) (*Stock, error) {
	return q.dao.doSelect(ctx, nil, q.buildQueryString())
}

func (q *StockQuery) SelectForUpdate(ctx context.Context, tx *wrap.Tx) (*Stock, error) {
	q.forUpdate = true
	return q.dao.doSelect(ctx, tx, q.buildQueryString())
}

func (q *StockQuery) SelectForShare(ctx context.Context, tx *wrap.Tx) (*Stock, error) {
	q.forShare = true
	return q.dao.doSelect(ctx, tx, q.buildQueryString())
}

func (q *StockQuery) SelectList(ctx context.Context) (list []*Stock, err error) {
	return q.dao.doSelectList(ctx, nil, q.buildQueryString())
}

func (q *StockQuery) SelectListForUpdate(ctx context.Context, tx *wrap.Tx) (list []*Stock, err error) {
	q.forUpdate = true
	return q.dao.doSelectList(ctx, tx, q.buildQueryString())
}

func (q *StockQuery) SelectListForShare(ctx context.Context, tx *wrap.Tx) (list []*Stock, err error) {
	q.forShare = true
	return q.dao.doSelectList(ctx, tx, q.buildQueryString())
}

func (q *StockQuery) Left() *StockQuery {
	q.whereBuffer.WriteString(" ( ")
	return q
}

func (q *StockQuery) Right() *StockQuery {
	q.whereBuffer.WriteString(" ) ")
	return q
}

func (q *StockQuery) And() *StockQuery {
	q.whereBuffer.WriteString(" AND ")
	return q
}

func (q *StockQuery) Or() *StockQuery {
	q.whereBuffer.WriteString(" OR ")
	return q
}

func (q *StockQuery) Not() *StockQuery {
	q.whereBuffer.WriteString(" NOT ")
	return q
}

func (q *StockQuery) Limit(startIncluded int64, count int64) *StockQuery {
	q.limitBuffer.WriteString(fmt.Sprintf(" limit %d,%d", startIncluded, count))
	return q
}

func (q *StockQuery) Sort(fieldName string, asc bool) *StockQuery {
	if asc {
		q.orderBuffer.WriteString(fmt.Sprintf(" order by %s asc", fieldName))
	} else {
		q.orderBuffer.WriteString(fmt.Sprintf(" order by %s desc", fieldName))
	}

	return q
}
func (q *StockQuery) Id_Equal(v int64) *StockQuery {
	q.whereBuffer.WriteString("id='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) Id_NotEqual(v int64) *StockQuery {
	q.whereBuffer.WriteString("id<>'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) Id_Less(v int64) *StockQuery {
	q.whereBuffer.WriteString("id<'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) Id_LessEqual(v int64) *StockQuery {
	q.whereBuffer.WriteString("id<='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) Id_Greater(v int64) *StockQuery {
	q.whereBuffer.WriteString("id>='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) Id_GreaterEqual(v int64) *StockQuery {
	q.whereBuffer.WriteString("id='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) StockId_Equal(v string) *StockQuery {
	q.whereBuffer.WriteString("stock_id='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) StockId_NotEqual(v string) *StockQuery {
	q.whereBuffer.WriteString("stock_id<>'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) StockId_Less(v string) *StockQuery {
	q.whereBuffer.WriteString("stock_id<'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) StockId_LessEqual(v string) *StockQuery {
	q.whereBuffer.WriteString("stock_id<='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) StockId_Greater(v string) *StockQuery {
	q.whereBuffer.WriteString("stock_id>='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) StockId_GreaterEqual(v string) *StockQuery {
	q.whereBuffer.WriteString("stock_id='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) ExchangeId_Equal(v string) *StockQuery {
	q.whereBuffer.WriteString("exchange_id='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) ExchangeId_NotEqual(v string) *StockQuery {
	q.whereBuffer.WriteString("exchange_id<>'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) ExchangeId_Less(v string) *StockQuery {
	q.whereBuffer.WriteString("exchange_id<'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) ExchangeId_LessEqual(v string) *StockQuery {
	q.whereBuffer.WriteString("exchange_id<='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) ExchangeId_Greater(v string) *StockQuery {
	q.whereBuffer.WriteString("exchange_id>='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) ExchangeId_GreaterEqual(v string) *StockQuery {
	q.whereBuffer.WriteString("exchange_id='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) StockCode_Equal(v string) *StockQuery {
	q.whereBuffer.WriteString("stock_code='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) StockCode_NotEqual(v string) *StockQuery {
	q.whereBuffer.WriteString("stock_code<>'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) StockCode_Less(v string) *StockQuery {
	q.whereBuffer.WriteString("stock_code<'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) StockCode_LessEqual(v string) *StockQuery {
	q.whereBuffer.WriteString("stock_code<='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) StockCode_Greater(v string) *StockQuery {
	q.whereBuffer.WriteString("stock_code>='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) StockCode_GreaterEqual(v string) *StockQuery {
	q.whereBuffer.WriteString("stock_code='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) StockNameCn_Equal(v string) *StockQuery {
	q.whereBuffer.WriteString("stock_name_cn='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) StockNameCn_NotEqual(v string) *StockQuery {
	q.whereBuffer.WriteString("stock_name_cn<>'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) StockNameCn_Less(v string) *StockQuery {
	q.whereBuffer.WriteString("stock_name_cn<'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) StockNameCn_LessEqual(v string) *StockQuery {
	q.whereBuffer.WriteString("stock_name_cn<='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) StockNameCn_Greater(v string) *StockQuery {
	q.whereBuffer.WriteString("stock_name_cn>='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) StockNameCn_GreaterEqual(v string) *StockQuery {
	q.whereBuffer.WriteString("stock_name_cn='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) StockNameEn_Equal(v string) *StockQuery {
	q.whereBuffer.WriteString("stock_name_en='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) StockNameEn_NotEqual(v string) *StockQuery {
	q.whereBuffer.WriteString("stock_name_en<>'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) StockNameEn_Less(v string) *StockQuery {
	q.whereBuffer.WriteString("stock_name_en<'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) StockNameEn_LessEqual(v string) *StockQuery {
	q.whereBuffer.WriteString("stock_name_en<='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) StockNameEn_Greater(v string) *StockQuery {
	q.whereBuffer.WriteString("stock_name_en>='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) StockNameEn_GreaterEqual(v string) *StockQuery {
	q.whereBuffer.WriteString("stock_name_en='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) LaunchDate_Equal(v time.Time) *StockQuery {
	q.whereBuffer.WriteString("launch_date='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) LaunchDate_NotEqual(v time.Time) *StockQuery {
	q.whereBuffer.WriteString("launch_date<>'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) LaunchDate_Less(v time.Time) *StockQuery {
	q.whereBuffer.WriteString("launch_date<'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) LaunchDate_LessEqual(v time.Time) *StockQuery {
	q.whereBuffer.WriteString("launch_date<='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) LaunchDate_Greater(v time.Time) *StockQuery {
	q.whereBuffer.WriteString("launch_date>='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) LaunchDate_GreaterEqual(v time.Time) *StockQuery {
	q.whereBuffer.WriteString("launch_date='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) CompanyNameCn_Equal(v string) *StockQuery {
	q.whereBuffer.WriteString("company_name_cn='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) CompanyNameCn_NotEqual(v string) *StockQuery {
	q.whereBuffer.WriteString("company_name_cn<>'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) CompanyNameCn_Less(v string) *StockQuery {
	q.whereBuffer.WriteString("company_name_cn<'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) CompanyNameCn_LessEqual(v string) *StockQuery {
	q.whereBuffer.WriteString("company_name_cn<='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) CompanyNameCn_Greater(v string) *StockQuery {
	q.whereBuffer.WriteString("company_name_cn>='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) CompanyNameCn_GreaterEqual(v string) *StockQuery {
	q.whereBuffer.WriteString("company_name_cn='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) CompanyNameEn_Equal(v string) *StockQuery {
	q.whereBuffer.WriteString("company_name_en='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) CompanyNameEn_NotEqual(v string) *StockQuery {
	q.whereBuffer.WriteString("company_name_en<>'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) CompanyNameEn_Less(v string) *StockQuery {
	q.whereBuffer.WriteString("company_name_en<'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) CompanyNameEn_LessEqual(v string) *StockQuery {
	q.whereBuffer.WriteString("company_name_en<='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) CompanyNameEn_Greater(v string) *StockQuery {
	q.whereBuffer.WriteString("company_name_en>='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) CompanyNameEn_GreaterEqual(v string) *StockQuery {
	q.whereBuffer.WriteString("company_name_en='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) WebsiteUrl_Equal(v string) *StockQuery {
	q.whereBuffer.WriteString("website_url='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) WebsiteUrl_NotEqual(v string) *StockQuery {
	q.whereBuffer.WriteString("website_url<>'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) WebsiteUrl_Less(v string) *StockQuery {
	q.whereBuffer.WriteString("website_url<'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) WebsiteUrl_LessEqual(v string) *StockQuery {
	q.whereBuffer.WriteString("website_url<='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) WebsiteUrl_Greater(v string) *StockQuery {
	q.whereBuffer.WriteString("website_url>='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) WebsiteUrl_GreaterEqual(v string) *StockQuery {
	q.whereBuffer.WriteString("website_url='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) IndustryName_Equal(v string) *StockQuery {
	q.whereBuffer.WriteString("industry_name='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) IndustryName_NotEqual(v string) *StockQuery {
	q.whereBuffer.WriteString("industry_name<>'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) IndustryName_Less(v string) *StockQuery {
	q.whereBuffer.WriteString("industry_name<'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) IndustryName_LessEqual(v string) *StockQuery {
	q.whereBuffer.WriteString("industry_name<='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) IndustryName_Greater(v string) *StockQuery {
	q.whereBuffer.WriteString("industry_name>='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) IndustryName_GreaterEqual(v string) *StockQuery {
	q.whereBuffer.WriteString("industry_name='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) CityNameCn_Equal(v string) *StockQuery {
	q.whereBuffer.WriteString("city_name_cn='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) CityNameCn_NotEqual(v string) *StockQuery {
	q.whereBuffer.WriteString("city_name_cn<>'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) CityNameCn_Less(v string) *StockQuery {
	q.whereBuffer.WriteString("city_name_cn<'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) CityNameCn_LessEqual(v string) *StockQuery {
	q.whereBuffer.WriteString("city_name_cn<='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) CityNameCn_Greater(v string) *StockQuery {
	q.whereBuffer.WriteString("city_name_cn>='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) CityNameCn_GreaterEqual(v string) *StockQuery {
	q.whereBuffer.WriteString("city_name_cn='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) CityNameEn_Equal(v string) *StockQuery {
	q.whereBuffer.WriteString("city_name_en='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) CityNameEn_NotEqual(v string) *StockQuery {
	q.whereBuffer.WriteString("city_name_en<>'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) CityNameEn_Less(v string) *StockQuery {
	q.whereBuffer.WriteString("city_name_en<'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) CityNameEn_LessEqual(v string) *StockQuery {
	q.whereBuffer.WriteString("city_name_en<='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) CityNameEn_Greater(v string) *StockQuery {
	q.whereBuffer.WriteString("city_name_en>='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) CityNameEn_GreaterEqual(v string) *StockQuery {
	q.whereBuffer.WriteString("city_name_en='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) ProvinceNameCn_Equal(v string) *StockQuery {
	q.whereBuffer.WriteString("province_name_cn='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) ProvinceNameCn_NotEqual(v string) *StockQuery {
	q.whereBuffer.WriteString("province_name_cn<>'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) ProvinceNameCn_Less(v string) *StockQuery {
	q.whereBuffer.WriteString("province_name_cn<'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) ProvinceNameCn_LessEqual(v string) *StockQuery {
	q.whereBuffer.WriteString("province_name_cn<='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) ProvinceNameCn_Greater(v string) *StockQuery {
	q.whereBuffer.WriteString("province_name_cn>='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) ProvinceNameCn_GreaterEqual(v string) *StockQuery {
	q.whereBuffer.WriteString("province_name_cn='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) ProvinceNameEn_Equal(v string) *StockQuery {
	q.whereBuffer.WriteString("province_name_en='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) ProvinceNameEn_NotEqual(v string) *StockQuery {
	q.whereBuffer.WriteString("province_name_en<>'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) ProvinceNameEn_Less(v string) *StockQuery {
	q.whereBuffer.WriteString("province_name_en<'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) ProvinceNameEn_LessEqual(v string) *StockQuery {
	q.whereBuffer.WriteString("province_name_en<='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) ProvinceNameEn_Greater(v string) *StockQuery {
	q.whereBuffer.WriteString("province_name_en>='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) ProvinceNameEn_GreaterEqual(v string) *StockQuery {
	q.whereBuffer.WriteString("province_name_en='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) CreateTime_Equal(v time.Time) *StockQuery {
	q.whereBuffer.WriteString("create_time='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) CreateTime_NotEqual(v time.Time) *StockQuery {
	q.whereBuffer.WriteString("create_time<>'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) CreateTime_Less(v time.Time) *StockQuery {
	q.whereBuffer.WriteString("create_time<'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) CreateTime_LessEqual(v time.Time) *StockQuery {
	q.whereBuffer.WriteString("create_time<='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) CreateTime_Greater(v time.Time) *StockQuery {
	q.whereBuffer.WriteString("create_time>='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) CreateTime_GreaterEqual(v time.Time) *StockQuery {
	q.whereBuffer.WriteString("create_time='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) UpdateTime_Equal(v time.Time) *StockQuery {
	q.whereBuffer.WriteString("update_time='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) UpdateTime_NotEqual(v time.Time) *StockQuery {
	q.whereBuffer.WriteString("update_time<>'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) UpdateTime_Less(v time.Time) *StockQuery {
	q.whereBuffer.WriteString("update_time<'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) UpdateTime_LessEqual(v time.Time) *StockQuery {
	q.whereBuffer.WriteString("update_time<='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) UpdateTime_Greater(v time.Time) *StockQuery {
	q.whereBuffer.WriteString("update_time>='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockQuery) UpdateTime_GreaterEqual(v time.Time) *StockQuery {
	q.whereBuffer.WriteString("update_time='" + fmt.Sprint(v) + "'")
	return q
}

type StockDao struct {
	logger     *zap.Logger
	db         *DB
	insertStmt *wrap.Stmt
	updateStmt *wrap.Stmt
	deleteStmt *wrap.Stmt
}

func NewStockDao(db *DB) (t *StockDao, err error) {
	t = &StockDao{}
	t.logger = log.TypedLogger(t)
	t.db = db
	err = t.init()
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (dao *StockDao) init() (err error) {
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

func (dao *StockDao) Insert(ctx context.Context, tx *wrap.Tx, e *Stock) (id int64, err error) {
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

func (dao *StockDao) Update(ctx context.Context, tx *wrap.Tx, e *Stock) (err error) {
	stmt := dao.updateStmt
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	result, err := stmt.Exec(ctx, e.StockId, e.ExchangeId, e.StockCode, e.StockNameCn, e.StockNameEn, e.LaunchDate, e.CompanyNameCn, e.CompanyNameEn, e.WebsiteUrl, e.IndustryName, e.CityNameCn, e.CityNameEn, e.ProvinceNameCn, e.ProvinceNameEn, e.CreateTime, e.UpdateTime, e.Id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return fmt.Errorf("rowsAffected:%s", rowsAffected)
	}

	return nil
}

func (dao *StockDao) Delete(ctx context.Context, tx *wrap.Tx, id int64) (err error) {
	stmt := dao.deleteStmt
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	result, err := stmt.Exec(ctx, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return fmt.Errorf("rowsAffected:%s", rowsAffected)
	}

	return nil
}

func (dao *StockDao) scanRow(row *wrap.Row) (*Stock, error) {
	e := &Stock{}
	err := row.Scan(&e.Id, &e.StockId, &e.ExchangeId, &e.StockCode, &e.StockNameCn, &e.StockNameEn, &e.LaunchDate, &e.CompanyNameCn, &e.CompanyNameEn, &e.WebsiteUrl, &e.IndustryName, &e.CityNameCn, &e.CityNameEn, &e.ProvinceNameCn, &e.ProvinceNameEn, &e.CreateTime, &e.UpdateTime)
	if err != nil {
		if err == wrap.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return e, nil
}

func (dao *StockDao) scanRows(rows *wrap.Rows) (list []*Stock, err error) {
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

func (dao *StockDao) doSelect(ctx context.Context, tx *wrap.Tx, query string) (*Stock, error) {
	row := dao.db.QueryRow(ctx, "SELECT "+STOCK_ALL_FIELDS_STRING+" FROM stock "+query)
	return dao.scanRow(row)
}

func (dao *StockDao) doSelectList(ctx context.Context, tx *wrap.Tx, query string) (list []*Stock, err error) {
	rows, err := dao.db.Query(ctx, "SELECT "+STOCK_ALL_FIELDS_STRING+" FROM stock "+query)
	if err != nil {
		dao.logger.Error("sqlDriver", zap.Error(err))
		return nil, err
	}

	return dao.scanRows(rows)
}

func (dao *StockDao) GetQuery() *StockQuery {
	return NewStockQuery(dao)
}

type DB struct {
	wrap.DB
	Exchange *ExchangeDao
	Stock    *StockDao
}

func NewDB(connectionString string) (d *DB, err error) {
	if connectionString == "" {
		return nil, fmt.Errorf("connectionString nil")
	}

	d = &DB{}

	db, err := wrap.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	d.DB = *db

	err = d.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	d.Exchange, err = NewExchangeDao(d)
	if err != nil {
		return nil, err
	}

	d.Stock, err = NewStockDao(d)
	if err != nil {
		return nil, err
	}

	return d, nil
}
