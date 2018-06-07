package mts

import (
	"github.com/morlay/goaliyun"
)

type SubmitPornJobRequest struct {
	Input                string `position:"Query" name:"Input"`
	UserData             string `position:"Query" name:"UserData"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PornConfig           string `position:"Query" name:"PornConfig"`
	PipelineId           string `position:"Query" name:"PipelineId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *SubmitPornJobRequest) Invoke(client goaliyun.Client) (*SubmitPornJobResponse, error) {
	resp := &SubmitPornJobResponse{}
	err := client.Request("mts", "SubmitPornJob", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SubmitPornJobResponse struct {
	RequestId goaliyun.String
	JobId     goaliyun.String
}
