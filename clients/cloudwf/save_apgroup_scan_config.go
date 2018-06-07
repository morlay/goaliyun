package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type SaveApgroupScanConfigRequest struct {
	JsonData  string `position:"Query" name:"JsonData"`
	ApgroupId int64  `position:"Query" name:"ApgroupId"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *SaveApgroupScanConfigRequest) Invoke(client goaliyun.Client) (*SaveApgroupScanConfigResponse, error) {
	resp := &SaveApgroupScanConfigResponse{}
	err := client.Request("cloudwf", "SaveApgroupScanConfig", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SaveApgroupScanConfigResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
