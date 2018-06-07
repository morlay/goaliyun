package hsm

import (
	"github.com/morlay/goaliyun"
)

type ConfigWhiteListRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId      string `position:"Query" name:"InstanceId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	WhiteList       string `position:"Query" name:"WhiteList"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ConfigWhiteListRequest) Invoke(client goaliyun.Client) (*ConfigWhiteListResponse, error) {
	resp := &ConfigWhiteListResponse{}
	err := client.Request("hsm", "ConfigWhiteList", "2018-01-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ConfigWhiteListResponse struct {
	RequestId goaliyun.String
}
