package ram

import (
	"github.com/morlay/goaliyun"
)

type DetachPolicyFromGroupRequest struct {
	PolicyType string `position:"Query" name:"PolicyType"`
	PolicyName string `position:"Query" name:"PolicyName"`
	GroupName  string `position:"Query" name:"GroupName"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *DetachPolicyFromGroupRequest) Invoke(client goaliyun.Client) (*DetachPolicyFromGroupResponse, error) {
	resp := &DetachPolicyFromGroupResponse{}
	err := client.Request("ram", "DetachPolicyFromGroup", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DetachPolicyFromGroupResponse struct {
	RequestId goaliyun.String
}
