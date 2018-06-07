package ccc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type CancelJobsRequest struct {
	All             string                        `position:"Query" name:"All"`
	JobIds          *CancelJobsJobIdList          `position:"Query" type:"Repeated" name:"JobId"`
	InstanceId      string                        `position:"Query" name:"InstanceId"`
	JobReferenceIds *CancelJobsJobReferenceIdList `position:"Query" type:"Repeated" name:"JobReferenceId"`
	GroupId         string                        `position:"Query" name:"GroupId"`
	ScenarioId      string                        `position:"Query" name:"ScenarioId"`
	RegionId        string                        `position:"Query" name:"RegionId"`
}

func (req *CancelJobsRequest) Invoke(client goaliyun.Client) (*CancelJobsResponse, error) {
	resp := &CancelJobsResponse{}
	err := client.Request("ccc", "CancelJobs", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CancelJobsResponse struct {
	RequestId      goaliyun.String
	Success        bool
	Code           goaliyun.String
	Message        goaliyun.String
	HttpStatusCode goaliyun.Integer
}

type CancelJobsJobIdList []string

func (list *CancelJobsJobIdList) UnmarshalJSON(data []byte) error {
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

type CancelJobsJobReferenceIdList []string

func (list *CancelJobsJobReferenceIdList) UnmarshalJSON(data []byte) error {
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
