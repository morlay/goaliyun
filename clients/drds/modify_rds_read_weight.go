package drds

import (
	"github.com/morlay/goaliyun"
)

type ModifyRdsReadWeightRequest struct {
	InstanceNames  string `position:"Query" name:"InstanceNames"`
	DbName         string `position:"Query" name:"DbName"`
	Weights        string `position:"Query" name:"Weights"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *ModifyRdsReadWeightRequest) Invoke(client goaliyun.Client) (*ModifyRdsReadWeightResponse, error) {
	resp := &ModifyRdsReadWeightResponse{}
	err := client.Request("drds", "ModifyRdsReadWeight", "2017-10-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyRdsReadWeightResponse struct {
	RequestId goaliyun.String
	Success   bool
}
