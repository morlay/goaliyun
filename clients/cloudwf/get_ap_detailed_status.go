package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type GetApDetailedStatusRequest struct {
	Id       int64  `position:"Query" name:"Id"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetApDetailedStatusRequest) Invoke(client goaliyun.Client) (*GetApDetailedStatusResponse, error) {
	resp := &GetApDetailedStatusResponse{}
	err := client.Request("cloudwf", "GetApDetailedStatus", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetApDetailedStatusResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
