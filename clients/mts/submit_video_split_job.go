package mts

import (
	"github.com/morlay/goaliyun"
)

type SubmitVideoSplitJobRequest struct {
	Input                string `position:"Query" name:"Input"`
	VideoSplitConfig     string `position:"Query" name:"VideoSplitConfig"`
	UserData             string `position:"Query" name:"UserData"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PipelineId           string `position:"Query" name:"PipelineId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *SubmitVideoSplitJobRequest) Invoke(client goaliyun.Client) (*SubmitVideoSplitJobResponse, error) {
	resp := &SubmitVideoSplitJobResponse{}
	err := client.Request("mts", "SubmitVideoSplitJob", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SubmitVideoSplitJobResponse struct {
	RequestId goaliyun.String
	JobId     goaliyun.String
}
