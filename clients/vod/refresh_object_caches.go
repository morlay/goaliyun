package vod

import (
	"github.com/morlay/goaliyun"
)

type RefreshObjectCachesRequest struct {
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ObjectPath           string `position:"Query" name:"ObjectPath"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	ObjectType           string `position:"Query" name:"ObjectType"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *RefreshObjectCachesRequest) Invoke(client goaliyun.Client) (*RefreshObjectCachesResponse, error) {
	resp := &RefreshObjectCachesResponse{}
	err := client.Request("vod", "RefreshObjectCaches", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RefreshObjectCachesResponse struct {
	RequestId     goaliyun.String
	RefreshTaskId goaliyun.String
}
