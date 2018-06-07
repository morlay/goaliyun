package cms

import (
	"github.com/morlay/goaliyun"
)

type DeleteNotifyPolicyRequest struct {
	PolicyType string `position:"Query" name:"PolicyType"`
	AlertName  string `position:"Query" name:"AlertName"`
	Id         string `position:"Query" name:"Id"`
	Dimensions string `position:"Query" name:"Dimensions"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *DeleteNotifyPolicyRequest) Invoke(client goaliyun.Client) (*DeleteNotifyPolicyResponse, error) {
	resp := &DeleteNotifyPolicyResponse{}
	err := client.Request("cms", "DeleteNotifyPolicy", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteNotifyPolicyResponse struct {
	Code    goaliyun.String
	Message goaliyun.String
	Success goaliyun.String
	TraceId goaliyun.String
	Result  goaliyun.Integer
}
