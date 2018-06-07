package mts

import (
	"github.com/morlay/goaliyun"
)

type ReportMediaDetailJobResultRequest struct {
	JobId                string `position:"Query" name:"JobId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Tag                  string `position:"Query" name:"Tag"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Results              string `position:"Query" name:"Results"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ReportMediaDetailJobResultRequest) Invoke(client goaliyun.Client) (*ReportMediaDetailJobResultResponse, error) {
	resp := &ReportMediaDetailJobResultResponse{}
	err := client.Request("mts", "ReportMediaDetailJobResult", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ReportMediaDetailJobResultResponse struct {
	RequestId goaliyun.String
	JobId     goaliyun.String
}
