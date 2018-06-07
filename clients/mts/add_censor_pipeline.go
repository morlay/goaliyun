package mts

import (
	"github.com/morlay/goaliyun"
)

type AddCensorPipelineRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	NotifyConfig         string `position:"Query" name:"NotifyConfig"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Priority             int64  `position:"Query" name:"Priority"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *AddCensorPipelineRequest) Invoke(client goaliyun.Client) (*AddCensorPipelineResponse, error) {
	resp := &AddCensorPipelineResponse{}
	err := client.Request("mts", "AddCensorPipeline", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddCensorPipelineResponse struct {
	RequestId goaliyun.String
	Pipeline  AddCensorPipelinePipeline
}

type AddCensorPipelinePipeline struct {
	Id           goaliyun.String
	Name         goaliyun.String
	Priority     goaliyun.Integer
	State        goaliyun.String
	NotifyConfig AddCensorPipelineNotifyConfig
}

type AddCensorPipelineNotifyConfig struct {
	Topic goaliyun.String
	Queue goaliyun.String
}
