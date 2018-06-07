package ehpc

import (
	"github.com/morlay/goaliyun"
)

type ModifyClusterAttributesRequest struct {
	Name        string `position:"Query" name:"Name"`
	Description string `position:"Query" name:"Description"`
	ClusterId   string `position:"Query" name:"ClusterId"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *ModifyClusterAttributesRequest) Invoke(client goaliyun.Client) (*ModifyClusterAttributesResponse, error) {
	resp := &ModifyClusterAttributesResponse{}
	err := client.Request("ehpc", "ModifyClusterAttributes", "2018-04-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyClusterAttributesResponse struct {
	RequestId goaliyun.String
}
