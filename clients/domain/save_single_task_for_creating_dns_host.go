package domain

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type SaveSingleTaskForCreatingDnsHostRequest struct {
	InstanceId string                                  `position:"Query" name:"InstanceId"`
	Ips        *SaveSingleTaskForCreatingDnsHostIpList `position:"Query" type:"Repeated" name:"Ip"`
	DnsName    string                                  `position:"Query" name:"DnsName"`
	Lang       string                                  `position:"Query" name:"Lang"`
	RegionId   string                                  `position:"Query" name:"RegionId"`
}

func (req *SaveSingleTaskForCreatingDnsHostRequest) Invoke(client goaliyun.Client) (*SaveSingleTaskForCreatingDnsHostResponse, error) {
	resp := &SaveSingleTaskForCreatingDnsHostResponse{}
	err := client.Request("domain", "SaveSingleTaskForCreatingDnsHost", "2018-01-29", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SaveSingleTaskForCreatingDnsHostResponse struct {
	RequestId goaliyun.String
	TaskNo    goaliyun.String
}

type SaveSingleTaskForCreatingDnsHostIpList []string

func (list *SaveSingleTaskForCreatingDnsHostIpList) UnmarshalJSON(data []byte) error {
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
