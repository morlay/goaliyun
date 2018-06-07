package mts

import (
	"github.com/morlay/goaliyun"
)

type CategoryTreeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CategoryTreeRequest) Invoke(client goaliyun.Client) (*CategoryTreeResponse, error) {
	resp := &CategoryTreeResponse{}
	err := client.Request("mts", "CategoryTree", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CategoryTreeResponse struct {
	RequestId    goaliyun.String
	CategoryTree goaliyun.String
}
