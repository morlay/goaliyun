package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type GetRadioRunHistoryTimeSerRequest struct {
	Id       int64  `position:"Query" name:"Id"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetRadioRunHistoryTimeSerRequest) Invoke(client goaliyun.Client) (*GetRadioRunHistoryTimeSerResponse, error) {
	resp := &GetRadioRunHistoryTimeSerResponse{}
	err := client.Request("cloudwf", "GetRadioRunHistoryTimeSer", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetRadioRunHistoryTimeSerResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
