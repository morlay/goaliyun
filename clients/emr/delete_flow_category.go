package emr

import (
	"github.com/morlay/goaliyun"
)

type DeleteFlowCategoryRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
	ProjectId       string `position:"Query" name:"ProjectId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DeleteFlowCategoryRequest) Invoke(client goaliyun.Client) (*DeleteFlowCategoryResponse, error) {
	resp := &DeleteFlowCategoryResponse{}
	err := client.Request("emr", "DeleteFlowCategory", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteFlowCategoryResponse struct {
	RequestId goaliyun.String
	Data      bool
}
