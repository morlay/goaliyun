package mts

import (
	"github.com/morlay/goaliyun"
)

type AddPipelineRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Role                 string `position:"Query" name:"Role"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	NotifyConfig         string `position:"Query" name:"NotifyConfig"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	SpeedLevel           int64  `position:"Query" name:"SpeedLevel"`
	Speed                string `position:"Query" name:"Speed"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *AddPipelineRequest) Invoke(client goaliyun.Client) (*AddPipelineResponse, error) {
	resp := &AddPipelineResponse{}
	err := client.Request("mts", "AddPipeline", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddPipelineResponse struct {
	RequestId goaliyun.String
	Pipeline  AddPipelinePipeline
}

type AddPipelinePipeline struct {
	Id           goaliyun.String
	Name         goaliyun.String
	State        goaliyun.String
	Speed        goaliyun.String
	SpeedLevel   goaliyun.Integer
	Role         goaliyun.String
	NotifyConfig AddPipelineNotifyConfig
}

type AddPipelineNotifyConfig struct {
	Topic     goaliyun.String
	QueueName goaliyun.String
}
