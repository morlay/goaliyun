package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type SaveApMapInfoRequest struct {
	JsonData string `position:"Query" name:"JsonData"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *SaveApMapInfoRequest) Invoke(client goaliyun.Client) (*SaveApMapInfoResponse, error) {
	resp := &SaveApMapInfoResponse{}
	err := client.Request("cloudwf", "SaveApMapInfo", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SaveApMapInfoResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
