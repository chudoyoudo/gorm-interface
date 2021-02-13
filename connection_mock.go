package gorm_interface

import (
    "database/sql"

    "github.com/stretchr/testify/mock"
    "gorm.io/gorm"
)

type ConnectionMock struct {
    mock.Mock
}

func (db *ConnectionMock) Create(value interface{}) Connection {
    args := db.Called(value)
    return args.Get(0).(Connection)
}

func (db *ConnectionMock) Save(value interface{}) Connection {
    args := db.Called(value)
    return args.Get(0).(Connection)
}

func (db *ConnectionMock) First(dest interface{}, conds ...interface{}) Connection {
    args := db.Called(dest, conds)
    return args.Get(0).(Connection)
}

func (db *ConnectionMock) Last(dest interface{}, conds ...interface{}) Connection {
    args := db.Called(dest, conds)
    return args.Get(0).(Connection)
}

func (db *ConnectionMock) Find(dest interface{}, conds ...interface{}) Connection {
    args := db.Called(dest, conds)
    return args.Get(0).(Connection)
}

func (db *ConnectionMock) Limit(limit int) Connection {
    args := db.Called(limit)
    return args.Get(0).(Connection)
}

func (db *ConnectionMock) Offset(offset int) Connection {
    args := db.Called(offset)
    return args.Get(0).(Connection)
}

func (db *ConnectionMock) Delete(value interface{}, conds ...interface{}) Connection {
    args := db.Called(value, conds)
    return args.Get(0).(Connection)
}

func (db *ConnectionMock) Count(count *int64) Connection {
    args := db.Called(count)
    return args.Get(0).(Connection)
}

func (db *ConnectionMock) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
    args := db.Called(fc, opts)
    return args.Error(0)
}

func (db *ConnectionMock) Begin(opts ...*sql.TxOptions) Connection {
    args := db.Called(opts)
    return args.Get(0).(Connection)
}

func (db *ConnectionMock) Commit() Connection {
    args := db.Called()
    return args.Get(0).(Connection)
}

func (db *ConnectionMock) Rollback() Connection {
    args := db.Called()
    return args.Get(0).(Connection)
}

func (db *ConnectionMock) Exec(sql string, values ...interface{}) Connection {
    args := db.Called(sql, values)
    return args.Get(0).(Connection)
}

func (db *ConnectionMock) Error() error {
    args := db.Called()
    return args.Error(0)
}

func (db *ConnectionMock) RowsAffected() int64 {
    args := db.Called()
    return int64(args.Int(0))
}
