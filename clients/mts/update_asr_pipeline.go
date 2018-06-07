package mts

import (
	"github.com/morlay/goaliyun"
)

type UpdateAsrPipelineRequest struct {
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

func (req *UpdateAsrPipelineRequest) Invoke(client goaliyun.Client) (*UpdateAsrPipelineResponse, error) {
	resp := &UpdateAsrPipelineResponse{}
	err := client.Request("mts", "UpdateAsrPipeline", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateAsrPipelineResponse struct {
	RequestId goaliyun.String
	Pipeline  UpdateAsrPipelinePipeline
}

type UpdateAsrPipelinePipeline struct {
	Id           goaliyun.String
	Name         goaliyun.String
	State        goaliyun.String
	Priority     goaliyun.Integer
	NotifyConfig UpdateAsrPipelineNotifyConfig
}

type UpdateAsrPipelineNotifyConfig struct {
	Topic     goaliyun.String
	QueueName goaliyun.String
}
