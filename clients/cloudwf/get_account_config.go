package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type GetAccountConfigRequest struct {
	Id       int64  `position:"Query" name:"Id"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetAccountConfigRequest) Invoke(client goaliyun.Client) (*GetAccountConfigResponse, error) {
	resp := &GetAccountConfigResponse{}
	err := client.Request("cloudwf", "GetAccountConfig", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetAccountConfigResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
