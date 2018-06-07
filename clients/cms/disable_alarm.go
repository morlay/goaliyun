package cms

import (
	"github.com/morlay/goaliyun"
)

type DisableAlarmRequest struct {
	Id       string `position:"Query" name:"Id"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *DisableAlarmRequest) Invoke(client goaliyun.Client) (*DisableAlarmResponse, error) {
	resp := &DisableAlarmResponse{}
	err := client.Request("cms", "DisableAlarm", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DisableAlarmResponse struct {
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
}
