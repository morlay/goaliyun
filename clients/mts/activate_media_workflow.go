package mts

import (
	"github.com/morlay/goaliyun"
)

type ActivateMediaWorkflowRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	MediaWorkflowId      string `position:"Query" name:"MediaWorkflowId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ActivateMediaWorkflowRequest) Invoke(client goaliyun.Client) (*ActivateMediaWorkflowResponse, error) {
	resp := &ActivateMediaWorkflowResponse{}
	err := client.Request("mts", "ActivateMediaWorkflow", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ActivateMediaWorkflowResponse struct {
	RequestId     goaliyun.String
	MediaWorkflow ActivateMediaWorkflowMediaWorkflow
}

type ActivateMediaWorkflowMediaWorkflow struct {
	MediaWorkflowId goaliyun.String
	Name            goaliyun.String
	Topology        goaliyun.String
	State           goaliyun.String
	CreationTime    goaliyun.String
}
