package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeCdnDomainDetailRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeCdnDomainDetailRequest) Invoke(client goaliyun.Client) (*DescribeCdnDomainDetailResponse, error) {
	resp := &DescribeCdnDomainDetailResponse{}
	err := client.Request("cdn", "DescribeCdnDomainDetail", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeCdnDomainDetailResponse struct {
	RequestId            goaliyun.String
	GetDomainDetailModel DescribeCdnDomainDetailGetDomainDetailModel
}

type DescribeCdnDomainDetailGetDomainDetailModel struct {
	GmtCreated              goaliyun.String
	GmtModified             goaliyun.String
	SourceType              goaliyun.String
	DomainStatus            goaliyun.String
	SourcePort              goaliyun.Integer
	CdnType                 goaliyun.String
	Cname                   goaliyun.String
	HttpsCname              goaliyun.String
	DomainName              goaliyun.String
	Description             goaliyun.String
	ServerCertificateStatus goaliyun.String
	ServerCertificate       goaliyun.String
	Region                  goaliyun.String
	Scope                   goaliyun.String
	CertificateName         goaliyun.String
	ResourceGroupId         goaliyun.String
	SourceModels            DescribeCdnDomainDetailSourceModelList
	Sources                 DescribeCdnDomainDetailSourceList
}

type DescribeCdnDomainDetailSourceModel struct {
	Content  goaliyun.String
	Type     goaliyun.String
	Port     goaliyun.Integer
	Enabled  goaliyun.String
	Priority goaliyun.String
}

type DescribeCdnDomainDetailSourceModelList []DescribeCdnDomainDetailSourceModel

func (list *DescribeCdnDomainDetailSourceModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCdnDomainDetailSourceModel)
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

type DescribeCdnDomainDetailSourceList []goaliyun.String

func (list *DescribeCdnDomainDetailSourceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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
