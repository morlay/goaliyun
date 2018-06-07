package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryMediaWorkflowListRequest struct {
	MediaWorkflowIds     string `position:"Query" name:"MediaWorkflowIds"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *QueryMediaWorkflowListRequest) Invoke(client goaliyun.Client) (*QueryMediaWorkflowListResponse, error) {
	resp := &QueryMediaWorkflowListResponse{}
	err := client.Request("mts", "QueryMediaWorkflowList", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryMediaWorkflowListResponse struct {
	RequestId                goaliyun.String
	MediaWorkflowList        QueryMediaWorkflowListMediaWorkflowList
	NonExistMediaWorkflowIds QueryMediaWorkflowListNonExistMediaWorkflowIdList
}

type QueryMediaWorkflowListMediaWorkflow struct {
	MediaWorkflowId goaliyun.String
	Name            goaliyun.String
	Topology        goaliyun.String
	TriggerMode     goaliyun.String
	State           goaliyun.String
	CreationTime    goaliyun.String
}

type QueryMediaWorkflowListMediaWorkflowList []QueryMediaWorkflowListMediaWorkflow

func (list *QueryMediaWorkflowListMediaWorkflowList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaWorkflowListMediaWorkflow)
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

type QueryMediaWorkflowListNonExistMediaWorkflowIdList []goaliyun.String

func (list *QueryMediaWorkflowListNonExistMediaWorkflowIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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
