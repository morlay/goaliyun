package mts

import (
	"github.com/morlay/goaliyun"
)

type DeleteCategoryRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	CateId               int64  `position:"Query" name:"CateId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteCategoryRequest) Invoke(client goaliyun.Client) (*DeleteCategoryResponse, error) {
	resp := &DeleteCategoryResponse{}
	err := client.Request("mts", "DeleteCategory", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteCategoryResponse struct {
	RequestId goaliyun.String
}
