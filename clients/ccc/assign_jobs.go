package ccc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type AssignJobsRequest struct {
	CallingNumbers *AssignJobsCallingNumberList `position:"Query" type:"Repeated" name:"CallingNumber"`
	InstanceId     string                       `position:"Query" name:"InstanceId"`
	GroupId        string                       `position:"Query" name:"GroupId"`
	StrategyJson   string                       `position:"Query" name:"StrategyJson"`
	ScenarioId     string                       `position:"Query" name:"ScenarioId"`
	JobsJsons      *AssignJobsJobsJsonList      `position:"Query" type:"Repeated" name:"JobsJson"`
	RegionId       string                       `position:"Query" name:"RegionId"`
}

func (req *AssignJobsRequest) Invoke(client goaliyun.Client) (*AssignJobsResponse, error) {
	resp := &AssignJobsResponse{}
	err := client.Request("ccc", "AssignJobs", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AssignJobsResponse struct {
	RequestId      goaliyun.String
	Success        bool
	Code           goaliyun.String
	Message        goaliyun.String
	HttpStatusCode goaliyun.Integer
	GroupId        goaliyun.String
}

type AssignJobsCallingNumberList []string

func (list *AssignJobsCallingNumberList) UnmarshalJSON(data []byte) error {
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

type AssignJobsJobsJsonList []string

func (list *AssignJobsJobsJsonList) UnmarshalJSON(data []byte) error {
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
