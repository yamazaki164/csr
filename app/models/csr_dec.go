package models

import (
	"crypto/x509"
	"encoding/base64"
	"github.com/revel/revel"
	"strings"
)

type CsrDec struct {
	CsrString string
	Csr       []byte
	ParsedCsr *x509.CertificateRequest
	IsError   bool
}

func NewCsrDec(csr string) *CsrDec {
	dec := &CsrDec{
		CsrString: csr,
	}

	return dec
}

func (this *CsrDec) Trim() {
	this.CsrString = strings.TrimPrefix(this.CsrString, CSR_PREFIX)
	this.CsrString = strings.TrimSuffix(this.CsrString, CSR_SUFFIX)
}

func (this *CsrDec) ToDecode() {
	var err error
	this.Csr, err = base64.StdEncoding.DecodeString(this.CsrString)
	if err != nil {
		this.IsError = true
		revel.ERROR.Println(err)
		return
	}

	this.ParsedCsr, err = x509.ParseCertificateRequest(this.Csr)
	if err != nil {
		this.IsError = true
		revel.ERROR.Println(err)
		return
	}

	this.IsError = false
}

func (this *CsrDec) Parse() {
	this.Trim()
	this.ToDecode()
}
