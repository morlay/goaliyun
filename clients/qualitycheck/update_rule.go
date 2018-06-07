package qualitycheck

import (
	"github.com/morlay/goaliyun"
)

type UpdateRuleRequest struct {
	JsonStr  string `position:"Query" name:"JsonStr"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *UpdateRuleRequest) Invoke(client goaliyun.Client) (*UpdateRuleResponse, error) {
	resp := &UpdateRuleResponse{}
	err := client.Request("qualitycheck", "UpdateRule", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateRuleResponse struct {
	RequestId goaliyun.String
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
	Data      goaliyun.String
}
