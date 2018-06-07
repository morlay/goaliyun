package mts

import (
	"github.com/morlay/goaliyun"
)

type UpdatePipelineRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Role                 string `position:"Query" name:"Role"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	State                string `position:"Query" name:"State"`
	NotifyConfig         string `position:"Query" name:"NotifyConfig"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PipelineId           string `position:"Query" name:"PipelineId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *UpdatePipelineRequest) Invoke(client goaliyun.Client) (*UpdatePipelineResponse, error) {
	resp := &UpdatePipelineResponse{}
	err := client.Request("mts", "UpdatePipeline", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdatePipelineResponse struct {
	RequestId goaliyun.String
	Pipeline  UpdatePipelinePipeline
}

type UpdatePipelinePipeline struct {
	Id           goaliyun.String
	Name         goaliyun.String
	State        goaliyun.String
	Speed        goaliyun.String
	Role         goaliyun.String
	NotifyConfig UpdatePipelineNotifyConfig
}

type UpdatePipelineNotifyConfig struct {
	Topic     goaliyun.String
	QueueName goaliyun.String
}
