package mts

import (
	"github.com/morlay/goaliyun"
)

type UpdateCensorPipelineRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	State                string `position:"Query" name:"State"`
	NotifyConfig         string `position:"Query" name:"NotifyConfig"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Priority             int64  `position:"Query" name:"Priority"`
	PipelineId           string `position:"Query" name:"PipelineId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *UpdateCensorPipelineRequest) Invoke(client goaliyun.Client) (*UpdateCensorPipelineResponse, error) {
	resp := &UpdateCensorPipelineResponse{}
	err := client.Request("mts", "UpdateCensorPipeline", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateCensorPipelineResponse struct {
	RequestId goaliyun.String
	Pipeline  UpdateCensorPipelinePipeline
}

type UpdateCensorPipelinePipeline struct {
	Id           goaliyun.String
	Name         goaliyun.String
	State        goaliyun.String
	Priority     goaliyun.Integer
	NotifyConfig UpdateCensorPipelineNotifyConfig
}

type UpdateCensorPipelineNotifyConfig struct {
	Topic goaliyun.String
	Queue goaliyun.String
}
