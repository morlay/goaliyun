package ram

import (
	"github.com/morlay/goaliyun"
)

type DetachPolicyFromUserRequest struct {
	PolicyType string `position:"Query" name:"PolicyType"`
	PolicyName string `position:"Query" name:"PolicyName"`
	UserName   string `position:"Query" name:"UserName"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *DetachPolicyFromUserRequest) Invoke(client goaliyun.Client) (*DetachPolicyFromUserResponse, error) {
	resp := &DetachPolicyFromUserResponse{}
	err := client.Request("ram", "DetachPolicyFromUser", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DetachPolicyFromUserResponse struct {
	RequestId goaliyun.String
}
