package qualitycheck

import (
	"github.com/morlay/goaliyun"
)

type InvalidRuleRequest struct {
	JsonStr  string `position:"Query" name:"JsonStr"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *InvalidRuleRequest) Invoke(client goaliyun.Client) (*InvalidRuleResponse, error) {
	resp := &InvalidRuleResponse{}
	err := client.Request("qualitycheck", "InvalidRule", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type InvalidRuleResponse struct {
	RequestId goaliyun.String
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
	Data      bool
}
