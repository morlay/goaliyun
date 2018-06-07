package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type DeleteApgroupSsidConfigRequest struct {
	ApgroupId int64  `position:"Query" name:"ApgroupId"`
	Id        int64  `position:"Query" name:"Id"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *DeleteApgroupSsidConfigRequest) Invoke(client goaliyun.Client) (*DeleteApgroupSsidConfigResponse, error) {
	resp := &DeleteApgroupSsidConfigResponse{}
	err := client.Request("cloudwf", "DeleteApgroupSsidConfig", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteApgroupSsidConfigResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
