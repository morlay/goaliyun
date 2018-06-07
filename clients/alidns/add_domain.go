package alidns

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type AddDomainRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
	GroupId      string `position:"Query" name:"GroupId"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *AddDomainRequest) Invoke(client goaliyun.Client) (*AddDomainResponse, error) {
	resp := &AddDomainResponse{}
	err := client.Request("alidns", "AddDomain", "2015-01-09", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddDomainResponse struct {
	RequestId  goaliyun.String
	DomainId   goaliyun.String
	DomainName goaliyun.String
	PunyCode   goaliyun.String
	GroupId    goaliyun.String
	GroupName  goaliyun.String
	DnsServers AddDomainDnsServerList
}

type AddDomainDnsServerList []goaliyun.String

func (list *AddDomainDnsServerList) UnmarshalJSON(data []byte) error {
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
