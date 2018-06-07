package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryCoverPipelineListRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	PipelineIds          string `position:"Query" name:"PipelineIds"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *QueryCoverPipelineListRequest) Invoke(client goaliyun.Client) (*QueryCoverPipelineListResponse, error) {
	resp := &QueryCoverPipelineListResponse{}
	err := client.Request("mts", "QueryCoverPipelineList", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryCoverPipelineListResponse struct {
	RequestId    goaliyun.String
	PipelineList QueryCoverPipelineListPipelineList
	NonExistIds  QueryCoverPipelineListNonExistIdList
}

type QueryCoverPipelineListPipeline struct {
	Id           goaliyun.String
	Name         goaliyun.String
	State        goaliyun.String
	Priority     goaliyun.String
	Role         goaliyun.String
	NotifyConfig QueryCoverPipelineListNotifyConfig
}

type QueryCoverPipelineListNotifyConfig struct {
	Topic goaliyun.String
	Queue goaliyun.String
}

type QueryCoverPipelineListPipelineList []QueryCoverPipelineListPipeline

func (list *QueryCoverPipelineListPipelineList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryCoverPipelineListPipeline)
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

type QueryCoverPipelineListNonExistIdList []goaliyun.String

func (list *QueryCoverPipelineListNonExistIdList) UnmarshalJSON(data []byte) error {
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
