package dcdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDcdnCertificateListRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeDcdnCertificateListRequest) Invoke(client goaliyun.Client) (*DescribeDcdnCertificateListResponse, error) {
	resp := &DescribeDcdnCertificateListResponse{}
	err := client.Request("dcdn", "DescribeDcdnCertificateList", "2018-01-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDcdnCertificateListResponse struct {
	RequestId            goaliyun.String
	CertificateListModel DescribeDcdnCertificateListCertificateListModel
}

type DescribeDcdnCertificateListCertificateListModel struct {
	Count    goaliyun.Integer
	CertList DescribeDcdnCertificateListCertList
}

type DescribeDcdnCertificateListCert struct {
	CertName    goaliyun.String
	CertId      goaliyun.Integer
	Fingerprint goaliyun.String
	Common      goaliyun.String
	Issuer      goaliyun.String
	LastTime    goaliyun.Integer
}

type DescribeDcdnCertificateListCertList []DescribeDcdnCertificateListCert

func (list *DescribeDcdnCertificateListCertList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDcdnCertificateListCert)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
