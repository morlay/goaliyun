package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type RepairApRequest struct {
	Id       int64  `position:"Query" name:"Id"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *RepairApRequest) Invoke(client goaliyun.Client) (*RepairApResponse, error) {
	resp := &RepairApResponse{}
	err := client.Request("cloudwf", "RepairAp", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RepairApResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
