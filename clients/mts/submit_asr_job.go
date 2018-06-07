package mts

import (
	"github.com/morlay/goaliyun"
)

type SubmitAsrJobRequest struct {
	Input                string `position:"Query" name:"Input"`
	UserData             string `position:"Query" name:"UserData"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	AsrConfig            string `position:"Query" name:"AsrConfig"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PipelineId           string `position:"Query" name:"PipelineId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *SubmitAsrJobRequest) Invoke(client goaliyun.Client) (*SubmitAsrJobResponse, error) {
	resp := &SubmitAsrJobResponse{}
	err := client.Request("mts", "SubmitAsrJob", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SubmitAsrJobResponse struct {
	RequestId goaliyun.String
	JobId     goaliyun.String
}
