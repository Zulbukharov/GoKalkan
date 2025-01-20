package gokalkan

import (
	"github.com/gokalkan/gokalkan/ckalkan"
)

func (cli *Client) LoadCertFromBuffer(cert []byte, certType ckalkan.CertCodeType) (err error) {
	return cli.kc.X509LoadCertificateFromBuffer(cert, certType)
}
