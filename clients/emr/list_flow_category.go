package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListFlowCategoryRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Root            string `position:"Query" name:"Root"`
	PageSize        int64  `position:"Query" name:"PageSize"`
	ProjectId       string `position:"Query" name:"ProjectId"`
	ParentId        string `position:"Query" name:"ParentId"`
	PageNumber      int64  `position:"Query" name:"PageNumber"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ListFlowCategoryRequest) Invoke(client goaliyun.Client) (*ListFlowCategoryResponse, error) {
	resp := &ListFlowCategoryResponse{}
	err := client.Request("emr", "ListFlowCategory", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListFlowCategoryResponse struct {
	RequestId  goaliyun.String
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	Total      goaliyun.Integer
	Categories ListFlowCategoryCategoryList
}

type ListFlowCategoryCategory struct {
	Id           goaliyun.String
	GmtCreate    goaliyun.Integer
	GmtModified  goaliyun.Integer
	Name         goaliyun.String
	ParentId     goaliyun.String
	Type         goaliyun.String
	CategoryType goaliyun.String
	ObjectType   goaliyun.String
	ObjectId     goaliyun.String
	ProjectId    goaliyun.String
}

type ListFlowCategoryCategoryList []ListFlowCategoryCategory

func (list *ListFlowCategoryCategoryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListFlowCategoryCategory)
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
