package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type GetStaDetailedStatusRequest struct {
	Id       int64  `position:"Query" name:"Id"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetStaDetailedStatusRequest) Invoke(client goaliyun.Client) (*GetStaDetailedStatusResponse, error) {
	resp := &GetStaDetailedStatusResponse{}
	err := client.Request("cloudwf", "GetStaDetailedStatus", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetStaDetailedStatusResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
