package emr

import (
	"github.com/morlay/goaliyun"
)

type DescribeFlowRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
	ProjectId       string `position:"Query" name:"ProjectId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribeFlowRequest) Invoke(client goaliyun.Client) (*DescribeFlowResponse, error) {
	resp := &DescribeFlowResponse{}
	err := client.Request("emr", "DescribeFlow", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeFlowResponse struct {
	RequestId     goaliyun.String
	Id            goaliyun.String
	GmtCreate     goaliyun.Integer
	GmtModified   goaliyun.Integer
	Name          goaliyun.String
	Description   goaliyun.String
	Type          goaliyun.String
	Status        goaliyun.String
	Periodic      bool
	StartSchedule goaliyun.Integer
	EndSchedule   goaliyun.Integer
	CronExpr      goaliyun.String
	CreateCluster bool
	ClusterId     goaliyun.String
	Graph         goaliyun.String
	CategoryId    goaliyun.String
}
