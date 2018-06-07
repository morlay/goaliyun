package mts

import (
	"github.com/morlay/goaliyun"
)

type SubmitFacerecogJobRequest struct {
	Input                string `position:"Query" name:"Input"`
	UserData             string `position:"Query" name:"UserData"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	FacerecogConfig      string `position:"Query" name:"FacerecogConfig"`
	PipelineId           string `position:"Query" name:"PipelineId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *SubmitFacerecogJobRequest) Invoke(client goaliyun.Client) (*SubmitFacerecogJobResponse, error) {
	resp := &SubmitFacerecogJobResponse{}
	err := client.Request("mts", "SubmitFacerecogJob", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SubmitFacerecogJobResponse struct {
	RequestId goaliyun.String
	JobId     goaliyun.String
}
