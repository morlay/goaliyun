package phoenixsp_inner

import (
	"github.com/morlay/goaliyun"
)

type ModifyPredictedEndTimeRequest struct {
	Caller       string `position:"Query" name:"Caller"`
	InstanceName string `position:"Query" name:"InstanceName"`
	EndTime      int64  `position:"Query" name:"EndTime"`
	AliUID       int64  `position:"Query" name:"AliUID"`
	Bid          string `position:"Query" name:"Bid"`
	ResourceType string `position:"Query" name:"ResourceType"`
	Operator     string `position:"Query" name:"Operator"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *ModifyPredictedEndTimeRequest) Invoke(client goaliyun.Client) (*ModifyPredictedEndTimeResponse, error) {
	resp := &ModifyPredictedEndTimeResponse{}
	err := client.Request("phoenixsp-inner", "ModifyPredictedEndTime", "2016-08-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyPredictedEndTimeResponse struct {
	RequestId goaliyun.String
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
}
