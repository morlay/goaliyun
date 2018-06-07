package vod

import (
	"github.com/morlay/goaliyun"
)

type AddCategoryRequest struct {
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	ParentId             int64  `position:"Query" name:"ParentId"`
	CateName             string `position:"Query" name:"CateName"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *AddCategoryRequest) Invoke(client goaliyun.Client) (*AddCategoryResponse, error) {
	resp := &AddCategoryResponse{}
	err := client.Request("vod", "AddCategory", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddCategoryResponse struct {
	RequestId goaliyun.String
	Category  AddCategoryCategory
}

type AddCategoryCategory struct {
	CateId   goaliyun.Integer
	CateName goaliyun.String
	ParentId goaliyun.Integer
	Level    goaliyun.Integer
}
