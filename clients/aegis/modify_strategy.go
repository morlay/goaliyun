package aegis

import (
	"github.com/morlay/goaliyun"
)

type ModifyStrategyRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	RiskSubTypeName string `position:"Query" name:"RiskSubTypeName"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	CycleStartTime  string `position:"Query" name:"CycleStartTime"`
	Name            string `position:"Query" name:"Name"`
	CycleDays       string `position:"Query" name:"CycleDays"`
	Id              string `position:"Query" name:"Id"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ModifyStrategyRequest) Invoke(client goaliyun.Client) (*ModifyStrategyResponse, error) {
	resp := &ModifyStrategyResponse{}
	err := client.Request("aegis", "ModifyStrategy", "2016-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyStrategyResponse struct {
	RequestId      goaliyun.String
	Success        bool
	TotalCount     goaliyun.Integer
	HttpStatusCode goaliyun.Integer
	Result         ModifyStrategyResult
}

type ModifyStrategyResult struct {
	StrategyId goaliyun.Integer
}
