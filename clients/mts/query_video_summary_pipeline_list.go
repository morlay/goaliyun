package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryVideoSummaryPipelineListRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	PipelineIds          string `position:"Query" name:"PipelineIds"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *QueryVideoSummaryPipelineListRequest) Invoke(client goaliyun.Client) (*QueryVideoSummaryPipelineListResponse, error) {
	resp := &QueryVideoSummaryPipelineListResponse{}
	err := client.Request("mts", "QueryVideoSummaryPipelineList", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryVideoSummaryPipelineListResponse struct {
	RequestId    goaliyun.String
	PipelineList QueryVideoSummaryPipelineListPipelineList
	NonExistIds  QueryVideoSummaryPipelineListNonExistIdList
}

type QueryVideoSummaryPipelineListPipeline struct {
	Id           goaliyun.String
	Name         goaliyun.String
	State        goaliyun.String
	Priority     goaliyun.String
	NotifyConfig QueryVideoSummaryPipelineListNotifyConfig
}

type QueryVideoSummaryPipelineListNotifyConfig struct {
	Topic     goaliyun.String
	QueueName goaliyun.String
}

type QueryVideoSummaryPipelineListPipelineList []QueryVideoSummaryPipelineListPipeline

func (list *QueryVideoSummaryPipelineListPipelineList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryVideoSummaryPipelineListPipeline)
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

type QueryVideoSummaryPipelineListNonExistIdList []goaliyun.String

func (list *QueryVideoSummaryPipelineListNonExistIdList) UnmarshalJSON(data []byte) error {
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
