package vod

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetCategoriesRequest struct {
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	CateId               int64  `position:"Query" name:"CateId"`
	PageNo               int64  `position:"Query" name:"PageNo"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *GetCategoriesRequest) Invoke(client goaliyun.Client) (*GetCategoriesResponse, error) {
	resp := &GetCategoriesResponse{}
	err := client.Request("vod", "GetCategories", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetCategoriesResponse struct {
	RequestId     goaliyun.String
	SubTotal      goaliyun.Integer
	SubCategories GetCategoriesCategoryList
	Category1     GetCategoriesCategory1
}

type GetCategoriesCategory struct {
	CateId   goaliyun.Integer
	CateName goaliyun.String
	Level    goaliyun.Integer
	ParentId goaliyun.Integer
}

type GetCategoriesCategory1 struct {
	CateId   goaliyun.Integer
	CateName goaliyun.String
	Level    goaliyun.Integer
	ParentId goaliyun.Integer
}

type GetCategoriesCategoryList []GetCategoriesCategory

func (list *GetCategoriesCategoryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetCategoriesCategory)
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
