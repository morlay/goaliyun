package alidns

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDomainInfoRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *DescribeDomainInfoRequest) Invoke(client goaliyun.Client) (*DescribeDomainInfoResponse, error) {
	resp := &DescribeDomainInfoResponse{}
	err := client.Request("alidns", "DescribeDomainInfo", "2015-01-09", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDomainInfoResponse struct {
	RequestId       goaliyun.String
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
	DnsServers      DescribeDomainInfoDnsServerList
}

type DescribeDomainInfoDnsServerList []goaliyun.String

func (list *DescribeDomainInfoDnsServerList) UnmarshalJSON(data []byte) error {
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
