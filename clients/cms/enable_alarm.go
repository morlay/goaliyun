package cms

import (
	"github.com/morlay/goaliyun"
)

type EnableAlarmRequest struct {
	Id       string `position:"Query" name:"Id"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *EnableAlarmRequest) Invoke(client goaliyun.Client) (*EnableAlarmResponse, error) {
	resp := &EnableAlarmResponse{}
	err := client.Request("cms", "EnableAlarm", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type EnableAlarmResponse struct {
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
}
