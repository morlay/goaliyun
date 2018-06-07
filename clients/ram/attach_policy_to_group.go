package ram

import (
	"github.com/morlay/goaliyun"
)

type AttachPolicyToGroupRequest struct {
	PolicyType string `position:"Query" name:"PolicyType"`
	PolicyName string `position:"Query" name:"PolicyName"`
	GroupName  string `position:"Query" name:"GroupName"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *AttachPolicyToGroupRequest) Invoke(client goaliyun.Client) (*AttachPolicyToGroupResponse, error) {
	resp := &AttachPolicyToGroupResponse{}
	err := client.Request("ram", "AttachPolicyToGroup", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AttachPolicyToGroupResponse struct {
	RequestId goaliyun.String
}
