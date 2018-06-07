package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListExecutionPlanInstancesRequest struct {
	OnlyLastInstance     string                                             `position:"Query" name:"OnlyLastInstance"`
	ResourceOwnerId      int64                                              `position:"Query" name:"ResourceOwnerId"`
	ExecutionPlanIdLists *ListExecutionPlanInstancesExecutionPlanIdListList `position:"Query" type:"Repeated" name:"ExecutionPlanIdList"`
	StatusLists          *ListExecutionPlanInstancesStatusListList          `position:"Query" type:"Repeated" name:"StatusList"`
	PageSize             int64                                              `position:"Query" name:"PageSize"`
	IsDesc               string                                             `position:"Query" name:"IsDesc"`
	PageNumber           int64                                              `position:"Query" name:"PageNumber"`
	RegionId             string                                             `position:"Query" name:"RegionId"`
}

func (req *ListExecutionPlanInstancesRequest) Invoke(client goaliyun.Client) (*ListExecutionPlanInstancesResponse, error) {
	resp := &ListExecutionPlanInstancesResponse{}
	err := client.Request("emr", "ListExecutionPlanInstances", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListExecutionPlanInstancesResponse struct {
	RequestId              goaliyun.String
	TotalCount             goaliyun.Integer
	PageNumber             goaliyun.Integer
	PageSize               goaliyun.Integer
	ExecutionPlanInstances ListExecutionPlanInstancesExecutionPlanInstanceList
}

type ListExecutionPlanInstancesExecutionPlanInstance struct {
	Id                goaliyun.String
	ExecutionPlanId   goaliyun.String
	ExecutionPlanName goaliyun.String
	StartTime         goaliyun.Integer
	RunTime           goaliyun.Integer
	ClusterId         goaliyun.String
	ClusterName       goaliyun.String
	ClusterType       goaliyun.String
	Status            goaliyun.String
	LogEnable         bool
	LogPath           goaliyun.String
	WorkflowApp       goaliyun.String
}

type ListExecutionPlanInstancesExecutionPlanIdListList []string

func (list *ListExecutionPlanInstancesExecutionPlanIdListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type ListExecutionPlanInstancesStatusListList []string

func (list *ListExecutionPlanInstancesStatusListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type ListExecutionPlanInstancesExecutionPlanInstanceList []ListExecutionPlanInstancesExecutionPlanInstance

func (list *ListExecutionPlanInstancesExecutionPlanInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListExecutionPlanInstancesExecutionPlanInstance)
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
