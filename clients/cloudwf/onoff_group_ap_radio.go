package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type OnoffGroupApRadioRequest struct {
	ApgroupId int64  `position:"Query" name:"ApgroupId"`
	Disabled  int64  `position:"Query" name:"Disabled"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *OnoffGroupApRadioRequest) Invoke(client goaliyun.Client) (*OnoffGroupApRadioResponse, error) {
	resp := &OnoffGroupApRadioResponse{}
	err := client.Request("cloudwf", "OnoffGroupApRadio", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OnoffGroupApRadioResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
