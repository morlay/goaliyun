package emr

import (
	"github.com/morlay/goaliyun"
)

type DescribeFlowCategoryRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
	ProjectId       string `position:"Query" name:"ProjectId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribeFlowCategoryRequest) Invoke(client goaliyun.Client) (*DescribeFlowCategoryResponse, error) {
	resp := &DescribeFlowCategoryResponse{}
	err := client.Request("emr", "DescribeFlowCategory", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeFlowCategoryResponse struct {
	RequestId    goaliyun.String
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
