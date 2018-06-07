package alidns

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDomainWhoisInfoRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
	GroupId      string `position:"Query" name:"GroupId"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *DescribeDomainWhoisInfoRequest) Invoke(client goaliyun.Client) (*DescribeDomainWhoisInfoResponse, error) {
	resp := &DescribeDomainWhoisInfoResponse{}
	err := client.Request("alidns", "DescribeDomainWhoisInfo", "2015-01-09", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDomainWhoisInfoResponse struct {
	RequestId        goaliyun.String
	RegistrantName   goaliyun.String
	RegistrantEmail  goaliyun.String
	Registrar        goaliyun.String
	RegistrationDate goaliyun.String
	ExpirationDate   goaliyun.String
	StatusList       DescribeDomainWhoisInfoStatusListList
	DnsServers       DescribeDomainWhoisInfoDnsServerList
}

type DescribeDomainWhoisInfoStatusListList []goaliyun.String

func (list *DescribeDomainWhoisInfoStatusListList) UnmarshalJSON(data []byte) error {
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

type DescribeDomainWhoisInfoDnsServerList []goaliyun.String

func (list *DescribeDomainWhoisInfoDnsServerList) UnmarshalJSON(data []byte) error {
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
