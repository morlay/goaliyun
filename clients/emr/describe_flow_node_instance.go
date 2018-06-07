package emr

import (
	"github.com/morlay/goaliyun"
)

type DescribeFlowNodeInstanceRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
	ProjectId       string `position:"Query" name:"ProjectId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribeFlowNodeInstanceRequest) Invoke(client goaliyun.Client) (*DescribeFlowNodeInstanceResponse, error) {
	resp := &DescribeFlowNodeInstanceResponse{}
	err := client.Request("emr", "DescribeFlowNodeInstance", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeFlowNodeInstanceResponse struct {
	RequestId      goaliyun.String
	Id             goaliyun.String
	GmtCreate      goaliyun.Integer
	GmtModified    goaliyun.Integer
	Type           goaliyun.String
	Status         goaliyun.String
	JobId          goaliyun.String
	JobName        goaliyun.String
	JobType        goaliyun.String
	JobParams      goaliyun.String
	FailAct        goaliyun.String
	MaxRetry       goaliyun.String
	RetryInterval  goaliyun.String
	NodeName       goaliyun.String
	FlowId         goaliyun.String
	ClusterId      goaliyun.String
	HostName       goaliyun.String
	ProjectId      goaliyun.String
	StartTime      goaliyun.Integer
	EndTime        goaliyun.Integer
	Retries        goaliyun.Integer
	ExternalId     goaliyun.String
	ExternalStatus goaliyun.String
	ExternalInfo   goaliyun.String
	ParamConf      goaliyun.String
	EnvConf        goaliyun.String
	RunConf        goaliyun.String
	ClusterName    goaliyun.String
}
