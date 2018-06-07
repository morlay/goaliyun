package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type GetOnlineStaTimeSerRequest struct {
	ZoomStart int64  `position:"Query" name:"ZoomStart"`
	CompanyId int64  `position:"Query" name:"CompanyId"`
	ApgroupId int64  `position:"Query" name:"ApgroupId"`
	Start     int64  `position:"Query" name:"Start"`
	ZoomEnd   int64  `position:"Query" name:"ZoomEnd"`
	End       int64  `position:"Query" name:"End"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *GetOnlineStaTimeSerRequest) Invoke(client goaliyun.Client) (*GetOnlineStaTimeSerResponse, error) {
	resp := &GetOnlineStaTimeSerResponse{}
	err := client.Request("cloudwf", "GetOnlineStaTimeSer", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetOnlineStaTimeSerResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
