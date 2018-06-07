package mts

import (
	"github.com/morlay/goaliyun"
)

type UpdateMediaWorkflowTriggerModeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	MediaWorkflowId      string `position:"Query" name:"MediaWorkflowId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	TriggerMode          string `position:"Query" name:"TriggerMode"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *UpdateMediaWorkflowTriggerModeRequest) Invoke(client goaliyun.Client) (*UpdateMediaWorkflowTriggerModeResponse, error) {
	resp := &UpdateMediaWorkflowTriggerModeResponse{}
	err := client.Request("mts", "UpdateMediaWorkflowTriggerMode", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateMediaWorkflowTriggerModeResponse struct {
	RequestId     goaliyun.String
	MediaWorkflow UpdateMediaWorkflowTriggerModeMediaWorkflow
}

type UpdateMediaWorkflowTriggerModeMediaWorkflow struct {
	MediaWorkflowId goaliyun.String
	Name            goaliyun.String
	Topology        goaliyun.String
	TriggerMode     goaliyun.String
	State           goaliyun.String
	CreationTime    goaliyun.String
}
