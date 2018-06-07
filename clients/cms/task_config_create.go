package cms

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type TaskConfigCreateRequest struct {
	InstanceLists *TaskConfigCreateInstanceListList `position:"Query" type:"Repeated" name:"InstanceList"`
	JsonData      string                            `position:"Query" name:"JsonData"`
	TaskType      string                            `position:"Query" name:"TaskType"`
	TaskScope     string                            `position:"Query" name:"TaskScope"`
	AlertConfig   string                            `position:"Query" name:"AlertConfig"`
	GroupId       int64                             `position:"Query" name:"GroupId"`
	TaskName      string                            `position:"Query" name:"TaskName"`
	GroupName     string                            `position:"Query" name:"GroupName"`
	RegionId      string                            `position:"Query" name:"RegionId"`
}

func (req *TaskConfigCreateRequest) Invoke(client goaliyun.Client) (*TaskConfigCreateResponse, error) {
	resp := &TaskConfigCreateResponse{}
	err := client.Request("cms", "TaskConfigCreate", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type TaskConfigCreateResponse struct {
	ErrorCode    goaliyun.Integer
	ErrorMessage goaliyun.String
	Success      bool
	RequestId    goaliyun.String
	TaskId       goaliyun.Integer
}

type TaskConfigCreateInstanceListList []string

func (list *TaskConfigCreateInstanceListList) UnmarshalJSON(data []byte) error {
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
