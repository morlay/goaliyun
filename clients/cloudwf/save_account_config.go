package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type SaveAccountConfigRequest struct {
	JsonData string `position:"Query" name:"JsonData"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *SaveAccountConfigRequest) Invoke(client goaliyun.Client) (*SaveAccountConfigResponse, error) {
	resp := &SaveAccountConfigResponse{}
	err := client.Request("cloudwf", "SaveAccountConfig", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SaveAccountConfigResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
