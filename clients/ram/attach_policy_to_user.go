package ram

import (
	"github.com/morlay/goaliyun"
)

type AttachPolicyToUserRequest struct {
	PolicyType string `position:"Query" name:"PolicyType"`
	PolicyName string `position:"Query" name:"PolicyName"`
	UserName   string `position:"Query" name:"UserName"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *AttachPolicyToUserRequest) Invoke(client goaliyun.Client) (*AttachPolicyToUserResponse, error) {
	resp := &AttachPolicyToUserResponse{}
	err := client.Request("ram", "AttachPolicyToUser", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AttachPolicyToUserResponse struct {
	RequestId goaliyun.String
}
