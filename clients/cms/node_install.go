package cms

import (
	"github.com/morlay/goaliyun"
)

type NodeInstallRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
	Force      string `position:"Query" name:"Force"`
	UserId     string `position:"Query" name:"UserId"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *NodeInstallRequest) Invoke(client goaliyun.Client) (*NodeInstallResponse, error) {
	resp := &NodeInstallResponse{}
	err := client.Request("cms", "NodeInstall", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type NodeInstallResponse struct {
	ErrorCode    goaliyun.Integer
	ErrorMessage goaliyun.String
	Success      bool
	RequestId    goaliyun.String
}
