package db

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/revel/revel"
	"path"
)

type Manager struct {
	DB  *gorm.DB
	Txn *gorm.DB
}

func NewManager() *Manager {
	var found bool
	var driver string
	if driver, found = revel.Config.String("db.driver"); !found {
		revel.ERROR.Fatal("db driver not found")
	}

	var spec string
	if spec, found = revel.Config.String("db.spec"); !found {
		revel.ERROR.Fatal("db spec not found.")
	}

	if driver == "sqlite3" {
		spec = path.Join(revel.AppPath, spec)
	}

	var err error
	var db gorm.DB
	db, err = gorm.Open(driver, spec)
	if err != nil {
		revel.ERROR.Fatal(err)
	}

	db.LogMode(revel.Config.BoolDefault("db.debug", false))

	m := &Manager{
		DB: &db,
	}

	return m
}

func (this *Manager) Connection() *gorm.DB {
	if this.Txn != nil {
		return this.Txn
	} else {
		return this.DB
	}
}

func (this *Manager) TransactWith(callback func() error) error {
	defer func() {
		if r := recover(); r != nil {
			this.Rollback()
		} else {
			this.Commit()
		}
	}()

	this.Begin()
	return callback()
}

func (this *Manager) Begin() {
	this.Txn = this.DB.Begin()
	if this.Txn.Error != nil {
		panic(this.Txn.Error)
	}
}

func (this *Manager) Rollback() error {
	if this.Txn == nil {
		return errors.New("transaction not used")
	}

	this.Txn.Rollback()
	if this.Txn.Error != nil && this.Txn.Error != sql.ErrTxDone {
		return this.Txn.Error
	}
	this.Txn = nil
	return nil
}

func (this *Manager) Commit() error {
	if this.Txn == nil {
		return errors.New("transaction not used")
	}

	this.Txn.Commit()
	if this.Txn.Error != nil && this.Txn.Error != sql.ErrTxDone {
		return this.Txn.Error
	}
	this.Txn = nil
	return nil
}

func (this *Manager) Close() {
	this.DB.Close()
	this.DB = nil
	this.Txn = nil
}
