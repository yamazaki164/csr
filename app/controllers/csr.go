package controllers

import (
	"csr/app/models"
	"csr/app/routes"
	"github.com/revel/revel"
)

type Csr struct {
	App
}

func (this *Csr) Index() revel.Result {
	//	this.Dbm().Connection().DropTable(&models.Csr{})
	//	this.Dbm().Connection().CreateTable(&models.Csr{})
	var csrs []models.Csr = nil
	this.Dbm().Connection().Find(&csrs)

	return this.Render(csrs)
}

func (this *Csr) New() revel.Result {
	return this.Render()
}

func (this *Csr) Create(csr *models.Csr) revel.Result {
	csr.SetValidations(this.Validation)
	if this.Validation.HasErrors() {
		this.Validation.Keep()
		this.FlashParams()

		this.Flash.Error("Create error")
		return this.Redirect(routes.Csr.New())
	}

	this.Dbm().Connection().Save(csr)

	this.Flash.Success("Create success")
	return this.Redirect(routes.Csr.Show(csr.ID))
}

func (this *Csr) Show(id uint) revel.Result {
	var csr models.Csr

	this.Dbm().Connection().Where("id = ?", id).First(&csr)
	if csr.ID == 0 {
		return this.NotFound("not found")
	}

	return this.Render(csr)
}

func (this *Csr) Delete(id uint) revel.Result {
	var csr models.Csr
	this.Dbm().Connection().Where("id = ?", id).First(&csr)
	if csr.ID == 0 {
		this.Flash.Error("Delete error")
		return this.Redirect(routes.Csr.Index())
	}

	this.Dbm().Connection().Delete(&csr)

	this.Flash.Success("Delete success")
	return this.Redirect(routes.Csr.Index())
}
