package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryCensorPipelineListRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	PipelineIds          string `position:"Query" name:"PipelineIds"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *QueryCensorPipelineListRequest) Invoke(client goaliyun.Client) (*QueryCensorPipelineListResponse, error) {
	resp := &QueryCensorPipelineListResponse{}
	err := client.Request("mts", "QueryCensorPipelineList", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryCensorPipelineListResponse struct {
	RequestId    goaliyun.String
	PipelineList QueryCensorPipelineListPipelineList
	NonExistIds  QueryCensorPipelineListNonExistIdList
}

type QueryCensorPipelineListPipeline struct {
	Id           goaliyun.String
	Name         goaliyun.String
	State        goaliyun.String
	Priority     goaliyun.String
	NotifyConfig QueryCensorPipelineListNotifyConfig
}

type QueryCensorPipelineListNotifyConfig struct {
	Topic goaliyun.String
	Queue goaliyun.String
}

type QueryCensorPipelineListPipelineList []QueryCensorPipelineListPipeline

func (list *QueryCensorPipelineListPipelineList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryCensorPipelineListPipeline)
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

type QueryCensorPipelineListNonExistIdList []goaliyun.String

func (list *QueryCensorPipelineListNonExistIdList) UnmarshalJSON(data []byte) error {
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
