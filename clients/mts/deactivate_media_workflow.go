package mts

import (
	"github.com/morlay/goaliyun"
)

type DeactivateMediaWorkflowRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	MediaWorkflowId      string `position:"Query" name:"MediaWorkflowId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeactivateMediaWorkflowRequest) Invoke(client goaliyun.Client) (*DeactivateMediaWorkflowResponse, error) {
	resp := &DeactivateMediaWorkflowResponse{}
	err := client.Request("mts", "DeactivateMediaWorkflow", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeactivateMediaWorkflowResponse struct {
	RequestId     goaliyun.String
	MediaWorkflow DeactivateMediaWorkflowMediaWorkflow
}

type DeactivateMediaWorkflowMediaWorkflow struct {
	MediaWorkflowId goaliyun.String
	Name            goaliyun.String
	Topology        goaliyun.String
	State           goaliyun.String
	CreationTime    goaliyun.String
}
