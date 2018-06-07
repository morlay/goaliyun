package cms

import (
	"github.com/morlay/goaliyun"
)

type DeleteAlarmRequest struct {
	Id       string `position:"Query" name:"Id"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *DeleteAlarmRequest) Invoke(client goaliyun.Client) (*DeleteAlarmResponse, error) {
	resp := &DeleteAlarmResponse{}
	err := client.Request("cms", "DeleteAlarm", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteAlarmResponse struct {
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
}
