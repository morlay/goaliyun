package alidns

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDomainsRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	KeyWord      string `position:"Query" name:"KeyWord"`
	GroupId      string `position:"Query" name:"GroupId"`
	PageNumber   int64  `position:"Query" name:"PageNumber"`
	PageSize     int64  `position:"Query" name:"PageSize"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *DescribeDomainsRequest) Invoke(client goaliyun.Client) (*DescribeDomainsResponse, error) {
	resp := &DescribeDomainsResponse{}
	err := client.Request("alidns", "DescribeDomains", "2015-01-09", req.RegionId, "", "").Do(req, resp)
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
	DomainId        goaliyun.String
	DomainName      goaliyun.String
	PunyCode        goaliyun.String
	AliDomain       bool
	RegistrantEmail goaliyun.String
	GroupId         goaliyun.String
	GroupName       goaliyun.String
	InstanceId      goaliyun.String
	VersionCode     goaliyun.String
	VersionName     goaliyun.String
	DnsServers      DescribeDomainsDnsServerList
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

type DescribeDomainsDnsServerList []goaliyun.String

func (list *DescribeDomainsDnsServerList) UnmarshalJSON(data []byte) error {
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
