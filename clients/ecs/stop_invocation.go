package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type StopInvocationRequest struct {
	ResourceOwnerId      int64                         `position:"Query" name:"ResourceOwnerId"`
	InvokeId             string                        `position:"Query" name:"InvokeId"`
	ResourceOwnerAccount string                        `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                        `position:"Query" name:"OwnerAccount"`
	OwnerId              int64                         `position:"Query" name:"OwnerId"`
	InstanceIds          *StopInvocationInstanceIdList `position:"Query" type:"Repeated" name:"InstanceId"`
	RegionId             string                        `position:"Query" name:"RegionId"`
}

func (req *StopInvocationRequest) Invoke(client goaliyun.Client) (*StopInvocationResponse, error) {
	resp := &StopInvocationResponse{}
	err := client.Request("ecs", "StopInvocation", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type StopInvocationResponse struct {
	RequestId goaliyun.String
}

type StopInvocationInstanceIdList []string

func (list *StopInvocationInstanceIdList) UnmarshalJSON(data []byte) error {
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
