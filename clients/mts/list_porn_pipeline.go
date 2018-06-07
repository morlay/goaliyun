package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListPornPipelineRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	State                string `position:"Query" name:"State"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ListPornPipelineRequest) Invoke(client goaliyun.Client) (*ListPornPipelineResponse, error) {
	resp := &ListPornPipelineResponse{}
	err := client.Request("mts", "ListPornPipeline", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListPornPipelineResponse struct {
	RequestId    goaliyun.String
	TotalCount   goaliyun.Integer
	PageNumber   goaliyun.Integer
	PageSize     goaliyun.Integer
	PipelineList ListPornPipelinePipelineList
}

type ListPornPipelinePipeline struct {
	Id           goaliyun.String
	Name         goaliyun.String
	State        goaliyun.String
	Priority     goaliyun.String
	NotifyConfig ListPornPipelineNotifyConfig
}

type ListPornPipelineNotifyConfig struct {
	Topic goaliyun.String
	Queue goaliyun.String
}

type ListPornPipelinePipelineList []ListPornPipelinePipeline

func (list *ListPornPipelinePipelineList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListPornPipelinePipeline)
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
