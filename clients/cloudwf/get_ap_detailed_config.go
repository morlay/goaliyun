package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type GetApDetailedConfigRequest struct {
	Id       int64  `position:"Query" name:"Id"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetApDetailedConfigRequest) Invoke(client goaliyun.Client) (*GetApDetailedConfigResponse, error) {
	resp := &GetApDetailedConfigResponse{}
	err := client.Request("cloudwf", "GetApDetailedConfig", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetApDetailedConfigResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
