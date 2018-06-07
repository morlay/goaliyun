package mts

import (
	"github.com/morlay/goaliyun"
)

type SubmitCoverJobRequest struct {
	Input                string `position:"Query" name:"Input"`
	UserData             string `position:"Query" name:"UserData"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	CoverConfig          string `position:"Query" name:"CoverConfig"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PipelineId           string `position:"Query" name:"PipelineId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *SubmitCoverJobRequest) Invoke(client goaliyun.Client) (*SubmitCoverJobResponse, error) {
	resp := &SubmitCoverJobResponse{}
	err := client.Request("mts", "SubmitCoverJob", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SubmitCoverJobResponse struct {
	RequestId goaliyun.String
	JobId     goaliyun.String
}
