package mts

import (
	"github.com/morlay/goaliyun"
)

type UpdateMediaCoverRequest struct {
	CoverURL             string `position:"Query" name:"CoverURL"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	MediaId              string `position:"Query" name:"MediaId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *UpdateMediaCoverRequest) Invoke(client goaliyun.Client) (*UpdateMediaCoverResponse, error) {
	resp := &UpdateMediaCoverResponse{}
	err := client.Request("mts", "UpdateMediaCover", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateMediaCoverResponse struct {
	RequestId goaliyun.String
}
