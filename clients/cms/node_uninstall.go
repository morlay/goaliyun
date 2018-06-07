package cms

import (
	"github.com/morlay/goaliyun"
)

type NodeUninstallRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *NodeUninstallRequest) Invoke(client goaliyun.Client) (*NodeUninstallResponse, error) {
	resp := &NodeUninstallResponse{}
	err := client.Request("cms", "NodeUninstall", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type NodeUninstallResponse struct {
	ErrorCode    goaliyun.Integer
	ErrorMessage goaliyun.String
	Success      bool
	RequestId    goaliyun.String
}
