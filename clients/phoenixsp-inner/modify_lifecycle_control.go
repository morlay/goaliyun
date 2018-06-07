package phoenixsp_inner

import (
	"github.com/morlay/goaliyun"
)

type ModifyLifecycleControlRequest struct {
	InstanceIDs      string `position:"Query" name:"InstanceIDs"`
	AliUID           int64  `position:"Query" name:"AliUID"`
	Bid              string `position:"Query" name:"Bid"`
	ResourceType     string `position:"Query" name:"ResourceType"`
	Operator         string `position:"Query" name:"Operator"`
	ControlLifecycle string `position:"Query" name:"ControlLifecycle"`
	RegionId         string `position:"Query" name:"RegionId"`
}

func (req *ModifyLifecycleControlRequest) Invoke(client goaliyun.Client) (*ModifyLifecycleControlResponse, error) {
	resp := &ModifyLifecycleControlResponse{}
	err := client.Request("phoenixsp-inner", "ModifyLifecycleControl", "2016-08-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyLifecycleControlResponse struct {
	RequestId goaliyun.String
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
}
