package ram

import (
	"github.com/morlay/goaliyun"
)

type DeletePolicyVersionRequest struct {
	VersionId  string `position:"Query" name:"VersionId"`
	PolicyName string `position:"Query" name:"PolicyName"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *DeletePolicyVersionRequest) Invoke(client goaliyun.Client) (*DeletePolicyVersionResponse, error) {
	resp := &DeletePolicyVersionResponse{}
	err := client.Request("ram", "DeletePolicyVersion", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeletePolicyVersionResponse struct {
	RequestId goaliyun.String
}
