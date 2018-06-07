package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListFailureJobExecutionInstancesRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Count           int64  `position:"Query" name:"Count"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ListFailureJobExecutionInstancesRequest) Invoke(client goaliyun.Client) (*ListFailureJobExecutionInstancesResponse, error) {
	resp := &ListFailureJobExecutionInstancesResponse{}
	err := client.Request("emr", "ListFailureJobExecutionInstances", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListFailureJobExecutionInstancesResponse struct {
	RequestId    goaliyun.String
	JobInstances ListFailureJobExecutionInstancesJobInstanceList
}

type ListFailureJobExecutionInstancesJobInstance struct {
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

type ListFailureJobExecutionInstancesJobInstanceList []ListFailureJobExecutionInstancesJobInstance

func (list *ListFailureJobExecutionInstancesJobInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListFailureJobExecutionInstancesJobInstance)
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
