package db

import (
	"csr/app/services/db"
	"github.com/revel/revel/testing"
)

type ManagerTest struct {
	testing.TestSuite
}

func (t *ManagerTest) Test_NewManager() {
	m := db.NewManager()

	t.Assert(m != nil)
	t.Assert(m.DB != nil)

	m.Close()
	t.Assert(m.DB == nil)
}

func (t *ManagerTest) Test_Connection() {
	m := db.NewManager()
	c := m.Connection()

	t.Assert(c != nil)

	t.Assert(m.Txn == nil)
	m.Begin()
	t.Assert(m.Txn != nil)

	m.Close()
	t.Assert(m.Txn == nil)
	t.Assert(m.DB == nil)
}

func (t *ManagerTest) Test_Begin() {
	m := db.NewManager()

	t.Assert(m.Txn == nil)
	m.Begin()
	t.Assert(m.Txn != nil)

	m.Rollback()
	t.Assert(m.Txn == nil)

	m.Close()
}

func (t *ManagerTest) Test_Commit() {
	m := db.NewManager()

	t.Assert(m.Txn == nil)
	ret := m.Commit()
	t.Assert(ret != nil)

	t.Assert(m.Txn == nil)
	m.Begin()
	t.Assert(m.Txn != nil)
	ret = m.Commit()
	t.Assert(ret == nil)
	t.Assert(m.Txn == nil)

	m.Close()
}

func (t *ManagerTest) Test_Rollback() {
	m := db.NewManager()

	t.Assert(m.Txn == nil)
	ret := m.Rollback()
	t.Assert(ret != nil)

	m.Begin()
	t.Assert(m.Txn != nil)

	ret = m.Rollback()
	t.Assert(ret == nil)
	t.Assert(m.Txn == nil)

	m.Close()
}

func (t *ManagerTest) Test_TransactWith() {
	m := db.NewManager()

	ret := m.TransactWith(func() error {
		t.Assert(m.Txn != nil)

		return nil
	})
	t.Assert(ret == nil)
	t.Assert(m.Txn == nil)

	m.Close()
}
