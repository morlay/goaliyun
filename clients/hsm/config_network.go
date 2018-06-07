package hsm

import (
	"github.com/morlay/goaliyun"
)

type ConfigNetworkRequest struct {
	VSwitchId       string `position:"Query" name:"VSwitchId"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId      string `position:"Query" name:"InstanceId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	VpcId           string `position:"Query" name:"VpcId"`
	Ip              string `position:"Query" name:"Ip"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ConfigNetworkRequest) Invoke(client goaliyun.Client) (*ConfigNetworkResponse, error) {
	resp := &ConfigNetworkResponse{}
	err := client.Request("hsm", "ConfigNetwork", "2018-01-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ConfigNetworkResponse struct {
	RequestId goaliyun.String
}
