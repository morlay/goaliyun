package mts

import (
	"github.com/morlay/goaliyun"
)

type UpdateMediaCategoryRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	CateId               int64  `position:"Query" name:"CateId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	MediaId              string `position:"Query" name:"MediaId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *UpdateMediaCategoryRequest) Invoke(client goaliyun.Client) (*UpdateMediaCategoryResponse, error) {
	resp := &UpdateMediaCategoryResponse{}
	err := client.Request("mts", "UpdateMediaCategory", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateMediaCategoryResponse struct {
	RequestId goaliyun.String
}
