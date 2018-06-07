package dcdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDcdnUserDomainsRequest struct {
	FuncFilter       string `position:"Query" name:"FuncFilter"`
	DomainName       string `position:"Query" name:"DomainName"`
	OwnerId          int64  `position:"Query" name:"OwnerId"`
	FuncId           string `position:"Query" name:"FuncId"`
	PageNumber       int64  `position:"Query" name:"PageNumber"`
	DomainStatus     string `position:"Query" name:"DomainStatus"`
	DomainSearchType string `position:"Query" name:"DomainSearchType"`
	CheckDomainShow  string `position:"Query" name:"CheckDomainShow"`
	ResourceGroupId  string `position:"Query" name:"ResourceGroupId"`
	SecurityToken    string `position:"Query" name:"SecurityToken"`
	PageSize         int64  `position:"Query" name:"PageSize"`
	RegionId         string `position:"Query" name:"RegionId"`
}

func (req *DescribeDcdnUserDomainsRequest) Invoke(client goaliyun.Client) (*DescribeDcdnUserDomainsResponse, error) {
	resp := &DescribeDcdnUserDomainsResponse{}
	err := client.Request("dcdn", "DescribeDcdnUserDomains", "2018-01-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDcdnUserDomainsResponse struct {
	RequestId   goaliyun.String
	PageNumber  goaliyun.Integer
	PageSize    goaliyun.Integer
	TotalCount  goaliyun.Integer
	OnlineCount goaliyun.Integer
	Domains     DescribeDcdnUserDomainsPageDataList
}

type DescribeDcdnUserDomainsPageData struct {
	DomainName      goaliyun.String
	Cname           goaliyun.String
	DomainStatus    goaliyun.String
	GmtCreated      goaliyun.String
	GmtModified     goaliyun.String
	Description     goaliyun.String
	SSLProtocol     goaliyun.String
	ResourceGroupId goaliyun.String
	Sandbox         goaliyun.String
	Sources         DescribeDcdnUserDomainsSourceList
}

type DescribeDcdnUserDomainsSource struct {
	Type     goaliyun.String
	Content  goaliyun.String
	Port     goaliyun.Integer
	Priority goaliyun.String
}

type DescribeDcdnUserDomainsPageDataList []DescribeDcdnUserDomainsPageData

func (list *DescribeDcdnUserDomainsPageDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDcdnUserDomainsPageData)
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

type DescribeDcdnUserDomainsSourceList []DescribeDcdnUserDomainsSource

func (list *DescribeDcdnUserDomainsSourceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDcdnUserDomainsSource)
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
