package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type GetOnlineApTimeSerRequest struct {
	ZoomStart int64  `position:"Query" name:"ZoomStart"`
	CompanyId int64  `position:"Query" name:"CompanyId"`
	ApgroupId int64  `position:"Query" name:"ApgroupId"`
	Start     int64  `position:"Query" name:"Start"`
	ZoomEnd   int64  `position:"Query" name:"ZoomEnd"`
	End       int64  `position:"Query" name:"End"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *GetOnlineApTimeSerRequest) Invoke(client goaliyun.Client) (*GetOnlineApTimeSerResponse, error) {
	resp := &GetOnlineApTimeSerResponse{}
	err := client.Request("cloudwf", "GetOnlineApTimeSer", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetOnlineApTimeSerResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
