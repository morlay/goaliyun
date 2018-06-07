package mts

import (
	"github.com/morlay/goaliyun"
)

type AddTerrorismPipelineRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	NotifyConfig         string `position:"Query" name:"NotifyConfig"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Priority             int64  `position:"Query" name:"Priority"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *AddTerrorismPipelineRequest) Invoke(client goaliyun.Client) (*AddTerrorismPipelineResponse, error) {
	resp := &AddTerrorismPipelineResponse{}
	err := client.Request("mts", "AddTerrorismPipeline", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddTerrorismPipelineResponse struct {
	RequestId goaliyun.String
	Pipeline  AddTerrorismPipelinePipeline
}

type AddTerrorismPipelinePipeline struct {
	Id           goaliyun.String
	Name         goaliyun.String
	Priority     goaliyun.Integer
	State        goaliyun.String
	NotifyConfig AddTerrorismPipelineNotifyConfig
}

type AddTerrorismPipelineNotifyConfig struct {
	Topic goaliyun.String
	Queue goaliyun.String
}
