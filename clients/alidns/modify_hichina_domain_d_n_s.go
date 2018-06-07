package alidns

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ModifyHichinaDomainDNSRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *ModifyHichinaDomainDNSRequest) Invoke(client goaliyun.Client) (*ModifyHichinaDomainDNSResponse, error) {
	resp := &ModifyHichinaDomainDNSResponse{}
	err := client.Request("alidns", "ModifyHichinaDomainDNS", "2015-01-09", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyHichinaDomainDNSResponse struct {
	RequestId          goaliyun.String
	OriginalDnsServers ModifyHichinaDomainDNSOriginalDnsServerList
	NewDnsServers      ModifyHichinaDomainDNSNewDnsServerList
}

type ModifyHichinaDomainDNSOriginalDnsServerList []goaliyun.String

func (list *ModifyHichinaDomainDNSOriginalDnsServerList) UnmarshalJSON(data []byte) error {
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

type ModifyHichinaDomainDNSNewDnsServerList []goaliyun.String

func (list *ModifyHichinaDomainDNSNewDnsServerList) UnmarshalJSON(data []byte) error {
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
