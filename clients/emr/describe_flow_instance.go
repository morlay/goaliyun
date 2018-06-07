package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeFlowInstanceRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
	ProjectId       string `position:"Query" name:"ProjectId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribeFlowInstanceRequest) Invoke(client goaliyun.Client) (*DescribeFlowInstanceResponse, error) {
	resp := &DescribeFlowInstanceResponse{}
	err := client.Request("emr", "DescribeFlowInstance", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeFlowInstanceResponse struct {
	RequestId    goaliyun.String
	Id           goaliyun.String
	GmtCreate    goaliyun.Integer
	GmtModified  goaliyun.Integer
	FlowId       goaliyun.String
	FlowName     goaliyun.String
	ProjectId    goaliyun.String
	Status       goaliyun.String
	ClusterId    goaliyun.String
	StartTime    goaliyun.Integer
	EndTime      goaliyun.Integer
	NodeInstance DescribeFlowInstanceNodeInstanceItemList
}

type DescribeFlowInstanceNodeInstanceItem struct {
	Id             goaliyun.String
	GmtCreate      goaliyun.Integer
	GmtModified    goaliyun.Integer
	Type           goaliyun.String
	Status         goaliyun.String
	JobId          goaliyun.String
	JobName        goaliyun.String
	JobType        goaliyun.String
	FailAct        goaliyun.String
	MaxRetry       goaliyun.String
	RetryInterval  goaliyun.String
	NodeName       goaliyun.String
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
}

type DescribeFlowInstanceNodeInstanceItemList []DescribeFlowInstanceNodeInstanceItem

func (list *DescribeFlowInstanceNodeInstanceItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeFlowInstanceNodeInstanceItem)
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
