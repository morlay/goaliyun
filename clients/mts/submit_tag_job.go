package mts

import (
	"github.com/morlay/goaliyun"
)

type SubmitTagJobRequest struct {
	Input                string `position:"Query" name:"Input"`
	UserData             string `position:"Query" name:"UserData"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	TagConfig            string `position:"Query" name:"TagConfig"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PipelineId           string `position:"Query" name:"PipelineId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *SubmitTagJobRequest) Invoke(client goaliyun.Client) (*SubmitTagJobResponse, error) {
	resp := &SubmitTagJobResponse{}
	err := client.Request("mts", "SubmitTagJob", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SubmitTagJobResponse struct {
	RequestId goaliyun.String
	JobId     goaliyun.String
}
