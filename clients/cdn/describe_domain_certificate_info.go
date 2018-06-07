package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDomainCertificateInfoRequest struct {
	DomainName string `position:"Query" name:"DomainName"`
	OwnerId    int64  `position:"Query" name:"OwnerId"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *DescribeDomainCertificateInfoRequest) Invoke(client goaliyun.Client) (*DescribeDomainCertificateInfoResponse, error) {
	resp := &DescribeDomainCertificateInfoResponse{}
	err := client.Request("cdn", "DescribeDomainCertificateInfo", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDomainCertificateInfoResponse struct {
	RequestId goaliyun.String
	CertInfos DescribeDomainCertificateInfoCertInfoList
}

type DescribeDomainCertificateInfoCertInfo struct {
	DomainName              goaliyun.String
	CertName                goaliyun.String
	CertDomainName          goaliyun.String
	CertExpireTime          goaliyun.String
	CertLife                goaliyun.String
	CertOrg                 goaliyun.String
	CertType                goaliyun.String
	ServerCertificateStatus goaliyun.String
	Status                  goaliyun.String
}

type DescribeDomainCertificateInfoCertInfoList []DescribeDomainCertificateInfoCertInfo

func (list *DescribeDomainCertificateInfoCertInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainCertificateInfoCertInfo)
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
