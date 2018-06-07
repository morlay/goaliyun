package emr

import (
	"github.com/morlay/goaliyun"
)

type DescribeFlowCategoryTreeRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Type            string `position:"Query" name:"Type"`
	ProjectId       string `position:"Query" name:"ProjectId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribeFlowCategoryTreeRequest) Invoke(client goaliyun.Client) (*DescribeFlowCategoryTreeResponse, error) {
	resp := &DescribeFlowCategoryTreeResponse{}
	err := client.Request("emr", "DescribeFlowCategoryTree", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeFlowCategoryTreeResponse struct {
	RequestId goaliyun.String
	Data      goaliyun.String
}
