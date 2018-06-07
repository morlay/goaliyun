package cms

import (
	"github.com/morlay/goaliyun"
)

type ListProductOfActiveAlertRequest struct {
	UserId   string `position:"Query" name:"UserId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *ListProductOfActiveAlertRequest) Invoke(client goaliyun.Client) (*ListProductOfActiveAlertResponse, error) {
	resp := &ListProductOfActiveAlertResponse{}
	err := client.Request("cms", "ListProductOfActiveAlert", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListProductOfActiveAlertResponse struct {
	RequestId  goaliyun.String
	Success    bool
	Code       goaliyun.Integer
	Message    goaliyun.String
	Datapoints goaliyun.String
}
