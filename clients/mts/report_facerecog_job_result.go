package mts

import (
	"github.com/morlay/goaliyun"
)

type ReportFacerecogJobResultRequest struct {
	JobId                string `position:"Query" name:"JobId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Facerecog            string `position:"Query" name:"Facerecog"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Details              string `position:"Query" name:"Details"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ReportFacerecogJobResultRequest) Invoke(client goaliyun.Client) (*ReportFacerecogJobResultResponse, error) {
	resp := &ReportFacerecogJobResultResponse{}
	err := client.Request("mts", "ReportFacerecogJobResult", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ReportFacerecogJobResultResponse struct {
	RequestId goaliyun.String
	JobId     goaliyun.String
}
