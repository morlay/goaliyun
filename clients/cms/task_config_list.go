package cms

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type TaskConfigListRequest struct {
	GroupId    int64  `position:"Query" name:"GroupId"`
	PageSize   int64  `position:"Query" name:"PageSize"`
	TaskName   string `position:"Query" name:"TaskName"`
	Id         int64  `position:"Query" name:"Id"`
	PageNumber int64  `position:"Query" name:"PageNumber"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *TaskConfigListRequest) Invoke(client goaliyun.Client) (*TaskConfigListResponse, error) {
	resp := &TaskConfigListResponse{}
	err := client.Request("cms", "TaskConfigList", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type TaskConfigListResponse struct {
	ErrorCode    goaliyun.Integer
	ErrorMessage goaliyun.String
	Success      bool
	RequestId    goaliyun.String
	PageNumber   goaliyun.Integer
	PageSize     goaliyun.Integer
	PageTotal    goaliyun.Integer
	Total        goaliyun.Integer
	TaskList     TaskConfigListNodeTaskConfigList
}

type TaskConfigListNodeTaskConfig struct {
	Id           goaliyun.Integer
	TaskName     goaliyun.String
	TaskType     goaliyun.String
	TaskScope    goaliyun.String
	Disabled     bool
	GroupId      goaliyun.Integer
	GroupName    goaliyun.String
	JsonData     goaliyun.String
	AlertConfig  goaliyun.String
	InstanceList TaskConfigListInstanceListList
}

type TaskConfigListNodeTaskConfigList []TaskConfigListNodeTaskConfig

func (list *TaskConfigListNodeTaskConfigList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]TaskConfigListNodeTaskConfig)
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

type TaskConfigListInstanceListList []goaliyun.String

func (list *TaskConfigListInstanceListList) UnmarshalJSON(data []byte) error {
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
