package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDomainsBySourceRequest struct {
	Sources       string `position:"Query" name:"Sources"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeDomainsBySourceRequest) Invoke(client goaliyun.Client) (*DescribeDomainsBySourceResponse, error) {
	resp := &DescribeDomainsBySourceResponse{}
	err := client.Request("cdn", "DescribeDomainsBySource", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDomainsBySourceResponse struct {
	RequestId   goaliyun.String
	Sources     goaliyun.String
	DomainsList DescribeDomainsBySourceDomainsDataList
}

type DescribeDomainsBySourceDomainsData struct {
	Source      goaliyun.String
	DomainInfos DescribeDomainsBySourceDomainInfoList
	Domains     DescribeDomainsBySourceDomainList
}

type DescribeDomainsBySourceDomainInfo struct {
	DomainName  goaliyun.String
	DomainCname goaliyun.String
	CreateTime  goaliyun.String
	UpdateTime  goaliyun.String
	Status      goaliyun.String
}

type DescribeDomainsBySourceDomainsDataList []DescribeDomainsBySourceDomainsData

func (list *DescribeDomainsBySourceDomainsDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainsBySourceDomainsData)
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

type DescribeDomainsBySourceDomainInfoList []DescribeDomainsBySourceDomainInfo

func (list *DescribeDomainsBySourceDomainInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainsBySourceDomainInfo)
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

type DescribeDomainsBySourceDomainList []goaliyun.String

func (list *DescribeDomainsBySourceDomainList) UnmarshalJSON(data []byte) error {
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
