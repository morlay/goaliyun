package cms

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type TaskConfigModifyRequest struct {
	InstanceLists *TaskConfigModifyInstanceListList `position:"Query" type:"Repeated" name:"InstanceList"`
	JsonData      string                            `position:"Query" name:"JsonData"`
	TaskType      string                            `position:"Query" name:"TaskType"`
	TaskScope     string                            `position:"Query" name:"TaskScope"`
	AlertConfig   string                            `position:"Query" name:"AlertConfig"`
	GroupId       int64                             `position:"Query" name:"GroupId"`
	TaskName      string                            `position:"Query" name:"TaskName"`
	Id            int64                             `position:"Query" name:"Id"`
	GroupName     string                            `position:"Query" name:"GroupName"`
	RegionId      string                            `position:"Query" name:"RegionId"`
}

func (req *TaskConfigModifyRequest) Invoke(client goaliyun.Client) (*TaskConfigModifyResponse, error) {
	resp := &TaskConfigModifyResponse{}
	err := client.Request("cms", "TaskConfigModify", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type TaskConfigModifyResponse struct {
	ErrorCode    goaliyun.Integer
	ErrorMessage goaliyun.String
	Success      bool
	RequestId    goaliyun.String
}

type TaskConfigModifyInstanceListList []string

func (list *TaskConfigModifyInstanceListList) UnmarshalJSON(data []byte) error {
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
