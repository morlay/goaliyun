package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type RepairGroupApRequest struct {
	Id       int64  `position:"Query" name:"Id"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *RepairGroupApRequest) Invoke(client goaliyun.Client) (*RepairGroupApResponse, error) {
	resp := &RepairGroupApResponse{}
	err := client.Request("cloudwf", "RepairGroupAp", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RepairGroupApResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
