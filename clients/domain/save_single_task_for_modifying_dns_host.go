package domain

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type SaveSingleTaskForModifyingDnsHostRequest struct {
	InstanceId string                                   `position:"Query" name:"InstanceId"`
	Ips        *SaveSingleTaskForModifyingDnsHostIpList `position:"Query" type:"Repeated" name:"Ip"`
	DnsName    string                                   `position:"Query" name:"DnsName"`
	Lang       string                                   `position:"Query" name:"Lang"`
	RegionId   string                                   `position:"Query" name:"RegionId"`
}

func (req *SaveSingleTaskForModifyingDnsHostRequest) Invoke(client goaliyun.Client) (*SaveSingleTaskForModifyingDnsHostResponse, error) {
	resp := &SaveSingleTaskForModifyingDnsHostResponse{}
	err := client.Request("domain", "SaveSingleTaskForModifyingDnsHost", "2018-01-29", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SaveSingleTaskForModifyingDnsHostResponse struct {
	RequestId goaliyun.String
	TaskNo    goaliyun.String
}

type SaveSingleTaskForModifyingDnsHostIpList []string

func (list *SaveSingleTaskForModifyingDnsHostIpList) UnmarshalJSON(data []byte) error {
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
