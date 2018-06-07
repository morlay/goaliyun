package mts

import (
	"github.com/morlay/goaliyun"
)

type SubmitCensorJobRequest struct {
	Input                string `position:"Query" name:"Input"`
	UserData             string `position:"Query" name:"UserData"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	CensorConfig         string `position:"Query" name:"CensorConfig"`
	PipelineId           string `position:"Query" name:"PipelineId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *SubmitCensorJobRequest) Invoke(client goaliyun.Client) (*SubmitCensorJobResponse, error) {
	resp := &SubmitCensorJobResponse{}
	err := client.Request("mts", "SubmitCensorJob", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SubmitCensorJobResponse struct {
	RequestId goaliyun.String
	JobId     goaliyun.String
}
