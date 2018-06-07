package mts

import (
	"github.com/morlay/goaliyun"
)

type SubmitFpShotJobRequest struct {
	Input                string `position:"Query" name:"Input"`
	UserData             string `position:"Query" name:"UserData"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	FpShotConfig         string `position:"Query" name:"FpShotConfig"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PipelineId           string `position:"Query" name:"PipelineId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *SubmitFpShotJobRequest) Invoke(client goaliyun.Client) (*SubmitFpShotJobResponse, error) {
	resp := &SubmitFpShotJobResponse{}
	err := client.Request("mts", "SubmitFpShotJob", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SubmitFpShotJobResponse struct {
	RequestId goaliyun.String
	JobId     goaliyun.String
}
