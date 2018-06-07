package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type KickStaRequest struct {
	Id       int64  `position:"Query" name:"Id"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *KickStaRequest) Invoke(client goaliyun.Client) (*KickStaResponse, error) {
	resp := &KickStaResponse{}
	err := client.Request("cloudwf", "KickSta", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type KickStaResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
