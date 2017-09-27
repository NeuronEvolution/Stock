package gen

import (
	"database/sql"
)

const EXCHANGE_TABLE_NAME = "exchange"

const EXCHANGE_FIELD_ID = "id"
const EXCHANGE_FIELD_EXCHANGE_ID = "exchange_id"
const EXCHANGE_FIELD_EXCHANGE_NAME_CN = "exchange_name_cn"
const EXCHANGE_FIELD_EXCHANGE_NAME_EN = "exchange_name_en"
const EXCHANGE_FIELD_CREATE_TIME = "create_time"
const EXCHANGE_FIELD_UPDATE_TIME = "update_time"

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

type ExchangeDao struct {
	db                     *sql.DB
	insertStmt             *sql.Stmt
	selectStmtAll          *sql.Stmt
	selectStmtById         *sql.Stmt
	selectStmtByUpdateTime *sql.Stmt
	selectStmtByExchangeId *sql.Stmt
	updateStmt             *sql.Stmt
	deleteStmt             *sql.Stmt
}

func NewExchangeDao(db *sql.DB) (t *ExchangeDao) {
	t = &ExchangeDao{}
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
	dao.insertStmt, err = dao.db.Prepare("INSERT INTO exchange (exchange_id,exchange_name_cn,exchange_name_en,create_time,update_time) VALUES (?,?,?,?,?)")
	if err != nil {
		return err
	}

	return nil
}

func (dao *ExchangeDao) prepareUpdateStmt() (err error) {
	dao.updateStmt, err = dao.db.Prepare("UPDATE exchange SET exchange_id=?,exchange_name_cn=?,exchange_name_en=?,create_time=?,update_time=? WHERE id=?")
	if err != nil {
		return err
	}

	return nil
}

func (dao *ExchangeDao) prepareDeleteStmt() (err error) {
	dao.deleteStmt, err = dao.db.Prepare("DELETE FROM exchange WHERE id=?")
	if err != nil {
		return err
	}

	return nil
}

func (dao *ExchangeDao) prepareSelectStmtAll() (err error) {
	dao.selectStmtAll, err = dao.db.Prepare("SELECT id,exchange_id,exchange_name_cn,exchange_name_en,create_time,update_time FROM exchange")
	if err != nil {
		return err
	}

	return nil
}

func (dao *ExchangeDao) prepareSelectStmtById() (err error) {
	dao.selectStmtById, err = dao.db.Prepare("SELECT id,exchange_id,exchange_name_cn,exchange_name_en,create_time,update_time FROM exchange WHERE id=?")
	if err != nil {
		return err
	}

	return nil
}

func (dao *ExchangeDao) prepareSelectStmtByUpdateTime() (err error) {
	dao.selectStmtByUpdateTime, err = dao.db.Prepare("SELECT id,exchange_id,exchange_name_cn,exchange_name_en,create_time,update_time FROM exchange WHERE update_time=?")
	if err != nil {
		return err
	}

	return nil
}

func (dao *ExchangeDao) prepareSelectStmtByExchangeId() (err error) {
	dao.selectStmtByExchangeId, err = dao.db.Prepare("SELECT id,exchange_id,exchange_name_cn,exchange_name_en,create_time,update_time FROM exchange WHERE exchange_id=?")
	if err != nil {
		return err
	}

	return nil
}

func (dao *ExchangeDao) Insert(tx *sql.Tx, e *Exchange) (id int64, err error) {
	stmt := dao.insertStmt
	if tx != nil {
		stmt = tx.Stmt(stmt)
	}

	result, err := stmt.Exec(e.ExchangeId, e.ExchangeNameCn, e.ExchangeNameEn, e.CreateTime, e.UpdateTime)
	if err != nil {
		return 0, err
	}

	id, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (dao *ExchangeDao) Update(tx *sql.Tx, e *Exchange) (int64, error) {
	stmt := dao.updateStmt
	if tx != nil {
		stmt = tx.Stmt(stmt)
	}

	result, err := stmt.Exec(e.ExchangeId, e.ExchangeNameCn, e.ExchangeNameEn, e.CreateTime, e.UpdateTime, e.Id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func (dao *ExchangeDao) Delete(tx *sql.Tx, id int64) (rowsAffected int64, err error) {
	stmt := dao.deleteStmt
	if tx != nil {
		stmt = tx.Stmt(stmt)
	}

	result, err := stmt.Exec(id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func (dao *ExchangeDao) ScanRow(row *sql.Row) (*Exchange, error) {
	e := &Exchange{}
	err := row.Scan(&e.Id, &e.ExchangeId, &e.ExchangeNameCn, &e.ExchangeNameEn, &e.CreateTime, &e.UpdateTime)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return e, nil
}

func (dao *ExchangeDao) ScanRows(rows *sql.Rows) (list []*Exchange, err error) {
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
		return nil, rows.Err()
	}

	return list, nil
}

func (dao *ExchangeDao) SelectAll(tx *sql.Tx) (list []*Exchange, err error) {
	stmt := dao.selectStmtAll
	if tx != nil {
		stmt = tx.Stmt(stmt)
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	return dao.ScanRows(rows)
}

func (dao *ExchangeDao) SelectById(tx *sql.Tx, Id int64) (*Exchange, error) {
	stmt := dao.selectStmtById
	if tx != nil {
		stmt = tx.Stmt(stmt)
	}

	row := stmt.QueryRow(Id)

	return dao.ScanRow(row)
}

func (dao *ExchangeDao) SelectByUpdateTime(tx *sql.Tx, UpdateTime string) (*Exchange, error) {
	stmt := dao.selectStmtByUpdateTime
	if tx != nil {
		stmt = tx.Stmt(stmt)
	}

	row := stmt.QueryRow(UpdateTime)

	return dao.ScanRow(row)
}

func (dao *ExchangeDao) SelectListByUpdateTime(tx *sql.Tx, UpdateTime string) (list []*Exchange, err error) {
	stmt := dao.selectStmtByUpdateTime
	if tx != nil {
		stmt = tx.Stmt(stmt)
	}

	rows, err := stmt.Query(UpdateTime)
	if err != nil {
		return nil, err
	}

	return dao.ScanRows(rows)
}

func (dao *ExchangeDao) SelectByExchangeId(tx *sql.Tx, ExchangeId string) (*Exchange, error) {
	stmt := dao.selectStmtByExchangeId
	if tx != nil {
		stmt = tx.Stmt(stmt)
	}

	row := stmt.QueryRow(ExchangeId)

	return dao.ScanRow(row)
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

type StockDao struct {
	db                                 *sql.DB
	insertStmt                         *sql.Stmt
	selectStmtAll                      *sql.Stmt
	selectStmtById                     *sql.Stmt
	selectStmtByUpdateTime             *sql.Stmt
	selectStmtByIndustryName           *sql.Stmt
	selectStmtByStockId                *sql.Stmt
	selectStmtByExchangeId             *sql.Stmt
	selectStmtByExchangeIdAndStockCode *sql.Stmt
	updateStmt                         *sql.Stmt
	deleteStmt                         *sql.Stmt
}

func NewStockDao(db *sql.DB) (t *StockDao) {
	t = &StockDao{}
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
	dao.insertStmt, err = dao.db.Prepare("INSERT INTO stock (stock_id,exchange_id,stock_code,stock_name_cn,stock_name_en,launch_date,company_name_cn,company_name_en,website_url,industry_name,city_name_cn,city_name_en,province_name_cn,province_name_en,create_time,update_time) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}

	return nil
}

func (dao *StockDao) prepareUpdateStmt() (err error) {
	dao.updateStmt, err = dao.db.Prepare("UPDATE stock SET stock_id=?,exchange_id=?,stock_code=?,stock_name_cn=?,stock_name_en=?,launch_date=?,company_name_cn=?,company_name_en=?,website_url=?,industry_name=?,city_name_cn=?,city_name_en=?,province_name_cn=?,province_name_en=?,create_time=?,update_time=? WHERE id=?")
	if err != nil {
		return err
	}

	return nil
}

func (dao *StockDao) prepareDeleteStmt() (err error) {
	dao.deleteStmt, err = dao.db.Prepare("DELETE FROM stock WHERE id=?")
	if err != nil {
		return err
	}

	return nil
}

func (dao *StockDao) prepareSelectStmtAll() (err error) {
	dao.selectStmtAll, err = dao.db.Prepare("SELECT id,stock_id,exchange_id,stock_code,stock_name_cn,stock_name_en,launch_date,company_name_cn,company_name_en,website_url,industry_name,city_name_cn,city_name_en,province_name_cn,province_name_en,create_time,update_time FROM stock")
	if err != nil {
		return err
	}

	return nil
}

func (dao *StockDao) prepareSelectStmtById() (err error) {
	dao.selectStmtById, err = dao.db.Prepare("SELECT id,stock_id,exchange_id,stock_code,stock_name_cn,stock_name_en,launch_date,company_name_cn,company_name_en,website_url,industry_name,city_name_cn,city_name_en,province_name_cn,province_name_en,create_time,update_time FROM stock WHERE id=?")
	if err != nil {
		return err
	}

	return nil
}

func (dao *StockDao) prepareSelectStmtByUpdateTime() (err error) {
	dao.selectStmtByUpdateTime, err = dao.db.Prepare("SELECT id,stock_id,exchange_id,stock_code,stock_name_cn,stock_name_en,launch_date,company_name_cn,company_name_en,website_url,industry_name,city_name_cn,city_name_en,province_name_cn,province_name_en,create_time,update_time FROM stock WHERE update_time=?")
	if err != nil {
		return err
	}

	return nil
}

func (dao *StockDao) prepareSelectStmtByIndustryName() (err error) {
	dao.selectStmtByIndustryName, err = dao.db.Prepare("SELECT id,stock_id,exchange_id,stock_code,stock_name_cn,stock_name_en,launch_date,company_name_cn,company_name_en,website_url,industry_name,city_name_cn,city_name_en,province_name_cn,province_name_en,create_time,update_time FROM stock WHERE industry_name=?")
	if err != nil {
		return err
	}

	return nil
}

func (dao *StockDao) prepareSelectStmtByStockId() (err error) {
	dao.selectStmtByStockId, err = dao.db.Prepare("SELECT id,stock_id,exchange_id,stock_code,stock_name_cn,stock_name_en,launch_date,company_name_cn,company_name_en,website_url,industry_name,city_name_cn,city_name_en,province_name_cn,province_name_en,create_time,update_time FROM stock WHERE stock_id=?")
	if err != nil {
		return err
	}

	return nil
}

func (dao *StockDao) prepareSelectStmtByExchangeId() (err error) {
	dao.selectStmtByExchangeId, err = dao.db.Prepare("SELECT id,stock_id,exchange_id,stock_code,stock_name_cn,stock_name_en,launch_date,company_name_cn,company_name_en,website_url,industry_name,city_name_cn,city_name_en,province_name_cn,province_name_en,create_time,update_time FROM stock WHERE exchange_id=?")
	if err != nil {
		return err
	}

	return nil
}

func (dao *StockDao) prepareSelectStmtByExchangeIdAndStockCode() (err error) {
	dao.selectStmtByExchangeIdAndStockCode, err = dao.db.Prepare("SELECT id,stock_id,exchange_id,stock_code,stock_name_cn,stock_name_en,launch_date,company_name_cn,company_name_en,website_url,industry_name,city_name_cn,city_name_en,province_name_cn,province_name_en,create_time,update_time FROM stock WHERE exchange_id=? AND stock_code=?")
	if err != nil {
		return err
	}

	return nil
}

func (dao *StockDao) Insert(tx *sql.Tx, e *Stock) (id int64, err error) {
	stmt := dao.insertStmt
	if tx != nil {
		stmt = tx.Stmt(stmt)
	}

	result, err := stmt.Exec(e.StockId, e.ExchangeId, e.StockCode, e.StockNameCn, e.StockNameEn, e.LaunchDate, e.CompanyNameCn, e.CompanyNameEn, e.WebsiteUrl, e.IndustryName, e.CityNameCn, e.CityNameEn, e.ProvinceNameCn, e.ProvinceNameEn, e.CreateTime, e.UpdateTime)
	if err != nil {
		return 0, err
	}

	id, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (dao *StockDao) Update(tx *sql.Tx, e *Stock) (int64, error) {
	stmt := dao.updateStmt
	if tx != nil {
		stmt = tx.Stmt(stmt)
	}

	result, err := stmt.Exec(e.StockId, e.ExchangeId, e.StockCode, e.StockNameCn, e.StockNameEn, e.LaunchDate, e.CompanyNameCn, e.CompanyNameEn, e.WebsiteUrl, e.IndustryName, e.CityNameCn, e.CityNameEn, e.ProvinceNameCn, e.ProvinceNameEn, e.CreateTime, e.UpdateTime, e.Id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func (dao *StockDao) Delete(tx *sql.Tx, id int64) (rowsAffected int64, err error) {
	stmt := dao.deleteStmt
	if tx != nil {
		stmt = tx.Stmt(stmt)
	}

	result, err := stmt.Exec(id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func (dao *StockDao) ScanRow(row *sql.Row) (*Stock, error) {
	e := &Stock{}
	err := row.Scan(&e.Id, &e.StockId, &e.ExchangeId, &e.StockCode, &e.StockNameCn, &e.StockNameEn, &e.LaunchDate, &e.CompanyNameCn, &e.CompanyNameEn, &e.WebsiteUrl, &e.IndustryName, &e.CityNameCn, &e.CityNameEn, &e.ProvinceNameCn, &e.ProvinceNameEn, &e.CreateTime, &e.UpdateTime)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return e, nil
}

func (dao *StockDao) ScanRows(rows *sql.Rows) (list []*Stock, err error) {
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
		return nil, rows.Err()
	}

	return list, nil
}

func (dao *StockDao) SelectAll(tx *sql.Tx) (list []*Stock, err error) {
	stmt := dao.selectStmtAll
	if tx != nil {
		stmt = tx.Stmt(stmt)
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	return dao.ScanRows(rows)
}

func (dao *StockDao) SelectById(tx *sql.Tx, Id int64) (*Stock, error) {
	stmt := dao.selectStmtById
	if tx != nil {
		stmt = tx.Stmt(stmt)
	}

	row := stmt.QueryRow(Id)

	return dao.ScanRow(row)
}

func (dao *StockDao) SelectByUpdateTime(tx *sql.Tx, UpdateTime string) (*Stock, error) {
	stmt := dao.selectStmtByUpdateTime
	if tx != nil {
		stmt = tx.Stmt(stmt)
	}

	row := stmt.QueryRow(UpdateTime)

	return dao.ScanRow(row)
}

func (dao *StockDao) SelectListByUpdateTime(tx *sql.Tx, UpdateTime string) (list []*Stock, err error) {
	stmt := dao.selectStmtByUpdateTime
	if tx != nil {
		stmt = tx.Stmt(stmt)
	}

	rows, err := stmt.Query(UpdateTime)
	if err != nil {
		return nil, err
	}

	return dao.ScanRows(rows)
}

func (dao *StockDao) SelectByIndustryName(tx *sql.Tx, IndustryName string) (*Stock, error) {
	stmt := dao.selectStmtByIndustryName
	if tx != nil {
		stmt = tx.Stmt(stmt)
	}

	row := stmt.QueryRow(IndustryName)

	return dao.ScanRow(row)
}

func (dao *StockDao) SelectListByIndustryName(tx *sql.Tx, IndustryName string) (list []*Stock, err error) {
	stmt := dao.selectStmtByIndustryName
	if tx != nil {
		stmt = tx.Stmt(stmt)
	}

	rows, err := stmt.Query(IndustryName)
	if err != nil {
		return nil, err
	}

	return dao.ScanRows(rows)
}

func (dao *StockDao) SelectByStockId(tx *sql.Tx, StockId string) (*Stock, error) {
	stmt := dao.selectStmtByStockId
	if tx != nil {
		stmt = tx.Stmt(stmt)
	}

	row := stmt.QueryRow(StockId)

	return dao.ScanRow(row)
}

func (dao *StockDao) SelectByExchangeIdAndStockCode(tx *sql.Tx, ExchangeId string, StockCode string) (*Stock, error) {
	stmt := dao.selectStmtByExchangeIdAndStockCode
	if tx != nil {
		stmt = tx.Stmt(stmt)
	}

	row := stmt.QueryRow(ExchangeId, StockCode)

	return dao.ScanRow(row)
}

func (dao *StockDao) SelectListByExchangeId(tx *sql.Tx, ExchangeId string) (list []*Stock, err error) {
	stmt := dao.selectStmtByExchangeId
	if tx != nil {
		stmt = tx.Stmt(stmt)
	}

	rows, err := stmt.Query(ExchangeId)
	if err != nil {
		return nil, err
	}

	return dao.ScanRows(rows)
}
