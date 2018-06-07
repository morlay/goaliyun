package mts

import (
	"github.com/morlay/goaliyun"
)

type SubmitMediaDetailJobRequest struct {
	Input                string `position:"Query" name:"Input"`
	UserData             string `position:"Query" name:"UserData"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	MediaDetailConfig    string `position:"Query" name:"MediaDetailConfig"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PipelineId           string `position:"Query" name:"PipelineId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *SubmitMediaDetailJobRequest) Invoke(client goaliyun.Client) (*SubmitMediaDetailJobResponse, error) {
	resp := &SubmitMediaDetailJobResponse{}
	err := client.Request("mts", "SubmitMediaDetailJob", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SubmitMediaDetailJobResponse struct {
	RequestId goaliyun.String
	JobId     goaliyun.String
}
