package emr

import (
	"github.com/morlay/goaliyun"
)

type CreateFlowCategoryRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Name            string `position:"Query" name:"Name"`
	Type            string `position:"Query" name:"Type"`
	ProjectId       string `position:"Query" name:"ProjectId"`
	ParentId        string `position:"Query" name:"ParentId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *CreateFlowCategoryRequest) Invoke(client goaliyun.Client) (*CreateFlowCategoryResponse, error) {
	resp := &CreateFlowCategoryResponse{}
	err := client.Request("emr", "CreateFlowCategory", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateFlowCategoryResponse struct {
	RequestId goaliyun.String
	Id        goaliyun.String
}
