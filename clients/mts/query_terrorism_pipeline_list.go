package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryTerrorismPipelineListRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	PipelineIds          string `position:"Query" name:"PipelineIds"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *QueryTerrorismPipelineListRequest) Invoke(client goaliyun.Client) (*QueryTerrorismPipelineListResponse, error) {
	resp := &QueryTerrorismPipelineListResponse{}
	err := client.Request("mts", "QueryTerrorismPipelineList", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryTerrorismPipelineListResponse struct {
	RequestId    goaliyun.String
	PipelineList QueryTerrorismPipelineListPipelineList
	NonExistIds  QueryTerrorismPipelineListNonExistIdList
}

type QueryTerrorismPipelineListPipeline struct {
	Id           goaliyun.String
	Name         goaliyun.String
	State        goaliyun.String
	Priority     goaliyun.String
	NotifyConfig QueryTerrorismPipelineListNotifyConfig
}

type QueryTerrorismPipelineListNotifyConfig struct {
	Topic goaliyun.String
	Queue goaliyun.String
}

type QueryTerrorismPipelineListPipelineList []QueryTerrorismPipelineListPipeline

func (list *QueryTerrorismPipelineListPipelineList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryTerrorismPipelineListPipeline)
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

type QueryTerrorismPipelineListNonExistIdList []goaliyun.String

func (list *QueryTerrorismPipelineListNonExistIdList) UnmarshalJSON(data []byte) error {
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
