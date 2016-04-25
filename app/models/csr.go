package models

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"github.com/revel/revel"
	"regexp"
	"strings"
	"time"
)

const (
	PRIVATE_KEY_PREFIX = "-----BEGIN RSA PRIVATE KEY-----"
	PRIVATE_KEY_SUFFIX = "-----END RSA PRIVATE KEY-----"

	CSR_PREFIX = "-----BEGIN CERTIFICATE REQUEST-----"
	CSR_SUFFIX = "-----END CERTIFICATE REQUEST-----"
)

type Csr struct {
	ID                 uint `gorm:"primary_key"`
	KeyBits            int
	CsrAlgorithm       string
	Country            string
	State              string
	Locality           string
	OrganizationalName string
	OrganizationalUnit string
	CommonName         string
	privRsa            *rsa.PrivateKey `sql:"-"`
	priv               []byte          `sql:"-"`
	csr                []byte          `sql:"-"`
	PrivString         string          `sql:"type:text"`
	CsrString          string          `sql:"type:text"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          *time.Time
}

func (this *Csr) SetValidations(v *revel.Validation) {
	v.Required(this.KeyBits).Message("KeyBits required")
	v.Required(this.KeyBits == 2048).Message("KeyBits not match")

	v.Required(this.CsrAlgorithm).Message("CsrAlgorithm required")
	v.Match(this.CsrAlgorithm, regexp.MustCompile("^sha(1|256)$")).Message("CsrAlgorithm not match")

	v.Required(this.Country).Message("Country required")
	v.Match(this.Country, regexp.MustCompile("^[A-Z]{2}$")).Message("Country not match")

	v.Required(this.State).Message("State required")
	v.Match(this.State, regexp.MustCompile("^[a-zA-Z\\-_]+$")).Message("State not match")

	v.Required(this.Locality).Message("Locality required")
	v.Match(this.Locality, regexp.MustCompile("^[a-zA-Z\\-_]+$")).Message("Locality not match")

	v.Required(this.OrganizationalName).Message("OrganizationalName required")
	v.Match(this.OrganizationalName, regexp.MustCompile("^[a-zA-Z0-9\\-_,\\. ()]+$")).Message("OrganizationalName not match")

	if len(this.OrganizationalUnit) > 0 {
		v.Match(this.OrganizationalUnit, regexp.MustCompile("^[a-zA-Z0-9\\-_,\\. ]+$")).Message("OrganizationalUnit not match")
	}

	v.Required(this.CommonName).Message("CommonName required")
	v.Match(this.CommonName, regexp.MustCompile("^[a-z0-9\\.\\-\\*]+$")).Message("CommonName not match")
}

func (this *Csr) keyTo(data []byte) string {
	ret := base64.StdEncoding.EncodeToString(data)

	var seplen int = 64
	var s []string = []string{}
	var str string = ""
	for i := 0; i < len(ret); i += seplen {
		str = ""
		if i+seplen < len(ret) {
			str = ret[i:(i + seplen)]
			s = append(s, str)
		} else {
			str = ret[i:]
			s = append(s, str)
		}
	}

	return strings.Join(s, "\n")
}

func (this *Csr) ToPrivateKey() {
	this.privRsa, _ = rsa.GenerateKey(rand.Reader, this.KeyBits)
	this.priv = x509.MarshalPKCS1PrivateKey(this.privRsa)

	this.PrivString = PRIVATE_KEY_PREFIX + "\n" + this.keyTo(this.priv) + "\n" + PRIVATE_KEY_SUFFIX
}

func (this *Csr) ToCsr() {
	data := &x509.CertificateRequest{}
	data.Subject.Country = []string{
		this.Country,
	}
	data.Subject.Province = []string{
		this.State,
	}
	data.Subject.Locality = []string{
		this.Locality,
	}
	if len(this.OrganizationalName) > 0 {
		data.Subject.Organization = []string{
			this.OrganizationalName,
		}
	}
	if len(this.OrganizationalUnit) > 0 {
		data.Subject.OrganizationalUnit = []string{
			this.OrganizationalUnit,
		}
	}
	data.Subject.CommonName = this.CommonName

	if this.CsrAlgorithm == "sha256" {
		data.SignatureAlgorithm = x509.SHA256WithRSA
	} else {
		data.SignatureAlgorithm = x509.SHA1WithRSA
	}

	data.Signature = this.priv

	this.csr, _ = x509.CreateCertificateRequest(rand.Reader, data, this.privRsa)

	this.CsrString = CSR_PREFIX + "\n" + this.keyTo(this.csr) + "\n" + CSR_SUFFIX
}

func (this *Csr) BeforeCreate() {
	this.ToPrivateKey()
	this.ToCsr()
}
