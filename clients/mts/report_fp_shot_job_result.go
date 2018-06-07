package mts

import (
	"github.com/morlay/goaliyun"
)

type ReportFpShotJobResultRequest struct {
	Result               string `position:"Query" name:"Result"`
	JobId                string `position:"Query" name:"JobId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Details              string `position:"Query" name:"Details"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ReportFpShotJobResultRequest) Invoke(client goaliyun.Client) (*ReportFpShotJobResultResponse, error) {
	resp := &ReportFpShotJobResultResponse{}
	err := client.Request("mts", "ReportFpShotJobResult", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ReportFpShotJobResultResponse struct {
	RequestId goaliyun.String
	JobId     goaliyun.String
}
