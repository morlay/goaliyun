package cms

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type TaskConfigUnhealthyRequest struct {
	TaskIdLists *TaskConfigUnhealthyTaskIdListList `position:"Query" type:"Repeated" name:"TaskIdList"`
	RegionId    string                             `position:"Query" name:"RegionId"`
}

func (req *TaskConfigUnhealthyRequest) Invoke(client goaliyun.Client) (*TaskConfigUnhealthyResponse, error) {
	resp := &TaskConfigUnhealthyResponse{}
	err := client.Request("cms", "TaskConfigUnhealthy", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type TaskConfigUnhealthyResponse struct {
	ErrorCode     goaliyun.Integer
	ErrorMessage  goaliyun.String
	Success       bool
	RequestId     goaliyun.String
	UnhealthyList TaskConfigUnhealthyNodeTaskInstanceList
}

type TaskConfigUnhealthyNodeTaskInstance struct {
	TaskId       goaliyun.Integer
	InstanceList TaskConfigUnhealthyInstanceListList
}

type TaskConfigUnhealthyTaskIdListList []int64

func (list *TaskConfigUnhealthyTaskIdListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]int64)
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

type TaskConfigUnhealthyNodeTaskInstanceList []TaskConfigUnhealthyNodeTaskInstance

func (list *TaskConfigUnhealthyNodeTaskInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]TaskConfigUnhealthyNodeTaskInstance)
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

type TaskConfigUnhealthyInstanceListList []goaliyun.String

func (list *TaskConfigUnhealthyInstanceListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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
