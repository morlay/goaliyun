package aegis

import (
	"github.com/morlay/goaliyun"
)

type ModifyBatchIgnoreVulRequest struct {
	Reason          string `position:"Query" name:"Reason"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	Info            string `position:"Query" name:"Info"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ModifyBatchIgnoreVulRequest) Invoke(client goaliyun.Client) (*ModifyBatchIgnoreVulResponse, error) {
	resp := &ModifyBatchIgnoreVulResponse{}
	err := client.Request("aegis", "ModifyBatchIgnoreVul", "2016-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyBatchIgnoreVulResponse struct {
	RequestId      goaliyun.String
	Success        bool
	HttpStatusCode goaliyun.Integer
}
