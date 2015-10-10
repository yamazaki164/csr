package models

import (
	"csr/app/models"
	"github.com/revel/revel"
	"github.com/revel/revel/testing"
)

type CsrTest struct {
	testing.TestSuite
}

func (t *CsrTest) Test_SetValidations_KeyBits() {
	csr := models.Csr{
		KeyBits:            0,
		CsrAlgorithm:       "sha256",
		Country:            "JP",
		State:              "Tokyo",
		Locality:           "Piyo",
		OrganizationalName: "fuga",
		OrganizationalUnit: "",
		CommonName:         "*.test.com",
	}
	v := &revel.Validation{}

	v.Clear()
	csr.KeyBits = 0
	csr.SetValidations(v)
	t.Assert(v.HasErrors() == true)

	v.Clear()
	csr.KeyBits = 2048
	csr.SetValidations(v)
	t.Assert(v.HasErrors() == false)
}

func (t *CsrTest) Test_SetValidations_CsrAlgorithm() {
	csr := models.Csr{
		KeyBits:            2048,
		CsrAlgorithm:       "sha512",
		Country:            "JP",
		State:              "Tokyo",
		Locality:           "Piyo",
		OrganizationalName: "fuga",
		OrganizationalUnit: "",
		CommonName:         "*.test.com",
	}
	v := &revel.Validation{}

	v.Clear()
	csr.CsrAlgorithm = ""
	csr.SetValidations(v)
	t.Assert(v.HasErrors() == true)

	v.Clear()
	csr.CsrAlgorithm = "sha1"
	csr.SetValidations(v)
	t.Assert(v.HasErrors() == false)

	v.Clear()
	csr.CsrAlgorithm = "sha256"
	csr.SetValidations(v)
	t.Assert(v.HasErrors() == false)

	v.Clear()
	csr.CsrAlgorithm = "sha"
	csr.SetValidations(v)
	t.Assert(v.HasErrors() == true)
}

func (t *CsrTest) Test_SetValidations_Country() {
	csr := models.Csr{
		KeyBits:            2048,
		CsrAlgorithm:       "sha256",
		Country:            "",
		State:              "Tokyo",
		Locality:           "Piyo",
		OrganizationalName: "fuga",
		OrganizationalUnit: "",
		CommonName:         "*.test.com",
	}
	v := &revel.Validation{}

	v.Clear()
	csr.Country = ""
	csr.SetValidations(v)
	t.Assert(v.HasErrors() == true)

	v.Clear()
	csr.Country = "jp"
	csr.SetValidations(v)
	t.Assert(v.HasErrors() == true)

	v.Clear()
	csr.Country = "J"
	csr.SetValidations(v)
	t.Assert(v.HasErrors() == true)

	v.Clear()
	csr.Country = "JP"
	csr.SetValidations(v)
	t.Assert(v.HasErrors() == false)
}

func (t *CsrTest) Test_SetValidations_State() {
	csr := models.Csr{
		KeyBits:            2048,
		CsrAlgorithm:       "sha256",
		Country:            "JP",
		State:              "",
		Locality:           "Piyo",
		OrganizationalName: "fuga",
		OrganizationalUnit: "",
		CommonName:         "*.test.com",
	}
	v := &revel.Validation{}

	v.Clear()
	csr.State = ""
	csr.SetValidations(v)
	t.Assert(v.HasErrors() == true)

	v.Clear()
	csr.State = "Tokyo"
	csr.SetValidations(v)
	t.Assert(v.HasErrors() == false)

	v.Clear()
	csr.State = "tokyo"
	csr.SetValidations(v)
	t.Assert(v.HasErrors() == false)
}

func (t *CsrTest) Test_SetValidations_Locality() {
	csr := models.Csr{
		KeyBits:            2048,
		CsrAlgorithm:       "sha256",
		Country:            "JP",
		State:              "tokyo",
		Locality:           "",
		OrganizationalName: "fuga",
		OrganizationalUnit: "",
		CommonName:         "*.test.com",
	}
	v := &revel.Validation{}

	v.Clear()
	csr.Locality = ""
	csr.SetValidations(v)
	t.Assert(v.HasErrors() == true)

	csr.Locality = "meguro-ku"
	v.Clear()
	csr.SetValidations(v)
	t.Assert(v.HasErrors() == false)
}

func (t *CsrTest) Test_SetValidations_OrganizationalName() {
	csr := models.Csr{
		KeyBits:            2048,
		CsrAlgorithm:       "sha256",
		Country:            "JP",
		State:              "tokyo",
		Locality:           "Meguro-ku",
		OrganizationalName: "",
		OrganizationalUnit: "",
		CommonName:         "*.test.com",
	}
	v := &revel.Validation{}

	v.Clear()
	csr.OrganizationalName = ""
	csr.SetValidations(v)
	t.Assert(v.HasErrors() == true)

	csr.OrganizationalName = "fuga,.ltd."
	v.Clear()
	csr.SetValidations(v)
	t.Assert(v.HasErrors() == false)
}

func (t *CsrTest) Test_SetValidations_OrganizationalUnit() {
	csr := models.Csr{
		KeyBits:            2048,
		CsrAlgorithm:       "sha256",
		Country:            "JP",
		State:              "tokyo",
		Locality:           "Meguro-ku",
		OrganizationalName: "fuga.,ltd.",
		OrganizationalUnit: "",
		CommonName:         "*.test.com",
	}
	v := &revel.Validation{}

	v.Clear()
	csr.OrganizationalUnit = ""
	csr.SetValidations(v)
	t.Assert(v.HasErrors() == false)

	csr.OrganizationalUnit = "fuga.,ltd."
	v.Clear()
	csr.SetValidations(v)
	t.Assert(v.HasErrors() == false)
}

func (t *CsrTest) Test_SetValidations_CommonName() {
	csr := models.Csr{
		KeyBits:            2048,
		CsrAlgorithm:       "sha256",
		Country:            "JP",
		State:              "tokyo",
		Locality:           "Meguro-ku",
		OrganizationalName: "fuga.,ltd.",
		OrganizationalUnit: "",
		CommonName:         "",
	}
	v := &revel.Validation{}

	v.Clear()
	csr.CommonName = ""
	csr.SetValidations(v)
	t.Assert(v.HasErrors() == true)

	v.Clear()
	csr.CommonName = "www.fuga_daf.com"
	csr.SetValidations(v)
	t.Assert(v.HasErrors() == true)

	v.Clear()
	csr.CommonName = "www.fuga.com"
	csr.SetValidations(v)
	t.Assert(v.HasErrors() == false)

	v.Clear()
	csr.CommonName = "*.fuga.com"
	csr.SetValidations(v)
	t.Assert(v.HasErrors() == false)
}

func (t *CsrTest) Test_ToPrivateKey() {
	csr := &models.Csr{
		KeyBits:            2048,
		CsrAlgorithm:       "sha256",
		Country:            "JP",
		State:              "Tokyo",
		Locality:           "Piyo",
		OrganizationalName: "fuga",
		OrganizationalUnit: "",
		CommonName:         "*.test.com",
	}

	csr.ToPrivateKey()
	t.Assert(csr.PrivString != "")
}

func (t *CsrTest) Test_ToCsr() {
	csr := &models.Csr{
		KeyBits:            2048,
		CsrAlgorithm:       "sha256",
		Country:            "JP",
		State:              "Tokyo",
		Locality:           "Piyo",
		OrganizationalName: "fuga",
		OrganizationalUnit: "",
		CommonName:         "*.test.com",
	}

	csr.ToPrivateKey()
	csr.ToCsr()

	t.Assert(csr.CsrString != "")
}
