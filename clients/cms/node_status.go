package cms

import (
	"github.com/morlay/goaliyun"
)

type NodeStatusRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *NodeStatusRequest) Invoke(client goaliyun.Client) (*NodeStatusResponse, error) {
	resp := &NodeStatusResponse{}
	err := client.Request("cms", "NodeStatus", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type NodeStatusResponse struct {
	ErrorCode    goaliyun.Integer
	ErrorMessage goaliyun.String
	Success      bool
	RequestId    goaliyun.String
	InstanceId   goaliyun.String
	AutoInstall  bool
	Status       goaliyun.String
}
