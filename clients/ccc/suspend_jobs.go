package ccc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type SuspendJobsRequest struct {
	All             string                         `position:"Query" name:"All"`
	JobIds          *SuspendJobsJobIdList          `position:"Query" type:"Repeated" name:"JobId"`
	InstanceId      string                         `position:"Query" name:"InstanceId"`
	JobReferenceIds *SuspendJobsJobReferenceIdList `position:"Query" type:"Repeated" name:"JobReferenceId"`
	GroupId         string                         `position:"Query" name:"GroupId"`
	ScenarioId      string                         `position:"Query" name:"ScenarioId"`
	RegionId        string                         `position:"Query" name:"RegionId"`
}

func (req *SuspendJobsRequest) Invoke(client goaliyun.Client) (*SuspendJobsResponse, error) {
	resp := &SuspendJobsResponse{}
	err := client.Request("ccc", "SuspendJobs", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SuspendJobsResponse struct {
	RequestId      goaliyun.String
	Success        bool
	Code           goaliyun.String
	Message        goaliyun.String
	HttpStatusCode goaliyun.Integer
}

type SuspendJobsJobIdList []string

func (list *SuspendJobsJobIdList) UnmarshalJSON(data []byte) error {
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

type SuspendJobsJobReferenceIdList []string

func (list *SuspendJobsJobReferenceIdList) UnmarshalJSON(data []byte) error {
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
