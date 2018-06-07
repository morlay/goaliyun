package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryAsrPipelineListRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	PipelineIds          string `position:"Query" name:"PipelineIds"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *QueryAsrPipelineListRequest) Invoke(client goaliyun.Client) (*QueryAsrPipelineListResponse, error) {
	resp := &QueryAsrPipelineListResponse{}
	err := client.Request("mts", "QueryAsrPipelineList", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryAsrPipelineListResponse struct {
	RequestId    goaliyun.String
	PipelineList QueryAsrPipelineListPipelineList
	NonExistIds  QueryAsrPipelineListNonExistIdList
}

type QueryAsrPipelineListPipeline struct {
	Id           goaliyun.String
	Name         goaliyun.String
	State        goaliyun.String
	Priority     goaliyun.String
	NotifyConfig QueryAsrPipelineListNotifyConfig
}

type QueryAsrPipelineListNotifyConfig struct {
	Topic     goaliyun.String
	QueueName goaliyun.String
}

type QueryAsrPipelineListPipelineList []QueryAsrPipelineListPipeline

func (list *QueryAsrPipelineListPipelineList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryAsrPipelineListPipeline)
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

type QueryAsrPipelineListNonExistIdList []goaliyun.String

func (list *QueryAsrPipelineListNonExistIdList) UnmarshalJSON(data []byte) error {
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
