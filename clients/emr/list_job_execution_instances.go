package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListJobExecutionInstancesRequest struct {
	ResourceOwnerId         int64  `position:"Query" name:"ResourceOwnerId"`
	ExecutionPlanInstanceId string `position:"Query" name:"ExecutionPlanInstanceId"`
	PageSize                int64  `position:"Query" name:"PageSize"`
	IsDesc                  string `position:"Query" name:"IsDesc"`
	PageNumber              int64  `position:"Query" name:"PageNumber"`
	RegionId                string `position:"Query" name:"RegionId"`
}

func (req *ListJobExecutionInstancesRequest) Invoke(client goaliyun.Client) (*ListJobExecutionInstancesResponse, error) {
	resp := &ListJobExecutionInstancesResponse{}
	err := client.Request("emr", "ListJobExecutionInstances", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListJobExecutionInstancesResponse struct {
	RequestId    goaliyun.String
	TotalCount   goaliyun.Integer
	PageNumber   goaliyun.Integer
	PageSize     goaliyun.Integer
	JobInstances ListJobExecutionInstancesJobInstanceList
}

type ListJobExecutionInstancesJobInstance struct {
	Id        goaliyun.String
	JobName   goaliyun.String
	StartTime goaliyun.Integer
	RunTime   goaliyun.Integer
	JobType   goaliyun.String
	JobId     goaliyun.String
	ClusterId goaliyun.String
	Status    goaliyun.String
	RetryInfo goaliyun.String
}

type ListJobExecutionInstancesJobInstanceList []ListJobExecutionInstancesJobInstance

func (list *ListJobExecutionInstancesJobInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobExecutionInstancesJobInstance)
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
