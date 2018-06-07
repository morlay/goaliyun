package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryPipelineListRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	PipelineIds          string `position:"Query" name:"PipelineIds"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *QueryPipelineListRequest) Invoke(client goaliyun.Client) (*QueryPipelineListResponse, error) {
	resp := &QueryPipelineListResponse{}
	err := client.Request("mts", "QueryPipelineList", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryPipelineListResponse struct {
	RequestId    goaliyun.String
	PipelineList QueryPipelineListPipelineList
	NonExistPids QueryPipelineListNonExistPidList
}

type QueryPipelineListPipeline struct {
	Id           goaliyun.String
	Name         goaliyun.String
	State        goaliyun.String
	Speed        goaliyun.String
	SpeedLevel   goaliyun.Integer
	Role         goaliyun.String
	NotifyConfig QueryPipelineListNotifyConfig
}

type QueryPipelineListNotifyConfig struct {
	Topic     goaliyun.String
	QueueName goaliyun.String
}

type QueryPipelineListPipelineList []QueryPipelineListPipeline

func (list *QueryPipelineListPipelineList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryPipelineListPipeline)
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

type QueryPipelineListNonExistPidList []goaliyun.String

func (list *QueryPipelineListNonExistPidList) UnmarshalJSON(data []byte) error {
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
