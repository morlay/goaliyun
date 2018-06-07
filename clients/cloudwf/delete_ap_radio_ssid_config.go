package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type DeleteApRadioSsidConfigRequest struct {
	InstantlyEffective int64  `position:"Query" name:"InstantlyEffective"`
	Id                 int64  `position:"Query" name:"Id"`
	RegionId           string `position:"Query" name:"RegionId"`
}

func (req *DeleteApRadioSsidConfigRequest) Invoke(client goaliyun.Client) (*DeleteApRadioSsidConfigResponse, error) {
	resp := &DeleteApRadioSsidConfigResponse{}
	err := client.Request("cloudwf", "DeleteApRadioSsidConfig", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteApRadioSsidConfigResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
