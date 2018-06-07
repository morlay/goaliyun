package market

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryMarketCategoriesRequest struct {
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *QueryMarketCategoriesRequest) Invoke(client goaliyun.Client) (*QueryMarketCategoriesResponse, error) {
	resp := &QueryMarketCategoriesResponse{}
	err := client.Request("market", "QueryMarketCategories", "2015-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryMarketCategoriesResponse struct {
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	TotalCount goaliyun.Integer
	RequestId  goaliyun.String
	Categories QueryMarketCategoriesCategoryList
}

type QueryMarketCategoriesCategory struct {
	Id              goaliyun.Integer
	CategoryCode    goaliyun.String
	CategoryName    goaliyun.String
	ChildCategories QueryMarketCategoriesChildCategoryList
}

type QueryMarketCategoriesChildCategory struct {
	Id           goaliyun.Integer
	CategoryCode goaliyun.String
	CategoryName goaliyun.String
}

type QueryMarketCategoriesCategoryList []QueryMarketCategoriesCategory

func (list *QueryMarketCategoriesCategoryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMarketCategoriesCategory)
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

type QueryMarketCategoriesChildCategoryList []QueryMarketCategoriesChildCategory

func (list *QueryMarketCategoriesChildCategoryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMarketCategoriesChildCategory)
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
