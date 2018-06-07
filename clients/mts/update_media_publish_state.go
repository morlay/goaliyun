package mts

import (
	"github.com/morlay/goaliyun"
)

type UpdateMediaPublishStateRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Publish              string `position:"Query" name:"Publish"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	MediaId              string `position:"Query" name:"MediaId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *UpdateMediaPublishStateRequest) Invoke(client goaliyun.Client) (*UpdateMediaPublishStateResponse, error) {
	resp := &UpdateMediaPublishStateResponse{}
	err := client.Request("mts", "UpdateMediaPublishState", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateMediaPublishStateResponse struct {
	RequestId goaliyun.String
}
