package cdn

import (
	"github.com/morlay/goaliyun"
)

type SetDomainServerCertificateRequest struct {
	PrivateKey              string `position:"Query" name:"PrivateKey"`
	ServerCertificateStatus string `position:"Query" name:"ServerCertificateStatus"`
	ServerCertificate       string `position:"Query" name:"ServerCertificate"`
	SecurityToken           string `position:"Query" name:"SecurityToken"`
	CertType                string `position:"Query" name:"CertType"`
	CertName                string `position:"Query" name:"CertName"`
	DomainName              string `position:"Query" name:"DomainName"`
	OwnerId                 int64  `position:"Query" name:"OwnerId"`
	Region                  string `position:"Query" name:"Region"`
	RegionId                string `position:"Query" name:"RegionId"`
}

func (req *SetDomainServerCertificateRequest) Invoke(client goaliyun.Client) (*SetDomainServerCertificateResponse, error) {
	resp := &SetDomainServerCertificateResponse{}
	err := client.Request("cdn", "SetDomainServerCertificate", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetDomainServerCertificateResponse struct {
	RequestId goaliyun.String
}
