package ccc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ResumeJobsRequest struct {
	All             string                        `position:"Query" name:"All"`
	JobIds          *ResumeJobsJobIdList          `position:"Query" type:"Repeated" name:"JobId"`
	InstanceId      string                        `position:"Query" name:"InstanceId"`
	JobReferenceIds *ResumeJobsJobReferenceIdList `position:"Query" type:"Repeated" name:"JobReferenceId"`
	GroupId         string                        `position:"Query" name:"GroupId"`
	ScenarioId      string                        `position:"Query" name:"ScenarioId"`
	RegionId        string                        `position:"Query" name:"RegionId"`
}

func (req *ResumeJobsRequest) Invoke(client goaliyun.Client) (*ResumeJobsResponse, error) {
	resp := &ResumeJobsResponse{}
	err := client.Request("ccc", "ResumeJobs", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ResumeJobsResponse struct {
	RequestId      goaliyun.String
	Success        bool
	Code           goaliyun.String
	Message        goaliyun.String
	HttpStatusCode goaliyun.Integer
}

type ResumeJobsJobIdList []string

func (list *ResumeJobsJobIdList) UnmarshalJSON(data []byte) error {
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

type ResumeJobsJobReferenceIdList []string

func (list *ResumeJobsJobReferenceIdList) UnmarshalJSON(data []byte) error {
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
