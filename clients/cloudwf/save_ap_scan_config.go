package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type SaveApScanConfigRequest struct {
	JsonData   string `position:"Query" name:"JsonData"`
	ApConfigId int64  `position:"Query" name:"ApConfigId"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *SaveApScanConfigRequest) Invoke(client goaliyun.Client) (*SaveApScanConfigResponse, error) {
	resp := &SaveApScanConfigResponse{}
	err := client.Request("cloudwf", "SaveApScanConfig", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SaveApScanConfigResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
