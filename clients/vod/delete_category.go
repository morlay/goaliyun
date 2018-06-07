package vod

import (
	"github.com/morlay/goaliyun"
)

type DeleteCategoryRequest struct {
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	CateId               int64  `position:"Query" name:"CateId"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteCategoryRequest) Invoke(client goaliyun.Client) (*DeleteCategoryResponse, error) {
	resp := &DeleteCategoryResponse{}
	err := client.Request("vod", "DeleteCategory", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteCategoryResponse struct {
	RequestId goaliyun.String
}
