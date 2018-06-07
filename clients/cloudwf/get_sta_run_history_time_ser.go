package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type GetStaRunHistoryTimeSerRequest struct {
	Id       int64  `position:"Query" name:"Id"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetStaRunHistoryTimeSerRequest) Invoke(client goaliyun.Client) (*GetStaRunHistoryTimeSerResponse, error) {
	resp := &GetStaRunHistoryTimeSerResponse{}
	err := client.Request("cloudwf", "GetStaRunHistoryTimeSer", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetStaRunHistoryTimeSerResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
