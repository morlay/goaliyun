package emr

import (
	"github.com/morlay/goaliyun"
)

type ModifyFlowCategoryRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Name            string `position:"Query" name:"Name"`
	Id              string `position:"Query" name:"Id"`
	ProjectId       string `position:"Query" name:"ProjectId"`
	ParentId        string `position:"Query" name:"ParentId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ModifyFlowCategoryRequest) Invoke(client goaliyun.Client) (*ModifyFlowCategoryResponse, error) {
	resp := &ModifyFlowCategoryResponse{}
	err := client.Request("emr", "ModifyFlowCategory", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyFlowCategoryResponse struct {
	RequestId goaliyun.String
	Data      bool
}
