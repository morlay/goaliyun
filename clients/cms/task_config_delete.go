package cms

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type TaskConfigDeleteRequest struct {
	IdLists  *TaskConfigDeleteIdListList `position:"Query" type:"Repeated" name:"IdList"`
	RegionId string                      `position:"Query" name:"RegionId"`
}

func (req *TaskConfigDeleteRequest) Invoke(client goaliyun.Client) (*TaskConfigDeleteResponse, error) {
	resp := &TaskConfigDeleteResponse{}
	err := client.Request("cms", "TaskConfigDelete", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type TaskConfigDeleteResponse struct {
	ErrorCode    goaliyun.Integer
	ErrorMessage goaliyun.String
	Success      bool
	RequestId    goaliyun.String
}

type TaskConfigDeleteIdListList []int64

func (list *TaskConfigDeleteIdListList) UnmarshalJSON(data []byte) error {
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
