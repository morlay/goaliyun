package cms

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type TaskConfigEnableRequest struct {
	IdLists  *TaskConfigEnableIdListList `position:"Query" type:"Repeated" name:"IdList"`
	Enabled  string                      `position:"Query" name:"Enabled"`
	RegionId string                      `position:"Query" name:"RegionId"`
}

func (req *TaskConfigEnableRequest) Invoke(client goaliyun.Client) (*TaskConfigEnableResponse, error) {
	resp := &TaskConfigEnableResponse{}
	err := client.Request("cms", "TaskConfigEnable", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type TaskConfigEnableResponse struct {
	ErrorCode    goaliyun.Integer
	ErrorMessage goaliyun.String
	Success      bool
	RequestId    goaliyun.String
}

type TaskConfigEnableIdListList []int64

func (list *TaskConfigEnableIdListList) UnmarshalJSON(data []byte) error {
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
