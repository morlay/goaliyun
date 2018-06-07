package hsm

import (
	"github.com/morlay/goaliyun"
)

type ModifyInstanceRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId      string `position:"Query" name:"InstanceId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	Remark          string `position:"Query" name:"Remark"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ModifyInstanceRequest) Invoke(client goaliyun.Client) (*ModifyInstanceResponse, error) {
	resp := &ModifyInstanceResponse{}
	err := client.Request("hsm", "ModifyInstance", "2018-01-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyInstanceResponse struct {
	RequestId goaliyun.String
}
