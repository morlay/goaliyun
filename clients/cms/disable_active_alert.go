package cms

import (
	"github.com/morlay/goaliyun"
)

type DisableActiveAlertRequest struct {
	Product  string `position:"Query" name:"Product"`
	UserId   string `position:"Query" name:"UserId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *DisableActiveAlertRequest) Invoke(client goaliyun.Client) (*DisableActiveAlertResponse, error) {
	resp := &DisableActiveAlertResponse{}
	err := client.Request("cms", "DisableActiveAlert", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DisableActiveAlertResponse struct {
	RequestId goaliyun.String
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
}
