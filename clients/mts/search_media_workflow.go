package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type SearchMediaWorkflowRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	StateList            string `position:"Query" name:"StateList"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *SearchMediaWorkflowRequest) Invoke(client goaliyun.Client) (*SearchMediaWorkflowResponse, error) {
	resp := &SearchMediaWorkflowResponse{}
	err := client.Request("mts", "SearchMediaWorkflow", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SearchMediaWorkflowResponse struct {
	RequestId         goaliyun.String
	TotalCount        goaliyun.Integer
	PageNumber        goaliyun.Integer
	PageSize          goaliyun.Integer
	MediaWorkflowList SearchMediaWorkflowMediaWorkflowList
}

type SearchMediaWorkflowMediaWorkflow struct {
	MediaWorkflowId goaliyun.String
	Name            goaliyun.String
	Topology        goaliyun.String
	TriggerMode     goaliyun.String
	State           goaliyun.String
	CreationTime    goaliyun.String
}

type SearchMediaWorkflowMediaWorkflowList []SearchMediaWorkflowMediaWorkflow

func (list *SearchMediaWorkflowMediaWorkflowList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SearchMediaWorkflowMediaWorkflow)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
