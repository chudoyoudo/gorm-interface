package gorm_interface

import (
    "database/sql"

    "github.com/stretchr/testify/mock"
    "gorm.io/gorm"
)

type ConnectionMock struct {
    mock.Mock
    Err     error
    RowsAff int64
}

func (c *ConnectionMock) Create(value interface{}) Connection {
    args := c.Called(value)
    return args.Get(0).(Connection)
}

func (c *ConnectionMock) Save(value interface{}) Connection {
    args := c.Called(value)
    return args.Get(0).(Connection)
}

func (c *ConnectionMock) First(dest interface{}, conds ...interface{}) Connection {
    args := c.Called(dest, conds)
    return args.Get(0).(Connection)
}

func (c *ConnectionMock) Last(dest interface{}, conds ...interface{}) Connection {
    args := c.Called(dest, conds)
    return args.Get(0).(Connection)
}

func (c *ConnectionMock) Order(value interface{}) Connection {
    args := c.Called(value)
    return args.Get(0).(Connection)
}

func (c *ConnectionMock) Find(dest interface{}, conds ...interface{}) Connection {
    args := c.Called(dest, conds)
    return args.Get(0).(Connection)
}

func (c *ConnectionMock) Limit(limit int) Connection {
    args := c.Called(limit)
    return args.Get(0).(Connection)
}

func (c *ConnectionMock) Offset(offset int) Connection {
    args := c.Called(offset)
    return args.Get(0).(Connection)
}

func (c *ConnectionMock) Delete(value interface{}, conds ...interface{}) Connection {
    args := c.Called(value, conds)
    return args.Get(0).(Connection)
}

func (c *ConnectionMock) Count(count *int64) Connection {
    args := c.Called(count)
    return args.Get(0).(Connection)
}

func (c *ConnectionMock) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
    args := c.Called(fc, opts)
    return args.Error(0)
}

func (c *ConnectionMock) Begin(opts ...*sql.TxOptions) Connection {
    args := c.Called(opts)
    return args.Get(0).(Connection)
}

func (c *ConnectionMock) Commit() Connection {
    args := c.Called()
    return args.Get(0).(Connection)
}

func (c *ConnectionMock) Rollback() Connection {
    args := c.Called()
    return args.Get(0).(Connection)
}

func (c *ConnectionMock) Exec(sql string, values ...interface{}) Connection {
    args := c.Called(sql, values)
    return args.Get(0).(Connection)
}

func (c *ConnectionMock) Error() error {
    return c.Err
}

func (c *ConnectionMock) RowsAffected() int64 {
    return c.RowsAff
}
