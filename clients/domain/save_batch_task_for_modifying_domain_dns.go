package domain

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type SaveBatchTaskForModifyingDomainDnsRequest struct {
	UserClientIp      string                                                  `position:"Query" name:"UserClientIp"`
	DomainNames       *SaveBatchTaskForModifyingDomainDnsDomainNameList       `position:"Query" type:"Repeated" name:"DomainName"`
	DomainNameServers *SaveBatchTaskForModifyingDomainDnsDomainNameServerList `position:"Query" type:"Repeated" name:"DomainNameServer"`
	Lang              string                                                  `position:"Query" name:"Lang"`
	AliyunDns         string                                                  `position:"Query" name:"AliyunDns"`
	RegionId          string                                                  `position:"Query" name:"RegionId"`
}

func (req *SaveBatchTaskForModifyingDomainDnsRequest) Invoke(client goaliyun.Client) (*SaveBatchTaskForModifyingDomainDnsResponse, error) {
	resp := &SaveBatchTaskForModifyingDomainDnsResponse{}
	err := client.Request("domain", "SaveBatchTaskForModifyingDomainDns", "2018-01-29", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SaveBatchTaskForModifyingDomainDnsResponse struct {
	RequestId goaliyun.String
	TaskNo    goaliyun.String
}

type SaveBatchTaskForModifyingDomainDnsDomainNameList []string

func (list *SaveBatchTaskForModifyingDomainDnsDomainNameList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type SaveBatchTaskForModifyingDomainDnsDomainNameServerList []string

func (list *SaveBatchTaskForModifyingDomainDnsDomainNameServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
