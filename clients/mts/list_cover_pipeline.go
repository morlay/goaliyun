package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListCoverPipelineRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	State                string `position:"Query" name:"State"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ListCoverPipelineRequest) Invoke(client goaliyun.Client) (*ListCoverPipelineResponse, error) {
	resp := &ListCoverPipelineResponse{}
	err := client.Request("mts", "ListCoverPipeline", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListCoverPipelineResponse struct {
	RequestId    goaliyun.String
	TotalCount   goaliyun.Integer
	PageNumber   goaliyun.Integer
	PageSize     goaliyun.Integer
	PipelineList ListCoverPipelinePipelineList
}

type ListCoverPipelinePipeline struct {
	UserId       goaliyun.Integer
	PipelineId   goaliyun.String
	Name         goaliyun.String
	State        goaliyun.String
	Priority     goaliyun.String
	QuotaNum     goaliyun.Integer
	QuotaUsed    goaliyun.Integer
	NotifyConfig goaliyun.String
	Role         goaliyun.String
	ExtendConfig goaliyun.String
}

type ListCoverPipelinePipelineList []ListCoverPipelinePipeline

func (list *ListCoverPipelinePipelineList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListCoverPipelinePipeline)
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
