package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type InvokeCommandRequest struct {
	ResourceOwnerId      int64                        `position:"Query" name:"ResourceOwnerId"`
	CommandId            string                       `position:"Query" name:"CommandId"`
	Frequency            string                       `position:"Query" name:"Frequency"`
	Timed                string                       `position:"Query" name:"Timed"`
	ResourceOwnerAccount string                       `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                       `position:"Query" name:"OwnerAccount"`
	OwnerId              int64                        `position:"Query" name:"OwnerId"`
	InstanceIds          *InvokeCommandInstanceIdList `position:"Query" type:"Repeated" name:"InstanceId"`
	RegionId             string                       `position:"Query" name:"RegionId"`
}

func (req *InvokeCommandRequest) Invoke(client goaliyun.Client) (*InvokeCommandResponse, error) {
	resp := &InvokeCommandResponse{}
	err := client.Request("ecs", "InvokeCommand", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type InvokeCommandResponse struct {
	RequestId goaliyun.String
	InvokeId  goaliyun.String
}

type InvokeCommandInstanceIdList []string

func (list *InvokeCommandInstanceIdList) UnmarshalJSON(data []byte) error {
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
