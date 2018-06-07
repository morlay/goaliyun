package mts

import (
	"github.com/morlay/goaliyun"
)

type AddCategoryRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ParentId             int64  `position:"Query" name:"ParentId"`
	CateName             string `position:"Query" name:"CateName"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *AddCategoryRequest) Invoke(client goaliyun.Client) (*AddCategoryResponse, error) {
	resp := &AddCategoryResponse{}
	err := client.Request("mts", "AddCategory", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddCategoryResponse struct {
	RequestId goaliyun.String
	Category  AddCategoryCategory
}

type AddCategoryCategory struct {
	CateId   goaliyun.String
	CateName goaliyun.String
	ParentId goaliyun.String
	Level    goaliyun.String
}
