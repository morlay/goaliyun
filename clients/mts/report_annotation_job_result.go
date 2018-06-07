package mts

import (
	"github.com/morlay/goaliyun"
)

type ReportAnnotationJobResultRequest struct {
	Annotation           string `position:"Query" name:"Annotation"`
	JobId                string `position:"Query" name:"JobId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Details              string `position:"Query" name:"Details"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ReportAnnotationJobResultRequest) Invoke(client goaliyun.Client) (*ReportAnnotationJobResultResponse, error) {
	resp := &ReportAnnotationJobResultResponse{}
	err := client.Request("mts", "ReportAnnotationJobResult", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ReportAnnotationJobResultResponse struct {
	RequestId goaliyun.String
	JobId     goaliyun.String
}
