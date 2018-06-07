package mts

import (
	"github.com/morlay/goaliyun"
)

type AddPornPipelineRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	NotifyConfig         string `position:"Query" name:"NotifyConfig"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Priority             int64  `position:"Query" name:"Priority"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *AddPornPipelineRequest) Invoke(client goaliyun.Client) (*AddPornPipelineResponse, error) {
	resp := &AddPornPipelineResponse{}
	err := client.Request("mts", "AddPornPipeline", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddPornPipelineResponse struct {
	RequestId goaliyun.String
	Pipeline  AddPornPipelinePipeline
}

type AddPornPipelinePipeline struct {
	Id           goaliyun.String
	Name         goaliyun.String
	Priority     goaliyun.Integer
	State        goaliyun.String
	NotifyConfig AddPornPipelineNotifyConfig
}

type AddPornPipelineNotifyConfig struct {
	Topic goaliyun.String
	Queue goaliyun.String
}
