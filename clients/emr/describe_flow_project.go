package emr

import (
	"github.com/morlay/goaliyun"
)

type DescribeFlowProjectRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ProjectId       string `position:"Query" name:"ProjectId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribeFlowProjectRequest) Invoke(client goaliyun.Client) (*DescribeFlowProjectResponse, error) {
	resp := &DescribeFlowProjectResponse{}
	err := client.Request("emr", "DescribeFlowProject", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeFlowProjectResponse struct {
	RequestId   goaliyun.String
	Id          goaliyun.String
	GmtCreate   goaliyun.Integer
	GmtModified goaliyun.Integer
	UserId      goaliyun.String
	Name        goaliyun.String
	Description goaliyun.String
}
