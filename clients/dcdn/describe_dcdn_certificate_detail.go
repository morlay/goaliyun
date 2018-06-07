package dcdn

import (
	"github.com/morlay/goaliyun"
)

type DescribeDcdnCertificateDetailRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	CertName      string `position:"Query" name:"CertName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeDcdnCertificateDetailRequest) Invoke(client goaliyun.Client) (*DescribeDcdnCertificateDetailResponse, error) {
	resp := &DescribeDcdnCertificateDetailResponse{}
	err := client.Request("dcdn", "DescribeDcdnCertificateDetail", "2018-01-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDcdnCertificateDetailResponse struct {
	RequestId goaliyun.String
	Cert      goaliyun.String
	Key       goaliyun.String
	CertId    goaliyun.Integer
	CertName  goaliyun.String
}
