package gorm_interface

import (
    "database/sql"

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
