package gorm_interface

import (
    "database/sql"

    "github.com/stretchr/testify/mock"
    "gorm.io/gorm"
)

type Connection interface {
    Create(value interface{}) (tx *gorm.DB)
    CreateInBatches(value interface{}, batchSize int) (tx *gorm.DB)
    Save(value interface{}) (tx *gorm.DB)
    First(dest interface{}, conds ...interface{}) (tx *gorm.DB)
    Take(dest interface{}, conds ...interface{}) (tx *gorm.DB)
    Last(dest interface{}, conds ...interface{}) (tx *gorm.DB)
    Find(dest interface{}, conds ...interface{}) (tx *gorm.DB)
    Limit(limit int) (tx *gorm.DB)
    Offset(offset int) (tx *gorm.DB)
    FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB
    FirstOrInit(dest interface{}, conds ...interface{}) (tx *gorm.DB)
    FirstOrCreate(dest interface{}, conds ...interface{}) (tx *gorm.DB)
    Update(column string, value interface{}) (tx *gorm.DB)
    Updates(values interface{}) (tx *gorm.DB)
    UpdateColumn(column string, value interface{}) (tx *gorm.DB)
    UpdateColumns(values interface{}) (tx *gorm.DB)
    Delete(value interface{}, conds ...interface{}) (tx *gorm.DB)
    Count(count *int64) (tx *gorm.DB)
    Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error)
    Begin(opts ...*sql.TxOptions) *gorm.DB
    Commit() *gorm.DB
    Rollback() *gorm.DB
    Exec(sql string, values ...interface{}) (tx *gorm.DB)
}

type ConnectionMock struct {
    mock.Mock
}

func (db *ConnectionMock) Create(value interface{}) (tx *gorm.DB) {
    args := db.Called(value)
    return args.Get(0).(*gorm.DB)
}

func (db *ConnectionMock) CreateInBatches(value interface{}, batchSize int) (tx *gorm.DB) {
    args := db.Called(value, batchSize)
    return args.Get(0).(*gorm.DB)
}

func (db *ConnectionMock) Save(value interface{}) (tx *gorm.DB) {
    args := db.Called(value)
    return args.Get(0).(*gorm.DB)
}

func (db *ConnectionMock) First(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
    args := db.Called(dest, conds)
    return args.Get(0).(*gorm.DB)
}

func (db *ConnectionMock) Take(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
    args := db.Called(dest, conds)
    return args.Get(0).(*gorm.DB)
}

func (db *ConnectionMock) Last(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
    args := db.Called(dest, conds)
    return args.Get(0).(*gorm.DB)
}

func (db *ConnectionMock) Find(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
    args := db.Called(dest, conds)
    return args.Get(0).(*gorm.DB)
}

func (db *ConnectionMock) Limit(limit int) (tx *gorm.DB) {
    args := db.Called(limit)
    return args.Get(0).(*gorm.DB)
}

func (db *ConnectionMock) Offset(offset int) (tx *gorm.DB) {
    args := db.Called(offset)
    return args.Get(0).(*gorm.DB)
}

func (db *ConnectionMock) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
    args := db.Called(dest, batchSize, fc)
    return args.Get(0).(*gorm.DB)
}

func (db *ConnectionMock) FirstOrInit(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
    args := db.Called(dest, conds)
    return args.Get(0).(*gorm.DB)
}

func (db *ConnectionMock) FirstOrCreate(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
    args := db.Called(dest, conds)
    return args.Get(0).(*gorm.DB)
}

func (db *ConnectionMock) Update(column string, value interface{}) (tx *gorm.DB) {
    args := db.Called(column, value)
    return args.Get(0).(*gorm.DB)
}

func (db *ConnectionMock) Updates(values interface{}) (tx *gorm.DB) {
    args := db.Called(values)
    return args.Get(0).(*gorm.DB)
}

func (db *ConnectionMock) UpdateColumn(column string, value interface{}) (tx *gorm.DB) {
    args := db.Called(column, value)
    return args.Get(0).(*gorm.DB)
}

func (db *ConnectionMock) UpdateColumns(values interface{}) (tx *gorm.DB) {
    args := db.Called(values)
    return args.Get(0).(*gorm.DB)
}

func (db *ConnectionMock) Delete(value interface{}, conds ...interface{}) (tx *gorm.DB) {
    args := db.Called(value, conds)
    return args.Get(0).(*gorm.DB)
}

func (db *ConnectionMock) Count(count *int64) (tx *gorm.DB) {
    args := db.Called(count)
    return args.Get(0).(*gorm.DB)
}

func (db *ConnectionMock) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
    args := db.Called(fc, opts)
    return args.Error(0)
}

func (db *ConnectionMock) Begin(opts ...*sql.TxOptions) *gorm.DB {
    args := db.Called(opts)
    return args.Get(0).(*gorm.DB)
}

func (db *ConnectionMock) Commit() *gorm.DB {
    args := db.Called()
    return args.Get(0).(*gorm.DB)
}

func (db *ConnectionMock) Rollback() *gorm.DB {
    args := db.Called()
    return args.Get(0).(*gorm.DB)
}

func (db *ConnectionMock) Exec(sql string, values ...interface{}) (tx *gorm.DB) {
    args := db.Called(sql, values)
    return args.Get(0).(*gorm.DB)
}
