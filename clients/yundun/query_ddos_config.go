package yundun

import (
	"github.com/morlay/goaliyun"
)

type QueryDdosConfigRequest struct {
	InstanceId   string `position:"Query" name:"InstanceId"`
	InstanceType string `position:"Query" name:"InstanceType"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *QueryDdosConfigRequest) Invoke(client goaliyun.Client) (*QueryDdosConfigResponse, error) {
	resp := &QueryDdosConfigResponse{}
	err := client.Request("yundun", "QueryDdosConfig", "2015-04-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryDdosConfigResponse struct {
	RequestId        goaliyun.String
	Bps              goaliyun.Integer
	Pps              goaliyun.Integer
	Qps              goaliyun.Integer
	Sipconn          goaliyun.Integer
	Sipnew           goaliyun.Integer
	Layer7Config     bool
	FlowPosition     goaliyun.Integer
	QpsPosition      goaliyun.Integer
	StrategyPosition goaliyun.Integer
	Level            goaliyun.Integer
	HoleBps          goaliyun.String
	ConfigType       goaliyun.String
}
