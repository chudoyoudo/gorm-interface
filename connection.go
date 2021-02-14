package gorm_interface

import (
    "database/sql"

    "github.com/golobby/container"
    "gorm.io/gorm"
)

var db *gorm.DB

func setDb(g *gorm.DB) {
    db = g
}

func getDb() *gorm.DB {
    if db == nil {
        container.Make(&db)
    }
    return db
}

type Connection interface {
    Create(value interface{}) Connection
    Save(value interface{}) Connection
    Updates(values interface{}) Connection
    First(dest interface{}, conds ...interface{}) Connection
    Last(dest interface{}, conds ...interface{}) Connection
    Find(dest interface{}, conds ...interface{}) Connection
    Order(value interface{}) Connection
    Limit(limit int) Connection
    Offset(offset int) Connection
    Delete(value interface{}, conds ...interface{}) Connection
    Count(count *int64) Connection
    Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error
    Begin(opts ...*sql.TxOptions) Connection
    Commit() Connection
    Rollback() Connection
    Exec(sql string, values ...interface{}) Connection
    Error() error
    RowsAffected() int64
}

func NewConnection() Connection {
    db := getDb()
    return &connection{db: db}
}

type connection struct {
    db *gorm.DB
}

func (c *connection) Create(value interface{}) Connection {
    db := c.db.Create(value)
    c.db = db
    return c
}

func (c *connection) Save(value interface{}) Connection {
    db := c.db.Save(value)
    c.db = db
    return c
}

func (c *connection) Updates(values interface{}) Connection {
    db := c.db.Updates(values)
    c.db = db
    return c
}

func (c *connection) First(dest interface{}, conds ...interface{}) Connection {
    db := c.db.First(dest, conds...)
    c.db = db
    return c
}

func (c *connection) Last(dest interface{}, conds ...interface{}) Connection {
    db := c.db.Last(dest, conds...)
    c.db = db
    return c
}

func (c *connection) Find(dest interface{}, conds ...interface{}) Connection {
    db := c.db.Find(dest, conds...)
    c.db = db
    return c
}

func (c *connection) Order(value interface{}) Connection {
    db := c.db.Order(value)
    c.db = db
    return c
}

func (c *connection) Limit(limit int) Connection {
    db := c.db.Limit(limit)
    c.db = db
    return c
}

func (c *connection) Offset(offset int) Connection {
    db := c.db.Offset(offset)
    c.db = db
    return c
}

func (c *connection) Delete(value interface{}, conds ...interface{}) Connection {
    db := c.db.Find(value, conds...)
    c.db = db
    return c
}

func (c *connection) Count(count *int64) Connection {
    db := c.db.Count(count)
    c.db = db
    return c
}

func (c *connection) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
    return c.db.Transaction(fc, opts...)
}

func (c *connection) Begin(opts ...*sql.TxOptions) Connection {
    db := c.db.Begin(opts...)
    c.db = db
    return c
}

func (c *connection) Commit() Connection {
    db := c.db.Begin()
    c.db = db
    return c
}

func (c *connection) Rollback() Connection {
    db := c.db.Rollback()
    c.db = db
    return c
}

func (c *connection) Exec(sql string, values ...interface{}) Connection {
    db := c.db.Exec(sql, values...)
    c.db = db
    return c
}

func (c *connection) Error() error {
    return c.db.Error
}

func (c *connection) RowsAffected() int64 {
    return c.db.RowsAffected
}
