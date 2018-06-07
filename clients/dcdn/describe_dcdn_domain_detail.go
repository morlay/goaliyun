package dcdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDcdnDomainDetailRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeDcdnDomainDetailRequest) Invoke(client goaliyun.Client) (*DescribeDcdnDomainDetailResponse, error) {
	resp := &DescribeDcdnDomainDetailResponse{}
	err := client.Request("dcdn", "DescribeDcdnDomainDetail", "2018-01-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDcdnDomainDetailResponse struct {
	RequestId    goaliyun.String
	DomainDetail DescribeDcdnDomainDetailDomainDetail
}

type DescribeDcdnDomainDetailDomainDetail struct {
	GmtCreated      goaliyun.String
	GmtModified     goaliyun.String
	DomainStatus    goaliyun.String
	Cname           goaliyun.String
	DomainName      goaliyun.String
	Description     goaliyun.String
	SSLProtocol     goaliyun.String
	SSLPub          goaliyun.String
	Scope           goaliyun.String
	CertName        goaliyun.String
	ResourceGroupId goaliyun.String
	Sources         DescribeDcdnDomainDetailSourceList
}

type DescribeDcdnDomainDetailSource struct {
	Content  goaliyun.String
	Type     goaliyun.String
	Port     goaliyun.Integer
	Enabled  goaliyun.String
	Priority goaliyun.String
}

type DescribeDcdnDomainDetailSourceList []DescribeDcdnDomainDetailSource

func (list *DescribeDcdnDomainDetailSourceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDcdnDomainDetailSource)
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
