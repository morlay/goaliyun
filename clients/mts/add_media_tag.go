package mts

import (
	"github.com/morlay/goaliyun"
)

type AddMediaTagRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Tag                  string `position:"Query" name:"Tag"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	MediaId              string `position:"Query" name:"MediaId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *AddMediaTagRequest) Invoke(client goaliyun.Client) (*AddMediaTagResponse, error) {
	resp := &AddMediaTagResponse{}
	err := client.Request("mts", "AddMediaTag", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddMediaTagResponse struct {
	RequestId goaliyun.String
}
