package mts

import (
	"github.com/morlay/goaliyun"
)

type DeletePipelineRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PipelineId           string `position:"Query" name:"PipelineId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeletePipelineRequest) Invoke(client goaliyun.Client) (*DeletePipelineResponse, error) {
	resp := &DeletePipelineResponse{}
	err := client.Request("mts", "DeletePipeline", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeletePipelineResponse struct {
	RequestId  goaliyun.String
	PipelineId goaliyun.String
}
