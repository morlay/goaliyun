package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type SaveApgroupSsidConfigRequest struct {
	JsonData string `position:"Query" name:"JsonData"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *SaveApgroupSsidConfigRequest) Invoke(client goaliyun.Client) (*SaveApgroupSsidConfigResponse, error) {
	resp := &SaveApgroupSsidConfigResponse{}
	err := client.Request("cloudwf", "SaveApgroupSsidConfig", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SaveApgroupSsidConfigResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
