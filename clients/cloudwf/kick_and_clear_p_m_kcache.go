package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type KickAndClearPMKcacheRequest struct {
	Id       int64  `position:"Query" name:"Id"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *KickAndClearPMKcacheRequest) Invoke(client goaliyun.Client) (*KickAndClearPMKcacheResponse, error) {
	resp := &KickAndClearPMKcacheResponse{}
	err := client.Request("cloudwf", "KickAndClearPMKcache", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type KickAndClearPMKcacheResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
