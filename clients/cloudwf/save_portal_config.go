package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type SavePortalConfigRequest struct {
	JsonData string `position:"Query" name:"JsonData"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *SavePortalConfigRequest) Invoke(client goaliyun.Client) (*SavePortalConfigResponse, error) {
	resp := &SavePortalConfigResponse{}
	err := client.Request("cloudwf", "SavePortalConfig", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SavePortalConfigResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
