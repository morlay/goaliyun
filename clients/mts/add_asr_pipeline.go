package mts

import (
	"github.com/morlay/goaliyun"
)

type AddAsrPipelineRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	NotifyConfig         string `position:"Query" name:"NotifyConfig"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Priority             int64  `position:"Query" name:"Priority"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *AddAsrPipelineRequest) Invoke(client goaliyun.Client) (*AddAsrPipelineResponse, error) {
	resp := &AddAsrPipelineResponse{}
	err := client.Request("mts", "AddAsrPipeline", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddAsrPipelineResponse struct {
	RequestId goaliyun.String
	Pipeline  AddAsrPipelinePipeline
}

type AddAsrPipelinePipeline struct {
	Id           goaliyun.String
	Name         goaliyun.String
	Priority     goaliyun.Integer
	State        goaliyun.String
	NotifyConfig AddAsrPipelineNotifyConfig
}

type AddAsrPipelineNotifyConfig struct {
	Topic goaliyun.String
	Queue goaliyun.String
}
