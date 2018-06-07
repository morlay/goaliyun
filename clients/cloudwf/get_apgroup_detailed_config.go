package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type GetApgroupDetailedConfigRequest struct {
	Id       int64  `position:"Query" name:"Id"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetApgroupDetailedConfigRequest) Invoke(client goaliyun.Client) (*GetApgroupDetailedConfigResponse, error) {
	resp := &GetApgroupDetailedConfigResponse{}
	err := client.Request("cloudwf", "GetApgroupDetailedConfig", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetApgroupDetailedConfigResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
