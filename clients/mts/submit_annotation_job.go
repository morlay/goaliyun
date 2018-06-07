package mts

import (
	"github.com/morlay/goaliyun"
)

type SubmitAnnotationJobRequest struct {
	Input                string `position:"Query" name:"Input"`
	UserData             string `position:"Query" name:"UserData"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	AnnotationConfig     string `position:"Query" name:"AnnotationConfig"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PipelineId           string `position:"Query" name:"PipelineId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *SubmitAnnotationJobRequest) Invoke(client goaliyun.Client) (*SubmitAnnotationJobResponse, error) {
	resp := &SubmitAnnotationJobResponse{}
	err := client.Request("mts", "SubmitAnnotationJob", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SubmitAnnotationJobResponse struct {
	RequestId goaliyun.String
	JobId     goaliyun.String
}
