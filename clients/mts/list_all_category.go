package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListAllCategoryRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ListAllCategoryRequest) Invoke(client goaliyun.Client) (*ListAllCategoryResponse, error) {
	resp := &ListAllCategoryResponse{}
	err := client.Request("mts", "ListAllCategory", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListAllCategoryResponse struct {
	RequestId    goaliyun.String
	CategoryList ListAllCategoryCategoryList
}

type ListAllCategoryCategory struct {
	CateId   goaliyun.String
	CateName goaliyun.String
	ParentId goaliyun.String
	Level    goaliyun.String
}

type ListAllCategoryCategoryList []ListAllCategoryCategory

func (list *ListAllCategoryCategoryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAllCategoryCategory)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
