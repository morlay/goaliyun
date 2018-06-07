package cms

import (
	"github.com/morlay/goaliyun"
)

type CreateNotifyPolicyRequest struct {
	PolicyType string `position:"Query" name:"PolicyType"`
	AlertName  string `position:"Query" name:"AlertName"`
	EndTime    int64  `position:"Query" name:"EndTime"`
	StartTime  int64  `position:"Query" name:"StartTime"`
	Dimensions string `position:"Query" name:"Dimensions"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *CreateNotifyPolicyRequest) Invoke(client goaliyun.Client) (*CreateNotifyPolicyResponse, error) {
	resp := &CreateNotifyPolicyResponse{}
	err := client.Request("cms", "CreateNotifyPolicy", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateNotifyPolicyResponse struct {
	Code    goaliyun.String
	Message goaliyun.String
	Success goaliyun.String
	TraceId goaliyun.String
	Result  goaliyun.Integer
}
