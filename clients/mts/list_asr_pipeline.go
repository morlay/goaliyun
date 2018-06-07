package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListAsrPipelineRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	State                string `position:"Query" name:"State"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ListAsrPipelineRequest) Invoke(client goaliyun.Client) (*ListAsrPipelineResponse, error) {
	resp := &ListAsrPipelineResponse{}
	err := client.Request("mts", "ListAsrPipeline", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListAsrPipelineResponse struct {
	RequestId    goaliyun.String
	TotalCount   goaliyun.Integer
	PageNumber   goaliyun.Integer
	PageSize     goaliyun.Integer
	PipelineList ListAsrPipelinePipelineList
}

type ListAsrPipelinePipeline struct {
	Id           goaliyun.String
	Name         goaliyun.String
	State        goaliyun.String
	Priority     goaliyun.String
	NotifyConfig ListAsrPipelineNotifyConfig
}

type ListAsrPipelineNotifyConfig struct {
	Topic     goaliyun.String
	QueueName goaliyun.String
}

type ListAsrPipelinePipelineList []ListAsrPipelinePipeline

func (list *ListAsrPipelinePipelineList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAsrPipelinePipeline)
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
