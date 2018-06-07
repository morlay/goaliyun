package cms

import (
	"github.com/morlay/goaliyun"
)

type PutEventRequest struct {
	EventInfo string `position:"Query" name:"EventInfo"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *PutEventRequest) Invoke(client goaliyun.Client) (*PutEventResponse, error) {
	resp := &PutEventResponse{}
	err := client.Request("cms", "PutEvent", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type PutEventResponse struct {
	RequestId goaliyun.String
	Code      goaliyun.String
	Message   goaliyun.String
	Data      goaliyun.String
}
