package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type DeletePositionMapRequest struct {
	MapId    int64  `position:"Query" name:"MapId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *DeletePositionMapRequest) Invoke(client goaliyun.Client) (*DeletePositionMapResponse, error) {
	resp := &DeletePositionMapResponse{}
	err := client.Request("cloudwf", "DeletePositionMap", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeletePositionMapResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
