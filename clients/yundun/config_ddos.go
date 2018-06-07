package yundun

import (
	"github.com/morlay/goaliyun"
)

type ConfigDdosRequest struct {
	InstanceId       string `position:"Query" name:"InstanceId"`
	InstanceType     string `position:"Query" name:"InstanceType"`
	FlowPosition     int64  `position:"Query" name:"FlowPosition"`
	StrategyPosition int64  `position:"Query" name:"StrategyPosition"`
	Level            int64  `position:"Query" name:"Level"`
	RegionId         string `position:"Query" name:"RegionId"`
}

func (req *ConfigDdosRequest) Invoke(client goaliyun.Client) (*ConfigDdosResponse, error) {
	resp := &ConfigDdosResponse{}
	err := client.Request("yundun", "ConfigDdos", "2015-04-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ConfigDdosResponse struct {
	RequestId goaliyun.String
}
