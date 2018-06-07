package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type GetApRunHistoryTimeSerRequest struct {
	Start    int64  `position:"Query" name:"Start"`
	End      int64  `position:"Query" name:"End"`
	Id       int64  `position:"Query" name:"Id"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetApRunHistoryTimeSerRequest) Invoke(client goaliyun.Client) (*GetApRunHistoryTimeSerResponse, error) {
	resp := &GetApRunHistoryTimeSerResponse{}
	err := client.Request("cloudwf", "GetApRunHistoryTimeSer", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetApRunHistoryTimeSerResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
