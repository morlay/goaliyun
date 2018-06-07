package domain_intl

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryDnsHostRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
	Lang       string `position:"Query" name:"Lang"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *QueryDnsHostRequest) Invoke(client goaliyun.Client) (*QueryDnsHostResponse, error) {
	resp := &QueryDnsHostResponse{}
	err := client.Request("domain-intl", "QueryDnsHost", "2017-12-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryDnsHostResponse struct {
	RequestId   goaliyun.String
	DnsHostList QueryDnsHostDnsHostList
}

type QueryDnsHostDnsHost struct {
	DnsName goaliyun.String
	IpList  QueryDnsHostIpListList
}

type QueryDnsHostDnsHostList []QueryDnsHostDnsHost

func (list *QueryDnsHostDnsHostList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryDnsHostDnsHost)
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

type QueryDnsHostIpListList []goaliyun.String

func (list *QueryDnsHostIpListList) UnmarshalJSON(data []byte) error {
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
