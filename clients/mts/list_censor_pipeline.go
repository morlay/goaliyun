package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListCensorPipelineRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	State                string `position:"Query" name:"State"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ListCensorPipelineRequest) Invoke(client goaliyun.Client) (*ListCensorPipelineResponse, error) {
	resp := &ListCensorPipelineResponse{}
	err := client.Request("mts", "ListCensorPipeline", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListCensorPipelineResponse struct {
	RequestId    goaliyun.String
	TotalCount   goaliyun.Integer
	PageNumber   goaliyun.Integer
	PageSize     goaliyun.Integer
	PipelineList ListCensorPipelinePipelineList
}

type ListCensorPipelinePipeline struct {
	Id           goaliyun.String
	Name         goaliyun.String
	State        goaliyun.String
	Priority     goaliyun.String
	NotifyConfig ListCensorPipelineNotifyConfig
}

type ListCensorPipelineNotifyConfig struct {
	Topic goaliyun.String
	Queue goaliyun.String
}

type ListCensorPipelinePipelineList []ListCensorPipelinePipeline

func (list *ListCensorPipelinePipelineList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListCensorPipelinePipeline)
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
