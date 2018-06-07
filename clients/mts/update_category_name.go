package mts

import (
	"github.com/morlay/goaliyun"
)

type UpdateCategoryNameRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	CateId               string `position:"Query" name:"CateId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	CateName             string `position:"Query" name:"CateName"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *UpdateCategoryNameRequest) Invoke(client goaliyun.Client) (*UpdateCategoryNameResponse, error) {
	resp := &UpdateCategoryNameResponse{}
	err := client.Request("mts", "UpdateCategoryName", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateCategoryNameResponse struct {
	RequestId goaliyun.String
}
