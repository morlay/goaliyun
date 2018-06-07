package cms

import (
	"github.com/morlay/goaliyun"
)

type NodeProcessCreateRequest struct {
	InstanceId  string `position:"Query" name:"InstanceId"`
	ProcessName string `position:"Query" name:"ProcessName"`
	Name        string `position:"Query" name:"Name"`
	ProcessUser string `position:"Query" name:"ProcessUser"`
	Command     string `position:"Query" name:"Command"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *NodeProcessCreateRequest) Invoke(client goaliyun.Client) (*NodeProcessCreateResponse, error) {
	resp := &NodeProcessCreateResponse{}
	err := client.Request("cms", "NodeProcessCreate", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type NodeProcessCreateResponse struct {
	ErrorCode    goaliyun.Integer
	ErrorMessage goaliyun.String
	Success      bool
	RequestId    goaliyun.String
}
