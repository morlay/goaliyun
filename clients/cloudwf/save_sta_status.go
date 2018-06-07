package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type SaveStaStatusRequest struct {
	Description string `position:"Query" name:"Description"`
	Id          int64  `position:"Query" name:"Id"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *SaveStaStatusRequest) Invoke(client goaliyun.Client) (*SaveStaStatusResponse, error) {
	resp := &SaveStaStatusResponse{}
	err := client.Request("cloudwf", "SaveStaStatus", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SaveStaStatusResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
