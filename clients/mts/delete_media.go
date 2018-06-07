package mts

import (
	"github.com/morlay/goaliyun"
)

type DeleteMediaRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	MediaIds             string `position:"Query" name:"MediaIds"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteMediaRequest) Invoke(client goaliyun.Client) (*DeleteMediaResponse, error) {
	resp := &DeleteMediaResponse{}
	err := client.Request("mts", "DeleteMedia", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteMediaResponse struct {
	RequestId goaliyun.String
}
