package mts

import (
	"github.com/morlay/goaliyun"
)

type DeleteMediaWorkflowRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	MediaWorkflowId      string `position:"Query" name:"MediaWorkflowId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteMediaWorkflowRequest) Invoke(client goaliyun.Client) (*DeleteMediaWorkflowResponse, error) {
	resp := &DeleteMediaWorkflowResponse{}
	err := client.Request("mts", "DeleteMediaWorkflow", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteMediaWorkflowResponse struct {
	RequestId     goaliyun.String
	MediaWorkflow DeleteMediaWorkflowMediaWorkflow
}

type DeleteMediaWorkflowMediaWorkflow struct {
	MediaWorkflowId goaliyun.String
	Name            goaliyun.String
	Topology        goaliyun.String
	State           goaliyun.String
	CreationTime    goaliyun.String
}
