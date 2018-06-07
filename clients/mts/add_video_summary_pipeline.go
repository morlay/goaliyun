package mts

import (
	"github.com/morlay/goaliyun"
)

type AddVideoSummaryPipelineRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	NotifyConfig         string `position:"Query" name:"NotifyConfig"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Priority             int64  `position:"Query" name:"Priority"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *AddVideoSummaryPipelineRequest) Invoke(client goaliyun.Client) (*AddVideoSummaryPipelineResponse, error) {
	resp := &AddVideoSummaryPipelineResponse{}
	err := client.Request("mts", "AddVideoSummaryPipeline", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddVideoSummaryPipelineResponse struct {
	RequestId goaliyun.String
	Pipeline  AddVideoSummaryPipelinePipeline
}

type AddVideoSummaryPipelinePipeline struct {
	Id           goaliyun.String
	Name         goaliyun.String
	Priority     goaliyun.Integer
	State        goaliyun.String
	NotifyConfig AddVideoSummaryPipelineNotifyConfig
}

type AddVideoSummaryPipelineNotifyConfig struct {
	Topic goaliyun.String
	Queue goaliyun.String
}
