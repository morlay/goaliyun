package aegis

import (
	"github.com/morlay/goaliyun"
)

type ModifyStrategyTargetRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	Type            string `position:"Query" name:"Type"`
	Config          string `position:"Query" name:"Config"`
	Target          string `position:"Query" name:"Target"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ModifyStrategyTargetRequest) Invoke(client goaliyun.Client) (*ModifyStrategyTargetResponse, error) {
	resp := &ModifyStrategyTargetResponse{}
	err := client.Request("aegis", "ModifyStrategyTarget", "2016-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyStrategyTargetResponse struct {
	RequestId goaliyun.String
}
