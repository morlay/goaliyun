package mts

import (
	"github.com/morlay/goaliyun"
)

type UpdateMediaWorkflowRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Topology             string `position:"Query" name:"Topology"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	MediaWorkflowId      string `position:"Query" name:"MediaWorkflowId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *UpdateMediaWorkflowRequest) Invoke(client goaliyun.Client) (*UpdateMediaWorkflowResponse, error) {
	resp := &UpdateMediaWorkflowResponse{}
	err := client.Request("mts", "UpdateMediaWorkflow", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateMediaWorkflowResponse struct {
	RequestId     goaliyun.String
	MediaWorkflow UpdateMediaWorkflowMediaWorkflow
}

type UpdateMediaWorkflowMediaWorkflow struct {
	MediaWorkflowId goaliyun.String
	Name            goaliyun.String
	Topology        goaliyun.String
	TriggerMode     goaliyun.String
	State           goaliyun.String
	CreationTime    goaliyun.String
}
