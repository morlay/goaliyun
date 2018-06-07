package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryPornPipelineListRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	PipelineIds          string `position:"Query" name:"PipelineIds"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *QueryPornPipelineListRequest) Invoke(client goaliyun.Client) (*QueryPornPipelineListResponse, error) {
	resp := &QueryPornPipelineListResponse{}
	err := client.Request("mts", "QueryPornPipelineList", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryPornPipelineListResponse struct {
	RequestId    goaliyun.String
	PipelineList QueryPornPipelineListPipelineList
	NonExistIds  QueryPornPipelineListNonExistIdList
}

type QueryPornPipelineListPipeline struct {
	Id           goaliyun.String
	Name         goaliyun.String
	State        goaliyun.String
	Priority     goaliyun.String
	NotifyConfig QueryPornPipelineListNotifyConfig
}

type QueryPornPipelineListNotifyConfig struct {
	Topic goaliyun.String
	Queue goaliyun.String
}

type QueryPornPipelineListPipelineList []QueryPornPipelineListPipeline

func (list *QueryPornPipelineListPipelineList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryPornPipelineListPipeline)
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

type QueryPornPipelineListNonExistIdList []goaliyun.String

func (list *QueryPornPipelineListNonExistIdList) UnmarshalJSON(data []byte) error {
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
