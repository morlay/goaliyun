package phoenixsp_inner

import (
	"github.com/morlay/goaliyun"
)

type InstanceLifecycleApiRequest struct {
	EventResult  string `position:"Query" name:"EventResult"`
	Async        string `position:"Query" name:"Async"`
	InstanceIds  string `position:"Query" name:"InstanceIds"`
	OrderId      string `position:"Query" name:"OrderId"`
	Success      string `position:"Query" name:"Success"`
	ExtraData    string `position:"Query" name:"ExtraData"`
	EventType    string `position:"Query" name:"EventType"`
	EventTime    int64  `position:"Query" name:"EventTime"`
	AliUid       int64  `position:"Query" name:"AliUid"`
	ResourceType string `position:"Query" name:"ResourceType"`
	EventSource  string `position:"Query" name:"EventSource"`
	Token        string `position:"Query" name:"Token"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *InstanceLifecycleApiRequest) Invoke(client goaliyun.Client) (*InstanceLifecycleApiResponse, error) {
	resp := &InstanceLifecycleApiResponse{}
	err := client.Request("phoenixsp-inner", "InstanceLifecycleApi", "2016-08-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type InstanceLifecycleApiResponse struct {
	RequestId goaliyun.String
	Code      goaliyun.String
	Message   goaliyun.String
	Success   bool
}
