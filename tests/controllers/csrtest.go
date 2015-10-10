package controllers

import (
	"csr/app/models"
	"csr/app/services/db"
	"github.com/revel/revel/testing"
	"net/url"
)

type CsrTest struct {
	testing.TestSuite
}

func (t *CsrTest) Before() {
	db := db.NewManager()

	db.TransactWith(func() error {
		for i := 0; i < 20; i++ {
			data := &models.Csr{
				KeyBits:            2048,
				CsrAlgorithm:       "sha256",
				Country:            "JP",
				State:              "Tokyo",
				Locality:           "Piyo" + string(i),
				OrganizationalName: "fuga" + string(i),
				OrganizationalUnit: "",
				CommonName:         "*.test.com",
			}

			db.Connection().Save(&data)
		}

		return nil
	})
}

func (t *CsrTest) After() {
	db := db.NewManager()
	db.Connection().Exec("delete from csrs;")
	db.Connection().Exec("vaccum;")
}

func (t *CsrTest) Test_Index() {
	t.Get("/")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
	t.AssertContains("Delete")
}

func (t *CsrTest) Test_New() {
	t.Get("/new")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

func (t *CsrTest) Test_Create_Fail() {
	data := make(url.Values)
	t.PostForm("/create", data)
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
	t.AssertContains("error")
}

func (t *CsrTest) Test_Create_Success() {
	data := make(url.Values)
	data.Add("csr.KeyBits", "2048")
	data.Add("csr.CsrAlgorithm", "sha256")
	data.Add("csr.Country", "JP")
	data.Add("csr.State", "Tokyo")
	data.Add("csr.Locality", "Piyo")
	data.Add("csr.OrganizationalName", "fuga")
	data.Add("csr.OrganizationalUnit", "")
	data.Add("csr.CommonName", "*.test.com")

	t.PostForm("/create", data)
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
	t.AssertContains("success")
}

func (t *CsrTest) Test_Show_Fail() {
	t.Get("/0")
	t.AssertNotFound()
	t.AssertContentType("text/html; charset=utf-8")
}

func (t *CsrTest) Test_Show_Success() {
	t.Get("/1")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
	t.AssertContains("<th>Server Key</th>")
}

func (t *CsrTest) Test_Delete_Fail() {
	data := make(url.Values)
	t.PostForm("/0/delete", data)
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
	t.AssertContains("error")
}

func (t *CsrTest) Test_Delete_Success() {
	data := make(url.Values)
	t.PostForm("/1/delete", data)
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
	t.AssertContains("success")
}

func (t *CsrTest) Test_Dec_Get() {
	t.Get("/dec")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}
