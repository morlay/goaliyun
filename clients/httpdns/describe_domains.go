package httpdns

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDomainsRequest struct {
	PageSize   int64  `position:"Query" name:"PageSize"`
	PageNumber int64  `position:"Query" name:"PageNumber"`
	AccountId  string `position:"Query" name:"AccountId"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *DescribeDomainsRequest) Invoke(client goaliyun.Client) (*DescribeDomainsResponse, error) {
	resp := &DescribeDomainsResponse{}
	err := client.Request("httpdns", "DescribeDomains", "2016-02-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDomainsResponse struct {
	RequestId  goaliyun.String
	TotalCount goaliyun.Integer
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	Domains    DescribeDomainsDomainList
}

type DescribeDomainsDomain struct {
	DomainName goaliyun.String
}

type DescribeDomainsDomainList []DescribeDomainsDomain

func (list *DescribeDomainsDomainList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainsDomain)
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
