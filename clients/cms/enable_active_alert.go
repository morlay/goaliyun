package cms

import (
	"github.com/morlay/goaliyun"
)

type EnableActiveAlertRequest struct {
	Product  string `position:"Query" name:"Product"`
	UserId   string `position:"Query" name:"UserId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *EnableActiveAlertRequest) Invoke(client goaliyun.Client) (*EnableActiveAlertResponse, error) {
	resp := &EnableActiveAlertResponse{}
	err := client.Request("cms", "EnableActiveAlert", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type EnableActiveAlertResponse struct {
	Success bool
	Code    goaliyun.String
	Message goaliyun.String
}
