package ram

import (
	"github.com/morlay/goaliyun"
)

type SetDefaultPolicyVersionRequest struct {
	VersionId  string `position:"Query" name:"VersionId"`
	PolicyName string `position:"Query" name:"PolicyName"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *SetDefaultPolicyVersionRequest) Invoke(client goaliyun.Client) (*SetDefaultPolicyVersionResponse, error) {
	resp := &SetDefaultPolicyVersionResponse{}
	err := client.Request("ram", "SetDefaultPolicyVersion", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetDefaultPolicyVersionResponse struct {
	RequestId goaliyun.String
}
