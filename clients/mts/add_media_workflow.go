package mts

import (
	"github.com/morlay/goaliyun"
)

type AddMediaWorkflowRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Topology             string `position:"Query" name:"Topology"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	TriggerMode          string `position:"Query" name:"TriggerMode"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *AddMediaWorkflowRequest) Invoke(client goaliyun.Client) (*AddMediaWorkflowResponse, error) {
	resp := &AddMediaWorkflowResponse{}
	err := client.Request("mts", "AddMediaWorkflow", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddMediaWorkflowResponse struct {
	RequestId     goaliyun.String
	MediaWorkflow AddMediaWorkflowMediaWorkflow
}

type AddMediaWorkflowMediaWorkflow struct {
	MediaWorkflowId goaliyun.String
	Name            goaliyun.String
	Topology        goaliyun.String
	TriggerMode     goaliyun.String
	State           goaliyun.String
	CreationTime    goaliyun.String
}
