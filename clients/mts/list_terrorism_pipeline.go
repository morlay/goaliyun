package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListTerrorismPipelineRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	State                string `position:"Query" name:"State"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ListTerrorismPipelineRequest) Invoke(client goaliyun.Client) (*ListTerrorismPipelineResponse, error) {
	resp := &ListTerrorismPipelineResponse{}
	err := client.Request("mts", "ListTerrorismPipeline", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListTerrorismPipelineResponse struct {
	RequestId    goaliyun.String
	TotalCount   goaliyun.Integer
	PageNumber   goaliyun.Integer
	PageSize     goaliyun.Integer
	PipelineList ListTerrorismPipelinePipelineList
}

type ListTerrorismPipelinePipeline struct {
	Id           goaliyun.String
	Name         goaliyun.String
	State        goaliyun.String
	Priority     goaliyun.String
	NotifyConfig ListTerrorismPipelineNotifyConfig
}

type ListTerrorismPipelineNotifyConfig struct {
	Topic goaliyun.String
	Queue goaliyun.String
}

type ListTerrorismPipelinePipelineList []ListTerrorismPipelinePipeline

func (list *ListTerrorismPipelinePipelineList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListTerrorismPipelinePipeline)
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
