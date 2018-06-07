package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type DeleteApgroupConfigRequest struct {
	Id       int64  `position:"Query" name:"Id"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *DeleteApgroupConfigRequest) Invoke(client goaliyun.Client) (*DeleteApgroupConfigResponse, error) {
	resp := &DeleteApgroupConfigResponse{}
	err := client.Request("cloudwf", "DeleteApgroupConfig", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteApgroupConfigResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
