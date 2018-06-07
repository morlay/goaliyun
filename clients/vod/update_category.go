package vod

import (
	"github.com/morlay/goaliyun"
)

type UpdateCategoryRequest struct {
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	CateId               int64  `position:"Query" name:"CateId"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	CateName             string `position:"Query" name:"CateName"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *UpdateCategoryRequest) Invoke(client goaliyun.Client) (*UpdateCategoryResponse, error) {
	resp := &UpdateCategoryResponse{}
	err := client.Request("vod", "UpdateCategory", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateCategoryResponse struct {
	RequestId goaliyun.String
}
