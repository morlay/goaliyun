package cms

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListMyGroupCategoriesRequest struct {
	GroupId  int64  `position:"Query" name:"GroupId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *ListMyGroupCategoriesRequest) Invoke(client goaliyun.Client) (*ListMyGroupCategoriesResponse, error) {
	resp := &ListMyGroupCategoriesResponse{}
	err := client.Request("cms", "ListMyGroupCategories", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListMyGroupCategoriesResponse struct {
	RequestId    goaliyun.String
	Success      bool
	ErrorCode    goaliyun.Integer
	ErrorMessage goaliyun.String
	Category     ListMyGroupCategoriesCategory
}

type ListMyGroupCategoriesCategory struct {
	GroupId       goaliyun.Integer
	CategoryItems ListMyGroupCategoriesCategoryItemList
}

type ListMyGroupCategoriesCategoryItem struct {
	Category goaliyun.String
	Count    goaliyun.Integer
}

type ListMyGroupCategoriesCategoryItemList []ListMyGroupCategoriesCategoryItem

func (list *ListMyGroupCategoriesCategoryItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListMyGroupCategoriesCategoryItem)
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
