package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type SaveApPortalConfigRequest struct {
	AuthKey      string `position:"Query" name:"AuthKey"`
	PortalUrl    string `position:"Query" name:"PortalUrl"`
	PortalStatus string `position:"Query" name:"PortalStatus"`
	Whitelist    string `position:"Query" name:"Whitelist"`
	CheckUrl     string `position:"Query" name:"CheckUrl"`
	ApConfigId   int64  `position:"Query" name:"ApConfigId"`
	AuthSecret   string `position:"Query" name:"AuthSecret"`
	WebAuthUrl   string `position:"Query" name:"WebAuthUrl"`
	Network      int64  `position:"Query" name:"Network"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *SaveApPortalConfigRequest) Invoke(client goaliyun.Client) (*SaveApPortalConfigResponse, error) {
	resp := &SaveApPortalConfigResponse{}
	err := client.Request("cloudwf", "SaveApPortalConfig", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SaveApPortalConfigResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
