package ram

import (
	"github.com/morlay/goaliyun"
)

type DeletePolicyRequest struct {
	PolicyName string `position:"Query" name:"PolicyName"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *DeletePolicyRequest) Invoke(client goaliyun.Client) (*DeletePolicyResponse, error) {
	resp := &DeletePolicyResponse{}
	err := client.Request("ram", "DeletePolicy", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeletePolicyResponse struct {
	RequestId goaliyun.String
}
