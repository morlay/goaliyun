package mts

import (
	"github.com/morlay/goaliyun"
)

type UpdatePornPipelineRequest struct {
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

func (req *UpdatePornPipelineRequest) Invoke(client goaliyun.Client) (*UpdatePornPipelineResponse, error) {
	resp := &UpdatePornPipelineResponse{}
	err := client.Request("mts", "UpdatePornPipeline", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdatePornPipelineResponse struct {
	RequestId goaliyun.String
	Pipeline  UpdatePornPipelinePipeline
}

type UpdatePornPipelinePipeline struct {
	Id           goaliyun.String
	Name         goaliyun.String
	State        goaliyun.String
	Priority     goaliyun.Integer
	NotifyConfig UpdatePornPipelineNotifyConfig
}

type UpdatePornPipelineNotifyConfig struct {
	Topic goaliyun.String
	Queue goaliyun.String
}
