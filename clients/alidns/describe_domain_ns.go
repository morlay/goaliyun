package alidns

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDomainNsRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *DescribeDomainNsRequest) Invoke(client goaliyun.Client) (*DescribeDomainNsResponse, error) {
	resp := &DescribeDomainNsResponse{}
	err := client.Request("alidns", "DescribeDomainNs", "2015-01-09", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDomainNsResponse struct {
	RequestId     goaliyun.String
	AllAliDns     bool
	IncludeAliDns bool
	DnsServers    DescribeDomainNsDnsServerList
}

type DescribeDomainNsDnsServerList []goaliyun.String

func (list *DescribeDomainNsDnsServerList) UnmarshalJSON(data []byte) error {
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
