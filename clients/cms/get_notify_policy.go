package cms

import (
	"github.com/morlay/goaliyun"
)

type GetNotifyPolicyRequest struct {
	PolicyType string `position:"Query" name:"PolicyType"`
	AlertName  string `position:"Query" name:"AlertName"`
	Id         string `position:"Query" name:"Id"`
	Dimensions string `position:"Query" name:"Dimensions"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *GetNotifyPolicyRequest) Invoke(client goaliyun.Client) (*GetNotifyPolicyResponse, error) {
	resp := &GetNotifyPolicyResponse{}
	err := client.Request("cms", "GetNotifyPolicy", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetNotifyPolicyResponse struct {
	Code    goaliyun.String
	Message goaliyun.String
	Success goaliyun.String
	TraceId goaliyun.String
	Result  GetNotifyPolicyResult
}

type GetNotifyPolicyResult struct {
	AlertName  goaliyun.String
	Dimensions goaliyun.String
	Type       goaliyun.String
	Id         goaliyun.String
	StartTime  goaliyun.Integer
	EndTime    goaliyun.Integer
}
