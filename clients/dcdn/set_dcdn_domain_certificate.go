package dcdn

import (
	"github.com/morlay/goaliyun"
)

type SetDcdnDomainCertificateRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	CertType      string `position:"Query" name:"CertType"`
	SSLPub        string `position:"Query" name:"SSLPub"`
	CertName      string `position:"Query" name:"CertName"`
	SSLProtocol   string `position:"Query" name:"SSLProtocol"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Region        string `position:"Query" name:"Region"`
	SSLPri        string `position:"Query" name:"SSLPri"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *SetDcdnDomainCertificateRequest) Invoke(client goaliyun.Client) (*SetDcdnDomainCertificateResponse, error) {
	resp := &SetDcdnDomainCertificateResponse{}
	err := client.Request("dcdn", "SetDcdnDomainCertificate", "2018-01-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetDcdnDomainCertificateResponse struct {
	RequestId goaliyun.String
}
