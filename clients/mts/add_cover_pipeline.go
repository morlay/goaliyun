package mts

import (
	"github.com/morlay/goaliyun"
)

type AddCoverPipelineRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Role                 string `position:"Query" name:"Role"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	NotifyConfig         string `position:"Query" name:"NotifyConfig"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Priority             string `position:"Query" name:"Priority"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *AddCoverPipelineRequest) Invoke(client goaliyun.Client) (*AddCoverPipelineResponse, error) {
	resp := &AddCoverPipelineResponse{}
	err := client.Request("mts", "AddCoverPipeline", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddCoverPipelineResponse struct {
	RequestId goaliyun.String
	Pipeline  AddCoverPipelinePipeline
}

type AddCoverPipelinePipeline struct {
	Id           goaliyun.String
	Name         goaliyun.String
	Priority     goaliyun.String
	State        goaliyun.String
	Role         goaliyun.String
	NotifyConfig AddCoverPipelineNotifyConfig
}

type AddCoverPipelineNotifyConfig struct {
	Topic goaliyun.String
	Queue goaliyun.String
}
