package domain

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type AcknowledgeTaskResultRequest struct {
	TaskDetailNos *AcknowledgeTaskResultTaskDetailNoList `position:"Query" type:"Repeated" name:"TaskDetailNo"`
	UserClientIp  string                                 `position:"Query" name:"UserClientIp"`
	Lang          string                                 `position:"Query" name:"Lang"`
	RegionId      string                                 `position:"Query" name:"RegionId"`
}

func (req *AcknowledgeTaskResultRequest) Invoke(client goaliyun.Client) (*AcknowledgeTaskResultResponse, error) {
	resp := &AcknowledgeTaskResultResponse{}
	err := client.Request("domain", "AcknowledgeTaskResult", "2018-01-29", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AcknowledgeTaskResultResponse struct {
	RequestId goaliyun.String
	Result    goaliyun.Integer
}

type AcknowledgeTaskResultTaskDetailNoList []string

func (list *AcknowledgeTaskResultTaskDetailNoList) UnmarshalJSON(data []byte) error {
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
