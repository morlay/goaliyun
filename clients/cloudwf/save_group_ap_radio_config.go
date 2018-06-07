package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type SaveGroupApRadioConfigRequest struct {
	JsonData string `position:"Query" name:"JsonData"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *SaveGroupApRadioConfigRequest) Invoke(client goaliyun.Client) (*SaveGroupApRadioConfigResponse, error) {
	resp := &SaveGroupApRadioConfigResponse{}
	err := client.Request("cloudwf", "SaveGroupApRadioConfig", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SaveGroupApRadioConfigResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
