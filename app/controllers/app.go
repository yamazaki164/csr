package controllers

import (
	"csr/app/services/db"
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
	dbm *db.Manager
}

func (this *App) Dbm() *db.Manager {
	if this.dbm == nil {
		this.dbm = db.NewManager()
	}

	return this.dbm
}

func (this *App) Close() revel.Result {
	if this.dbm != nil {
		this.dbm.Close()
	}

	return nil
}

func init() {
	revel.InterceptMethod((*App).Close, revel.AFTER)
}
