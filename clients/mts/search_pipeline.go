package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type SearchPipelineRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	State                string `position:"Query" name:"State"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *SearchPipelineRequest) Invoke(client goaliyun.Client) (*SearchPipelineResponse, error) {
	resp := &SearchPipelineResponse{}
	err := client.Request("mts", "SearchPipeline", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SearchPipelineResponse struct {
	RequestId    goaliyun.String
	TotalCount   goaliyun.Integer
	PageNumber   goaliyun.Integer
	PageSize     goaliyun.Integer
	PipelineList SearchPipelinePipelineList
}

type SearchPipelinePipeline struct {
	Id           goaliyun.String
	Name         goaliyun.String
	State        goaliyun.String
	Speed        goaliyun.String
	SpeedLevel   goaliyun.Integer
	Role         goaliyun.String
	NotifyConfig SearchPipelineNotifyConfig
}

type SearchPipelineNotifyConfig struct {
	Topic     goaliyun.String
	QueueName goaliyun.String
}

type SearchPipelinePipelineList []SearchPipelinePipeline

func (list *SearchPipelinePipelineList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SearchPipelinePipeline)
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
