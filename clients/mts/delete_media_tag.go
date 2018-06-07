package mts

import (
	"github.com/morlay/goaliyun"
)

type DeleteMediaTagRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Tag                  string `position:"Query" name:"Tag"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	MediaId              string `position:"Query" name:"MediaId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteMediaTagRequest) Invoke(client goaliyun.Client) (*DeleteMediaTagResponse, error) {
	resp := &DeleteMediaTagResponse{}
	err := client.Request("mts", "DeleteMediaTag", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteMediaTagResponse struct {
	RequestId goaliyun.String
}
