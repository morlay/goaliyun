package cms

import (
	"github.com/morlay/goaliyun"
)

type NodeProcessDeleteRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
	Name       string `position:"Query" name:"Name"`
	Id         string `position:"Query" name:"Id"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *NodeProcessDeleteRequest) Invoke(client goaliyun.Client) (*NodeProcessDeleteResponse, error) {
	resp := &NodeProcessDeleteResponse{}
	err := client.Request("cms", "NodeProcessDelete", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type NodeProcessDeleteResponse struct {
	ErrorCode    goaliyun.Integer
	ErrorMessage goaliyun.String
	Success      bool
	RequestId    goaliyun.String
}
