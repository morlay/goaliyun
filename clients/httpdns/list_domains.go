package httpdns

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListDomainsRequest struct {
	PageSize   int64  `position:"Query" name:"PageSize"`
	PageNumber int64  `position:"Query" name:"PageNumber"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *ListDomainsRequest) Invoke(client goaliyun.Client) (*ListDomainsResponse, error) {
	resp := &ListDomainsResponse{}
	err := client.Request("httpdns", "ListDomains", "2016-02-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListDomainsResponse struct {
	RequestId   goaliyun.String
	TotalCount  goaliyun.Integer
	PageNumber  goaliyun.Integer
	PageSize    goaliyun.Integer
	DomainInfos ListDomainsDomainInfoList
}

type ListDomainsDomainInfo struct {
	DomainName    goaliyun.String
	Resolved      goaliyun.Integer
	ResolvedHttps goaliyun.Integer
}

type ListDomainsDomainInfoList []ListDomainsDomainInfo

func (list *ListDomainsDomainInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListDomainsDomainInfo)
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
