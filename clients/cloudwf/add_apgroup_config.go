package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type AddApgroupConfigRequest struct {
	ParentApgroupId int64  `position:"Query" name:"ParentApgroupId"`
	Name            string `position:"Query" name:"Name"`
	Description     string `position:"Query" name:"Description"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *AddApgroupConfigRequest) Invoke(client goaliyun.Client) (*AddApgroupConfigResponse, error) {
	resp := &AddApgroupConfigResponse{}
	err := client.Request("cloudwf", "AddApgroupConfig", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddApgroupConfigResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
