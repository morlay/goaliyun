package mts

import (
	"github.com/morlay/goaliyun"
)

type UpdateCoverPipelineRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Role                 string `position:"Query" name:"Role"`
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

func (req *UpdateCoverPipelineRequest) Invoke(client goaliyun.Client) (*UpdateCoverPipelineResponse, error) {
	resp := &UpdateCoverPipelineResponse{}
	err := client.Request("mts", "UpdateCoverPipeline", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateCoverPipelineResponse struct {
	RequestId goaliyun.String
	Pipeline  UpdateCoverPipelinePipeline
}

type UpdateCoverPipelinePipeline struct {
	Id           goaliyun.String
	Name         goaliyun.String
	State        goaliyun.String
	Priority     goaliyun.Integer
	Role         goaliyun.String
	NotifyConfig UpdateCoverPipelineNotifyConfig
}

type UpdateCoverPipelineNotifyConfig struct {
	Topic goaliyun.String
	Queue goaliyun.String
}
