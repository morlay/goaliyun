package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListVideoSummaryPipelineRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	State                string `position:"Query" name:"State"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ListVideoSummaryPipelineRequest) Invoke(client goaliyun.Client) (*ListVideoSummaryPipelineResponse, error) {
	resp := &ListVideoSummaryPipelineResponse{}
	err := client.Request("mts", "ListVideoSummaryPipeline", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListVideoSummaryPipelineResponse struct {
	RequestId    goaliyun.String
	TotalCount   goaliyun.Integer
	PageNumber   goaliyun.Integer
	PageSize     goaliyun.Integer
	PipelineList ListVideoSummaryPipelinePipelineList
}

type ListVideoSummaryPipelinePipeline struct {
	Id           goaliyun.String
	Name         goaliyun.String
	State        goaliyun.String
	Priority     goaliyun.String
	NotifyConfig ListVideoSummaryPipelineNotifyConfig
}

type ListVideoSummaryPipelineNotifyConfig struct {
	Topic     goaliyun.String
	QueueName goaliyun.String
}

type ListVideoSummaryPipelinePipelineList []ListVideoSummaryPipelinePipeline

func (list *ListVideoSummaryPipelinePipelineList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListVideoSummaryPipelinePipeline)
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
