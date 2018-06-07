package vod

import (
	"github.com/morlay/goaliyun"
)

type PushObjectCacheRequest struct {
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ObjectPath           string `position:"Query" name:"ObjectPath"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *PushObjectCacheRequest) Invoke(client goaliyun.Client) (*PushObjectCacheResponse, error) {
	resp := &PushObjectCacheResponse{}
	err := client.Request("vod", "PushObjectCache", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type PushObjectCacheResponse struct {
	RequestId  goaliyun.String
	PushTaskId goaliyun.String
}
